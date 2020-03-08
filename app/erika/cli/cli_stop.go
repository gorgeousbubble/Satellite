package cli

import (
	"flag"
	"fmt"
	"os"
	"satellite/app/erika/app"
	"satellite/app/erika/logs"
)

var stopCli = flag.NewFlagSet("stop", flag.ExitOnError)
var stopHelp bool

func init() {
	stopCli.BoolVar(&stopHelp, "help", false, "usage of erika stop")
}

func ParseCliStop() {
	// parse cli stop
	err := stopCli.Parse(os.Args[2:])
	if err != nil {
		logs.Error("Error Parse stop Command.")
		os.Exit(1)
	}
	// help info
	if stopHelp {
		stopCli.Usage()
		os.Exit(1)
	}
	// handle command parameters
	handleCliStop()
}

func handleCliStop() {
	err := app.Stop()
	if err != nil {
		fmt.Println("Error stop Erika:", err)
		os.Exit(1)
	}
}
