package events

import (
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/pkg/platform/runtime/artifact"
)

type ProgressOutputer interface {
	BuildStarted(int64) error
	BuildIncrement() error
	BuildCompleted(bool) error

	InstallationStarted(int64) error
	InstallationIncrement() error

	ArtifactStepStarted(artifact.ArtifactID, ArtifactSetupStep, string, int64) error
	ArtifactStepIncrement(artifact.ArtifactID, ArtifactSetupStep, int64) error
	ArtifactStepCompleted(artifact.ArtifactID, ArtifactSetupStep) error
	ArtifactStepFailure(artifact.ArtifactID, ArtifactSetupStep) error
}

// RuntimeEventConsumer is a struct that handles incoming SetupUpdate events in a single go-routine such that they can be forwarded to a progressOutputer.
// State-ful operations should be handled in this struct rather than in the progressOutputer in order to keep the calls to the progressOutputer as simple as possible.
type RuntimeEventConsumer struct {
	progressOut        ProgressOutputer
	totalArtifacts     int64
	numBuildFailures   int64
	numInstallFailures int64
}

func NewRuntimeEventConsumer(progressOut ProgressOutputer) *RuntimeEventConsumer {
	return &RuntimeEventConsumer{
		progressOut: progressOut,
	}
}

func (eh *RuntimeEventConsumer) handleEvent(ev BaseEventer) error {
	switch t := ev.(type) {
	case TotalArtifactEvent:
		eh.totalArtifacts = int64(t.Total())
		return nil
	case BuildStartEvent:
		if eh.totalArtifacts == 0 {
			return errs.New("total number of artifacts has not been set yet.")
		}
		return eh.progressOut.BuildStarted(eh.totalArtifacts)
	case BuildCompleteEvent:
		return eh.progressOut.BuildCompleted(eh.numBuildFailures > 0)
	case ArtifactEventer:
		return eh.handleArtifactEvent(t)
	default:
		logging.Debug("Received unhandled event: %s", ev.String())
	}

	return nil
}

func (eh *RuntimeEventConsumer) handleBuildArtifactEvent(ev ArtifactEventer) error {
	switch t := ev.(type) {
	case ArtifactCompleteEvent:
		return eh.progressOut.BuildIncrement()
	case ArtifactFailureEvent:
		eh.numBuildFailures++
		return nil
	default:
		logging.Debug("unhandled build artifact event: %s", t.String())
	}
	return nil
}

func (eh *RuntimeEventConsumer) handleArtifactEvent(ev ArtifactEventer) error {
	if ev.Step() == Build {
		return eh.handleBuildArtifactEvent(ev)
	}
	switch t := ev.(type) {
	case ArtifactStartEvent:
		if t.Step() == Download {
			if eh.totalArtifacts == 0 {
				return errs.New("total number of artifacts has not been set yet.")
			}
			err := eh.progressOut.InstallationStarted(eh.totalArtifacts)
			if err != nil {
				return err
			}
		}
		name, artBytes := t.ArtifactName(), t.Total()
		return eh.progressOut.ArtifactStepStarted(t.ArtifactID(), t.Step(), name, int64(artBytes))
	case ArtifactProgressEvent:
		by := t.Progress()
		return eh.progressOut.ArtifactStepIncrement(t.ArtifactID(), t.Step(), int64(by))
	case ArtifactCompleteEvent:
		if t.Step() == Install {
			err := eh.progressOut.InstallationIncrement()
			if err != nil {
				return err
			}
		}
		return eh.progressOut.ArtifactStepCompleted(t.ArtifactID(), t.Step())
	case ArtifactFailureEvent:
		eh.numInstallFailures++
		return eh.progressOut.ArtifactStepFailure(t.ArtifactID(), t.Step())
	default:
		logging.Debug("Unhandled artifact event: %s", ev.String())
	}

	return nil
}

// Run should be run in a go routine
func (eh *RuntimeEventConsumer) Run(ch <-chan BaseEventer) error {
	for ev := range ch {
		err := eh.handleEvent(ev)
		if err != nil {
			logging.Error("Cancel progress reporting due to invalid state transition: %w", err)
			// consume remaining events before returning
			for range ch {
			}
			return err
		}
	}
	return nil
}