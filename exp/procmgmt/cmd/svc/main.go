package main

import (
	"fmt"
	"os"
	"os/signal"
	"path"
	"sync"
	"syscall"

	"github.com/ActiveState/cli/exp/procmgmt/internal/proccomm"
	"github.com/ActiveState/cli/exp/procmgmt/internal/serve"
)

func main() {
	if err := run(); err != nil {
		cmd := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "%s: %s\n", cmd, err)
		os.Exit(1)
	}
}

func run() error {
	srv := serve.New()
	addr, err := srv.Run()
	if err != nil {
		return err
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan struct{})
	errs := make(chan error)

	go func() {
		defer close(done)

		select {
		case sig := <-sigs:
			fmt.Printf("handling signal: %s\n", sig)
		case err = <-errs:
			fmt.Fprintf(os.Stderr, "errored early: %s\n", err)
		}

		fmt.Println("closing server")
		if err := srv.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		pm := proccomm.New("name", "id", addr)
		if err = pm.Listen(done); err != nil {
			errs <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		srv.Wait() //nolint // add error handling
	}()

	wg.Wait()
	return nil
}
