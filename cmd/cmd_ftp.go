package cmd

import (
	"flag"
	"log"
	"os"
	. "satellite/global"
	"satellite/nets"
)

var ftpCmd = flag.NewFlagSet(CmdFtp, flag.ExitOnError)
var ftpIp string
var ftpPort string

func init() {
	ftpCmd.StringVar(&ftpIp, "ip", "127.0.0.1", "ip address: ipv4 address witch http server listen, such as \"127.0.0.1\"")
	ftpCmd.StringVar(&ftpPort, "port", "16514", "port: port number witch http server listen, such as \"16514\"")
}

func ParseCmdFtp() {
	// check args number
	if len(os.Args) == 2 {
		ftpCmd.Usage()
		os.Exit(1)
	}
	// parse command http
	err := ftpCmd.Parse(os.Args[2:])
	if err != nil {
		log.Println("Error Parse Ftp Command.")
		os.Exit(1)
	}
	// handle command parameters
	handleCmdFtp(ftpIp, ftpPort)
}

func handleCmdFtp(ip string, port string) {
	nets.StartFtpServer(ip, port)
}
