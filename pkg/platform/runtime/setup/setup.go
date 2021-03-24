package setup

import (
	"context"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/gammazero/workerpool"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/internal/analytics"
	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/download"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/runbits/proxyreader"
	"github.com/ActiveState/cli/internal/unarchiver"
	"github.com/ActiveState/cli/pkg/platform/api/buildlogstream"
	"github.com/ActiveState/cli/pkg/platform/api/headchef"
	"github.com/ActiveState/cli/pkg/platform/api/headchef/headchef_models"
	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
	apimodel "github.com/ActiveState/cli/pkg/platform/model"
	"github.com/ActiveState/cli/pkg/platform/runtime/artifact"
	"github.com/ActiveState/cli/pkg/platform/runtime/envdef"
	"github.com/ActiveState/cli/pkg/platform/runtime/model"
	"github.com/ActiveState/cli/pkg/platform/runtime/setup/buildlog"
	"github.com/ActiveState/cli/pkg/platform/runtime/setup/events"
	"github.com/ActiveState/cli/pkg/platform/runtime/setup/implementations/alternative"
	"github.com/ActiveState/cli/pkg/platform/runtime/setup/implementations/camel"
	"github.com/ActiveState/cli/pkg/platform/runtime/store"
	"github.com/ActiveState/cli/pkg/project"
)

// MaxConcurrency is maximum number of parallel artifact installations
const MaxConcurrency = 10

// NotInstalledError is an error returned when the runtime is not completely installed yet.
var NotInstalledError = errs.New("Runtime is not completely installed.")

type ArtifactSetupErrors struct {
	errs []error
}

func (a *ArtifactSetupErrors) Error() string {
	var errors []string
	for _, err := range a.errs {
		errors = append(errors, errs.Join(err, " :: ").Error())
	}
	return "Not all artifacts could be installed, errors: \n" + strings.Join(errors, "\n")
}

func (a *ArtifactSetupErrors) Errors() []error {
	return a.errs
}

func (a *ArtifactSetupErrors) UserError() string {
	var errStrings []string
	for _, err := range a.errs {
		errStrings = append(errStrings, locale.JoinErrors(err, ":").UserError())
	}
	return locale.Tl("setup_artifacts_err", "Not all artifacts could be installed: \n", strings.Join(errStrings, "\n"))
}

// MessageHandler is the interface for callback functions that are called during
// runtime set-up when progress messages can be forwarded to the user
type MessageHandler interface {
	buildlog.BuildLogMessageHandler

	// ChangeSummary summarizes the changes to the current project during the InstallRuntime() call.
	// This summary is printed as soon as possible, providing the State Tool user with an idea of the complexity of the requested build.
	// The arguments are for the changes introduced in the latest commit that this Setup is setting up.
	// TODO: Decide if we want to have a method to de-activate the change summary for activations where it does not make sense.
	ChangeSummary(artifacts map[artifact.ArtifactID]artifact.ArtifactRecipe, requested artifact.ArtifactChangeset, changed artifact.ArtifactChangeset)
	TotalArtifacts(total int)
	ArtifactStepStarting(events.ArtifactSetupStep, artifact.ArtifactID, string, int)
	ArtifactStepProgress(events.ArtifactSetupStep, artifact.ArtifactID, int)
	ArtifactStepCompleted(events.ArtifactSetupStep, artifact.ArtifactID)
	ArtifactStepFailed(events.ArtifactSetupStep, artifact.ArtifactID, string)
}

type Targeter interface {
	CommitUUID() strfmt.UUID
	Name() string
	Owner() string
	Dir() string
}

// Setup provides methods to setup a fully-function runtime that *only* requires interactions with the local file system.
type Setup struct {
	model      ModelProvider
	target     Targeter
	msgHandler MessageHandler
	store      *store.Store
}

// ModelProvider is the interface for all functions that involve backend communication
type ModelProvider interface {
	ResolveRecipe(commitID strfmt.UUID, owner, projectName string) (*inventory_models.Recipe, error)
	RequestBuild(recipeID, commitID strfmt.UUID, owner, project string) (headchef.BuildStatusEnum, *headchef_models.BuildStatusResponse, error)
	FetchBuildResult(commitID strfmt.UUID, owner, project string) (*model.BuildResult, error)
	SignS3URL(uri *url.URL) (*url.URL, error)
}

