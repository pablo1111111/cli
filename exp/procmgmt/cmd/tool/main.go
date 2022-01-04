package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ActiveState/cli/exp/procmgmt/internal/proccomm"
	"github.com/ActiveState/cli/exp/procmgmt/internal/serve"
	"github.com/ActiveState/cli/internal/exeutils"
)

func main() {
	pm := proccomm.New("name", "id", "")
	addr, err := pm.HTTPAddr()
	if err != nil {
		if _, err = exeutils.ExecuteAndForget("../svc/build/svc", nil); err != nil {
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
