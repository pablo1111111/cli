package installation

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ActiveState/cli/internal/rtutils"
)

func InstallPath() (string, error) {
	// Facilitate use-case of running executables from the build dir while developing
	if !rtutils.BuiltViaCI && strings.Contains(path.Clean(os.Args[0]), "/build/") {
		return filepath.Dir(os.Args[0]), nil
	}
	return defaultInstallPath()
}