// ArtifactSetuper is the interface for an implementation of artifact setup functions
// These need to be specialized for each BuildEngine type
type ArtifactSetuper interface {
	EnvDef(tmpInstallDir string) (*envdef.EnvironmentDefinition, error)
	Unarchiver() unarchiver.Unarchiver
}

// New returns a new Setup instance that can install a Runtime locally on the machine.
func New(target Targeter, msgHandler MessageHandler) *Setup {
	return NewWithModel(target, msgHandler, model.NewDefault())
}

// NewWithModel returns a new Setup instance with a customized model eg., for testing purposes
func NewWithModel(target Targeter, msgHandler MessageHandler, model ModelProvider) *Setup {
	return &Setup{model, target, msgHandler, nil}
}

// Update installs the runtime locally (or updates it if it's already partially installed)
func (s *Setup) Update() error {
	err := s.update()
	if err != nil {
		analytics.EventWithLabel(analytics.CatRuntime, analytics.ActRuntimeFailure, analytics.LblRtFailUpdate)
		return err
	}
	return nil
}

func (s *Setup) update() error {
	// Request build
	buildResult, err := s.model.FetchBuildResult(s.target.CommitUUID(), s.target.Owner(), s.target.Name())
	if err != nil {
		return errs.Wrap(err, "Failed to fetch build result")
	}

	if buildResult.BuildStatus == headchef.Started {
		analytics.Event(analytics.CatRuntime, analytics.ActRuntimeBuild)
		ns := project.Namespaced{
			Owner:   s.target.Owner(),
			Project: s.target.Name(),
		}
		analytics.EventWithLabel(analytics.CatRuntime, analytics.ActBuildProject, ns.String())
	}

	// Compute and handle the change summary
	artifacts := artifact.NewMapFromRecipe(buildResult.Recipe)

	s.store, err = store.New(s.target.Dir())
	if err != nil {
		return errs.Wrap(err, "Could not create runtime store")
	}
	oldRecipe, err := s.store.Recipe()
	if err != nil {
		logging.Debug("Could not load existing recipe.  Maybe it is a new installation: %v", err)
	}
	requestedArtifacts := artifact.NewArtifactChangesetByRecipe(oldRecipe, buildResult.Recipe, true)
	changedArtifacts := artifact.NewArtifactChangesetByRecipe(oldRecipe, buildResult.Recipe, false)
	s.msgHandler.ChangeSummary(artifacts, requestedArtifacts, changedArtifacts)

	storedArtifacts, err := s.store.Artifacts()
	if err != nil {
		return locale.WrapError(err, "err_stored_artifacts", "Could not unmarshal stored artifacts, your install may be corrupted.")
	}
	s.deleteOutdatedArtifacts(changedArtifacts, storedArtifacts)

	// if we get here, we dowload artifacts
	analytics.Event(analytics.CatRuntime, analytics.ActRuntimeDownload)

	s.msgHandler.TotalArtifacts(len(artifacts))
	if buildResult.BuildReady {
		err := s.installFromBuildResult(buildResult, artifacts)
		if err != nil {
			return err
		}
	} else {
		if buildResult.BuildEngine == model.Camel {
			return locale.NewInputError("build_status_in_progress", "", apimodel.ProjectURL(s.target.Owner(), s.target.Name(), s.target.CommitUUID().String()))
		}
		err := s.installFromBuildLog(buildResult, artifacts)
		if err != nil {
			return err
		}
	}

	if err := s.store.UpdateEnviron(buildResult.OrderedArtifacts()); err != nil {
		return errs.Wrap(err, "Could not save combined environment file")
	}

	// clean up temp directory
	tempDir := filepath.Join(s.store.InstallPath(), constants.LocalRuntimeTempDirectory)
	err = os.RemoveAll(tempDir)
	if err != nil {
		logging.Errorf("Failed to remove temporary installation directory %s: %v", tempDir, err)
	}

	if err := s.store.MarkInstallationComplete(s.target.CommitUUID()); err != nil {
		return errs.Wrap(err, "Could not mark install as complete.")
	}

	return nil
}

