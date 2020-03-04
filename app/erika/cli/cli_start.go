package cli

import (
	"flag"
	"os"
	"satellite/app/erika/app"
	"satellite/app/erika/logs"
)

var startCli = flag.NewFlagSet("start", flag.ExitOnError)
var startHelp bool
var startIp string
var startPort string

func init() {
	startCli.StringVar(&startIp, "ip", "127.0.0.1", "ip address: ipv4 address witch http server listen, such as \"127.0.0.1\"")
	startCli.StringVar(&startPort, "port", "14514", "port: port number witch http server listen, such as \"14514\"")
	startCli.BoolVar(&startHelp, "help", false, "usage of erika start")
}

func ParseCliStart() {
	// parse cli start
	err := startCli.Parse(os.Args[2:])
	if err != nil {
		logs.Error("Error Parse Http Command.")
		os.Exit(1)
	}
	// help info
	if startHelp {
		startCli.Usage()
		os.Exit(1)
	}
	// handle command parameters
	handleCliStart(startIp, startPort)
}

func handleCliStart(ip string, port string) {
	app.Run(ip, port)
}
