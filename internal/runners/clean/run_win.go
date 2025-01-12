//go:build windows
// +build windows

package clean

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ActiveState/cli/internal/appinfo"
	"github.com/ActiveState/cli/internal/assets"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/exeutils"
	"github.com/ActiveState/cli/internal/installation"
	"github.com/ActiveState/cli/internal/installation/storage"
	"github.com/ActiveState/cli/internal/language"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/internal/scriptfile"
)

func (u *Uninstall) runUninstall() error {
	// we aggregate installation errors, such that we can display all installation problems in the end
	// TODO: This behavior should be replaced with a proper rollback mechanism https://www.pivotaltracker.com/story/show/178134918
	var aggErr error
	logFile, err := ioutil.TempFile("", "state-clean-uninstall")
	if err != nil {
		aggErr = locale.WrapError(aggErr, "err_clean_logfile", "Could not create temporary log file")
	}

	err = removeInstall(logFile.Name(), u.cfg.ConfigPath(), u.cfg.GetString(installation.CfgTransitionalStateToolPath))
	if err != nil {
		aggErr = locale.WrapError(aggErr, "uninstall_remove_executables_err", "Failed to remove all State Tool files in installation directory {{.V0}}", filepath.Dir(appinfo.StateApp().Exec()))
	}

	err = removeCache(storage.CachePath())
	if err != nil {
		aggErr = locale.WrapError(aggErr, "uninstall_remove_cache_err", "Failed to remove cache directory {{.V0}}.", storage.CachePath())
	}

	err = undoPrepare(u.cfg)
	if err != nil {
		aggErr = locale.WrapError(aggErr, "uninstall_prepare_err", "Failed to undo some installation steps.")
	}

	err = removeEnvPaths(u.cfg)
	if err != nil {
		aggErr = locale.WrapError(aggErr, "uninstall_remove_paths_err", "Failed to remove PATH entries from environment")
	}

	if aggErr != nil {
		return aggErr
	}

	u.out.Print(locale.Tr("clean_message_windows", logFile.Name()))
	return nil
}

func removeConfig(configPath string, out output.Outputer) error {
	logFile, err := ioutil.TempFile("", "state-clean-config")
	if err != nil {
		return locale.WrapError(err, "err_clean_logfile", "Could not create temporary log file")
	}

	out.Print(locale.Tr("clean_config_message_windows", logFile.Name()))
	return removePaths(logFile.Name(), configPath)
}

func removeInstall(logFile, configPath, transitionalStateTool string) error {
	svcInfo := appinfo.SvcApp()
	trayInfo := appinfo.TrayApp()
	var aggErr error
	for _, info := range []*appinfo.AppInfo{svcInfo, trayInfo} {
		if err := os.Remove(info.LegacyExec()); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				aggErr = errs.Wrap(aggErr, "Could not remove (legacy) %s: %v", info.LegacyExec(), err)
			}
		}

		err := os.Remove(info.Exec())
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			aggErr = locale.WrapError(aggErr, "uninstall_rm_exec", "Could not remove executable: {{.V0}}. Error: {{.V1}}.", info.Exec(), err.Error())
		}
	}

	if aggErr != nil {
		return aggErr
	}

	paths := []string{appinfo.StateApp().Exec(), configPath}
	// If the transitional state tool path is known, we remove it.  This is done in the background, because the transitional State Tool can be the initiator of the uninstall request
	if transitionalStateTool != "" {
		paths = append(paths, transitionalStateTool)
	}

	return removePaths(logFile, paths...)
}

func removePaths(logFile string, paths ...string) error {
	logging.Debug("Removing paths: %v", paths)
	scriptName := "removePaths"
	scriptBlock, err := assets.ReadFileBytes(fmt.Sprintf("scripts/%s.bat", scriptName))
	if err != nil {
		return err
	}
	sf, err := scriptfile.New(language.Batch, scriptName, string(scriptBlock))
	if err != nil {
		return locale.WrapError(err, "err_clean_script", "Could not create new scriptfile")
	}

	exe, err := os.Executable()
	if err != nil {
		return locale.WrapError(err, "err_clean_executable", "Could not get executable name")
	}

	args := []string{"/C", sf.Filename(), logFile, fmt.Sprintf("%d", os.Getpid()), filepath.Base(exe)}
	args = append(args, paths...)
	_, err = exeutils.ExecuteAndForget("cmd.exe", args)
	if err != nil {
		return locale.WrapError(err, "err_clean_start", "Could not start remove direcotry script")
	}

	return nil
}