func (s *Setup) deleteOutdatedArtifacts(changeset artifact.ArtifactChangeset, storedArtifacted map[artifact.ArtifactID]store.StoredArtifact) error {
	del := map[strfmt.UUID]struct{}{}
	for _, upd := range changeset.Updated {
		del[upd.FromID] = struct{}{}
	}
	for _, id := range changeset.Removed {
		del[id] = struct{}{}
	}

	for _, artf := range storedArtifacted {
		if _, deleteMe := del[artf.ArtifactID]; !deleteMe {
			continue
		}

		for _, file := range artf.Files {
			if !fileutils.TargetExists(file) {
				continue // don't care it's already deleted (might have been deleted by another artifact that supplied the same file)
			}
			if err := os.Remove(file); err != nil {
				return locale.WrapError(err, "err_rm_artf", "", "Could not remove old package file at {{.V0}}.", file)
			}
		}

		if err := s.store.DeleteArtifactStore(artf.ArtifactID); err != nil {
			return errs.Wrap(err, "Could not delete artifact store")
		}
	}

	return nil
}

func (s *Setup) installFromBuildResult(buildResult *model.BuildResult, artifacts map[artifact.ArtifactID]artifact.ArtifactRecipe) error {
	var errors []error
	wp := workerpool.New(MaxConcurrency)
	downloads, err := artifact.NewDownloadsFromBuild(buildResult.BuildStatusResponse, buildResult.BuildEngine == model.Camel)
	if err != nil {
		return errs.Wrap(err, "Could not fetch artifacts to download.")
	}
	for _, a := range downloads {
		func(a artifact.ArtifactDownload) {
			wp.Submit(func() {
				name := "unknown"
				if a, ok := artifacts[a.ArtifactID]; ok {
					name = a.Name
				}
				if err := s.setupArtifact(buildResult.BuildEngine, a.ArtifactID, a.UnsignedURI, name); err != nil {
					errors = append(errors, err)
				}
			})
		}(a)
	}

	wp.StopWait()

	if len(errors) > 0 {
		return &ArtifactSetupErrors{errors}
	}

	return nil
}

func (s *Setup) installFromBuildLog(buildResult *model.BuildResult, artifacts map[artifact.ArtifactID]artifact.ArtifactRecipe) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := buildlogstream.Connect(ctx)
	if err != nil {
		return errs.Wrap(err, "Could not get build updates")
	}
	defer conn.Close()

	buildLog, err := buildlog.New(artifacts, conn, s.msgHandler, *buildResult.Recipe.RecipeID)

	var errors []error
	wp := workerpool.New(MaxConcurrency)

	go func() {
		for a := range buildLog.BuiltArtifactsChannel() {
			func(a artifact.ArtifactDownload) {
				wp.Submit(func() {
					name := artifacts[a.ArtifactID].Name
					if err := s.setupArtifact(buildResult.BuildEngine, a.ArtifactID, a.UnsignedURI, name); err != nil {
						errors = append(errors, err)
					}
				})
			}(a)
		}
	}()

	if err = buildLog.Wait(); err != nil {
		return err
	}

	wp.StopWait()

	if len(errors) > 0 {
		return &ArtifactSetupErrors{errors}
	}

	return nil
}

