package main

import (
	"fmt"
	"os"
	"os/signal"
	"path"
	"sync"
	"syscall"

	"github.com/ActiveState/cli/exp/procmgmt/internal/procmgmt"
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
	srv := serve.New(":5150")
	srv.Run()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan struct{})

	go func() {
		defer close(done)

		sig := <-sigs
		fmt.Printf("handling signal: %s\n", sig)

		fmt.Println("closing server")
		if err := srv.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		pm := procmgmt.New()
		pm.Listen(done) //nolint // add error handling
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		srv.Wait() //nolint // add error handling
	}()

	wg.Wait()
	return nil
}
