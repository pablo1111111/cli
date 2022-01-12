package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"syscall"
	"time"

	"github.com/ActiveState/cli/exp/file-opener/internal/syserr"
)

var (
	cmdName   = path.Base(os.Args[0])
	logErrorf = func(format string, a ...interface{}) {
		fmt.Fprintf(os.Stderr, cmdName+": "+format+"\n", a...)
	}
	logf = func(format string, a ...interface{}) {
		fmt.Fprintf(os.Stdout, cmdName+": "+format+"\n", a...)
	}
)

func main() {
	if err := run(); err != nil {
		if errors.Is(err, syserr.EMFILE) {
			logErrorf("EMFILE error detected")
		}
		logErrorf("%s", err)
		os.Exit(1)
	}
}

func run() error {
	var (
		count     int
		tmpDir    = filepath.Join(os.TempDir(), cmdName)
		tmpPrefix = "file"
	)

	flag.IntVar(&count, "c", count, "file count to open")
	flag.Parse()

	if _, err := os.Stat(tmpDir); errors.Is(err, os.ErrNotExist) {
		if err = os.Mkdir(tmpDir, os.ModePerm); err != nil {
			return err
		}
	}
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			logErrorf("failed to rm tmp dir (%s): %s", tmpDir, err)
		}
	}()

	if count <= 0 {
		count = 16
	}
	logf("using file count %d", count)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		logErrorf("handling signal: %s", sig)
		cancel()
	}()

	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			logf("ending early (count: %d)", i)
			break
		default:
		}

		_, err := ioutil.TempFile(tmpDir, tmpPrefix)
		if err != nil {
			return fmt.Errorf("at count %d: %w", i, err)
		}
	}

	select {
	case <-ctx.Done():
	case <-time.After(time.Second * 30):
	}

	return nil
}
