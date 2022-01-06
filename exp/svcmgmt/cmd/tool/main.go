package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ActiveState/cli/exp/internal/serve"
	"github.com/ActiveState/cli/exp/procmgmt/internal/proccomm"
	"github.com/ActiveState/cli/internal/exeutils"
)

func main() {
	var (
		id = "default"
	)

	flag.StringVar(&id, "id", id, "identifier")
	flag.Parse()

	pm := proccomm.New("name", id, "")
	addr, err := pm.HTTPAddr()
	if err != nil {
		args := []string{"-id", id}

		if _, err = exeutils.ExecuteAndForget("../svc/build/svc", args); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println("starting service")
		time.Sleep(time.Second)

		addr, err = pm.HTTPAddr()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	c := serve.NewClient(addr)
	info, err := c.GetInfo()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Print(info)
}
