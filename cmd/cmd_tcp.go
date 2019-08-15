package cmd

import (
	"flag"
	"log"
	"os"
	. "satellite/global"
	"satellite/nets"
)

var tcpCmd = flag.NewFlagSet(CmdTcp, flag.ExitOnError)
var tcpIp string
var tcpPort string

func init() {
	tcpCmd.StringVar(&tcpIp, "ip", "127.0.0.1", "ip address: ipv4 address witch tcp server listen, such as \"127.0.0.1\"")
	tcpCmd.StringVar(&tcpPort, "port", "6000", "port: port number witch tcp server listen, such as \"6000\"")
}

func ParseCmdTcp() {
	// check args number
	if len(os.Args) == 2 {
		tcpCmd.Usage()
		os.Exit(1)
	}
	// parse command tcp
	err := tcpCmd.Parse(os.Args[2:])
	if err != nil {
		log.Println("Error Parse Tcp Command.")
		os.Exit(1)
	}
	// handle command parameters
	handleCmdTcp(tcpIp, tcpPort)
}

func handleCmdTcp(ip string, port string) {
	nets.StartTcpServer(ip, port)
}