// setupArtifact sets up an individual artifact
// The artifact is downloaded, unpacked and then processed by the artifact setup implementation
func (s *Setup) setupArtifact(buildEngine model.BuildEngine, a artifact.ArtifactID, unsignedURI, artifactName string) error {
	as, err := s.selectArtifactSetupImplementation(buildEngine, a)
	if err != nil {
		return errs.Wrap(err, "Failed to select artifact setup implementation")
	}

	targetDir := filepath.Join(s.store.InstallPath(), constants.LocalRuntimeTempDirectory)
	if err := fileutils.MkdirUnlessExists(targetDir); err != nil {
		return errs.Wrap(err, "Could not create temp runtime dir")
	}

	unarchiver := as.Unarchiver()
	archivePath := filepath.Join(targetDir, a.String()+unarchiver.Ext())
	downloadProgress := events.NewSubProgressProducer(s.msgHandler, events.Download, a, artifactName)
	if err := s.downloadArtifact(unsignedURI, archivePath, downloadProgress); err != nil {
		err := errs.Wrap(err, "Could not download artifact %s", unsignedURI)
		s.msgHandler.ArtifactStepFailed(events.Download, a, err.Error())
		return err
	}
	s.msgHandler.ArtifactStepCompleted(events.Download, a)

	unpackedDir := filepath.Join(targetDir, a.String())
	logging.Debug("Unarchiving %s (%s) to %s", archivePath, unsignedURI, unpackedDir)
	// clean up the unpacked dir
	defer os.RemoveAll(unpackedDir)

	unpackProgress := events.NewSubProgressProducer(s.msgHandler, events.Unpack, a, artifactName)
	numFiles, err := s.unpackArtifact(unarchiver, archivePath, unpackedDir, unpackProgress)
	if err != nil {
		err := errs.Wrap(err, "Could not unpack artifact %s", archivePath)
		s.msgHandler.ArtifactStepFailed(events.Unpack, a, err.Error())
		return err
	}
	s.msgHandler.ArtifactStepCompleted(events.Unpack, a)

	envDef, err := as.EnvDef(unpackedDir)
	if err != nil {
		return errs.Wrap(err, "Could not collect env info for artifact")
	}

	var files []string
	onMoveFile := func(fromPath, toPath string) {
		if !fileutils.IsDir(toPath) {
			s.msgHandler.ArtifactStepProgress(events.Install, a, 1)
			files = append(files, toPath)
		}
	}
	s.msgHandler.ArtifactStepStarting(events.Install, a, artifactName, numFiles)
	err = fileutils.MoveAllFilesRecursively(
		filepath.Join(unpackedDir, envDef.InstallDir),
		s.store.InstallPath(), onMoveFile,
	)
	if err != nil {
		err := errs.Wrap(err, "Move artifact failed")
		s.msgHandler.ArtifactStepFailed(events.Install, a, err.Error())
		return err
	}
	s.msgHandler.ArtifactStepCompleted(events.Install, a)

	cnst := envdef.NewConstants(s.store.InstallPath())
	envDef = envDef.ExpandVariables(cnst)
	err = envDef.ApplyFileTransforms(s.store.InstallPath(), cnst)
	if err != nil {
		return locale.WrapError(err, "runtime_alternative_file_transforms_err", "", "Could not apply necessary file transformations after unpacking")
	}

	// Install PPM Shim if artifact provides Perl executable
	if activePerlPath := envDef.FindBinPathFor(constants.ActivePerlExecutable); activePerlPath != "" {
		err = installPPMShim(activePerlPath)
		if err != nil {
			return errs.Wrap(err, "Failed to install the PPM shim command at %s", activePerlPath)
		}
	}

	if err := s.store.StoreArtifact(store.NewStoredArtifact(a, files, envDef)); err != nil {
		return errs.Wrap(err, "Could not store artifact meta info")
	}

	return nil
}

// downloadArtifact retrieves the tarball for an artifactID
// Note: the tarball may also be retrieved from a local cache directory if that is available.
func (s *Setup) downloadArtifact(unsignedURI string, targetFile string, progress *events.SubProgressProducer) error {
	artifactURL, err := url.Parse(unsignedURI)
	if err != nil {
		return errs.Wrap(err, "Could not parse artifact URL %s.", unsignedURI)
	}

	downloadURL, err := s.model.SignS3URL(artifactURL)
	if err != nil {
		return errs.Wrap(err, "Could not sign artifact URL %s.", unsignedURI)
	}

	b, err := download.GetWithProgress(downloadURL.String(), progress)
	if err != nil {
		return errs.Wrap(err, "Download %s failed", downloadURL)
	}
	if err := fileutils.WriteFile(targetFile, b); err != nil {
		return errs.Wrap(err, "Writing download to target file %s failed", targetFile)
	}
	return nil
}

func (s *Setup) unpackArtifact(ua unarchiver.Unarchiver, tarballPath string, targetDir string, progress *events.SubProgressProducer) (int, error) {
	f, i, err := ua.PrepareUnpacking(tarballPath, targetDir)
	progress.TotalSize(int(i))
	defer f.Close()
	if err != nil {
		return 0, errs.Wrap(err, "Prepare for unpacking failed")
	}
	var numUnpackedFiles int
	ua.SetNotifier(func(_ string, _ int64, isDir bool) {
		if !isDir {
			numUnpackedFiles++
		}
	})
	proxy := proxyreader.NewProxyReader(progress, f)
	return numUnpackedFiles, ua.Unarchive(proxy, i, targetDir)
}

func (s *Setup) selectArtifactSetupImplementation(buildEngine model.BuildEngine, a artifact.ArtifactID) (ArtifactSetuper, error) {
	switch buildEngine {
	case model.Alternative:
		return alternative.NewArtifactSetup(a, s.store), nil
	case model.Camel:
		return camel.NewArtifactSetup(a, s.store), nil
	default:
		return nil, errs.New("Unknown build engine: %s", buildEngine)
	}
}