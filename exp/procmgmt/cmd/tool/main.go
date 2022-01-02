package main

import (
	"fmt"
	"os"

	"github.com/ActiveState/cli/exp/procmgmt/internal/procmgmt"
	"github.com/ActiveState/cli/exp/procmgmt/internal/serve"
)

func main() {
	pm := procmgmt.New()
	fmt.Println(pm.OK())

	c := serve.NewClient(":5150")
	info, err := c.GetInfo()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Print(info)
}
