package cmd

import (
	"flag"
	"log"
	"os"
	. "satellite/global"
	"satellite/nets"
)

var httpsCmd = flag.NewFlagSet(CmdHttps, flag.ExitOnError)
var httpsIp string
var httpsPort string

func init() {
	httpsCmd.StringVar(&httpsIp, "ip", "127.0.0.1", "ip address: ipv4 address witch http server listen, such as \"127.0.0.1\"")
	httpsCmd.StringVar(&httpsPort, "port", "9000", "port: port number witch http server listen, such as \"9000\"")
}

func ParseCmdHttps() {
	// check args number
	if len(os.Args) == 2 {
		httpsCmd.Usage()
		os.Exit(1)
	}
	// parse command https
	err := httpsCmd.Parse(os.Args[2:])
	if err != nil {
		log.Println("Error Parse Https Command.")
		os.Exit(1)
	}
	// handle command parameters
	handleCmdHttps(httpsIp, httpsPort)
}

func handleCmdHttps(ip string, port string) {
	nets.StartHttpsServer(ip, port)
}
