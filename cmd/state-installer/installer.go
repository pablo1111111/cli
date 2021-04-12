package main

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/shirou/gopsutil/process"

	"github.com/ActiveState/cli/cmd/state-installer/internal/installer"
	"github.com/ActiveState/cli/internal/appinfo"
	"github.com/ActiveState/cli/internal/config"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/exeutils"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/installation"
	"github.com/ActiveState/cli/internal/osutils"
	"github.com/ActiveState/cli/internal/subshell"
	"github.com/ActiveState/cli/internal/subshell/sscommon"
)

type params struct {
	logFile     string
	installPath string
}

func parseParams(args ...string) (*params, error) {
	var p params

	for _, arg := range args[1:] {
		if strings.HasPrefix(arg, "--log-file=") {
			p.logFile = strings.TrimPrefix(arg, "--log-file=")
		} else {
			p.installPath = arg
		}
	}

	if p.installPath == "" {
		installPath, err := installation.InstallPath()
		if err != nil {
			return nil, errs.Wrap(err, "Retrieving installPath")
		}
		p.installPath = installPath
	}

	return &p, nil
}

func main() {
	params, err := parseParams(os.Args...)
	if err != nil {
		log.Printf("Error parsing command line parameters: %v", err)
	}
	// If a log file is set, update the default logger to also append logs to that file. Otherwise it is really difficult to debug what is going on.
	if params.logFile != "" {
		f, err := os.OpenFile(params.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Printf("error initializing log file: %v", err)
		}
		defer f.Close()
		log.SetOutput(io.MultiWriter(os.Stderr, f))
	}

	if err := run(); err != nil {
		// Todo This is running in the background, so these error messages will not be seen and only be written to the log file.
		// https://www.pivotaltracker.com/story/show/177691644
		log.Println(errs.Join(err, ": ").Error())
		log.Printf("To retry run %s", strings.Join(os.Args, " "))
		os.Exit(1)
	}
	log.Println("Installation was successful.")
}

func run() error {
	exe, err := osutils.Executable()
	if err != nil {
		return errs.Wrap(err, "Could not detect executable path")
	}

	cfg, err := config.New()
	if err != nil {
		return errs.Wrap(err, "Could not initialize config")
	}

	var installPath string
	if len(os.Args) > 2 {
		installPath = os.Args[2]
	} else {
		installPath, err = installation.InstallPath()
		if err != nil {
			return errs.Wrap(err, "Retrieving installPath")
		}
	}

	svcInfo := appinfo.SvcApp()
	trayInfo := appinfo.TrayApp()

	// Todo: https://www.pivotaltracker.com/story/show/177585085
	// Yes this is awkward right now
	if err := stopTrayApp(cfg); err != nil {
		return errs.Wrap(err, "Failed to stop %s", trayInfo.Name())
	}

	// Stop state-svc before accessing its files
	if fileutils.FileExists(svcInfo.Exec()) {
		exitCode, _, err := exeutils.Execute(svcInfo.Exec(), []string{"stop"}, nil)
		if err != nil {
			return errs.Wrap(err, "Stopping %s returned error", svcInfo.Name())
		}
		if exitCode != 0 {
			return errs.New("Stopping %s exited with code %d", svcInfo.Name(), exitCode)
		}
	}

	tmpDir := filepath.Dir(exe)
	err = installer.Install(tmpDir, installPath, log.Default())
	if err != nil {
		return errs.Wrap(err, "Installation failed")
	}

	shell := subshell.New(cfg)
	err = shell.WriteUserEnv(cfg, map[string]string{"PATH": installPath}, sscommon.InstallID, true)
	if err != nil {
		return errs.Wrap(err, "Could not update PATH")
	}

	// Run _prepare after updates to facilitate anything the new version of the state tool might need to set up
	// Yes this is awkward, followup story here: https://www.pivotaltracker.com/story/show/176507898
	if stdout, stderr, err := exeutils.ExecSimple(os.Args[0], "_prepare"); err != nil {
		log.Printf("_prepare failed after update: %v\n\nstdout: %s\n\nstderr: %s", err, stdout, stderr)
	}

	if err := exeutils.ExecuteAndForget(trayInfo.Exec()); err != nil {
		return errs.Wrap(err, "Could not start %s", trayInfo.Exec())
	}

	return nil
}

func stopTrayApp(cfg *config.Instance) error {
	trayPid := cfg.GetInt(config.ConfigKeyTrayPid)
	if trayPid <= 0 {
		return nil
	}

	proc, err := process.NewProcess(int32(trayPid))
	if err != nil {
		if errors.Is(err, process.ErrorProcessNotRunning) {
			return nil
		}
		return errs.Wrap(err, "Could not detect if state-tray pid exists")
	}
	if err := proc.Kill(); err != nil {
		return errs.Wrap(err, "Could not kill state-tray")
	}

	return nil
}