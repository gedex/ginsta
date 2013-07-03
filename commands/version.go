package commands

import (
	"fmt"
	"os"
)

var cmdVersion = &Command{
	Callback: runVersion,
	Usage:    "version",
	Short:    "Show version",
	Long:     `Shows this client version.`,
}

func runVersion(r *Runner, cmd *Command, args []string) {
	fmt.Println(r.Client.Name, "version", r.Client.Version)
	os.Exit(0)
}
