package scriptfile

import (
	"os"
	"path"
	"runtime"
	"testing"
)

func TestScriptFile(t *testing.T) {
	sf, err := New(Bash, "echo hello")
	noError(t.Fatalf, err)

	func() { // scope for cleanup
		defer sf.Clean()

		t.Run("file name has extension", func(t *testing.T) {
			ext := path.Ext(sf.Filename())
			gt(t.Errorf, int64(len(ext)), 0)
		})

		info, err := os.Stat(sf.Filename())

		t.Run("file exists", func(t *testing.T) {
			if info == nil {
				t.Fatalf("got %v, want os.FileInfo", info)
			}
			noError(t.Fatalf, err)
		})

		t.Run("file not empty", func(t *testing.T) {
			gt(t.Errorf, info.Size(), int64(len("echo hello")))
		})

		t.Run("file executable", func(t *testing.T) {
			if runtime.GOOS == "windows" {
				gt(t.Errorf, int64(0400&info.Mode()), 0)
				return
			}

			gt(t.Errorf, int64(0110&info.Mode()), 0)
		})
	}()

	t.Run("file cleaned up", func(t *testing.T) {
		_, err := os.Stat(sf.Filename())
		if err == nil || !os.IsNotExist(err) {
			t.Errorf("got %v, want not exist error", err)
		}
	})

	t.Run("file lacking header", func(t *testing.T) {
		sf, err = New(Batch, "echo hello")
		noError(t.Fatalf, err)
		info, err := os.Stat(sf.Filename())
		noError(t.Fatalf, err)

		eq(t.Errorf, info.Size(), int64(len("echo hello")))
	})
}

type failFunc func(format string, args ...interface{})

func noError(ff failFunc, err error) {
	if err != nil {
		ff("unexpected error: %v", err)
	}
}

func gt(ff failFunc, a, b int64) {
	if a <= b {
		ff("got %v, want > %v", a, b)
	}
}

func eq(ff failFunc, a, b int64) {
	if a != b {
		ff("got %v, want %v", a, b)
	}
}