package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	. "satellite/global"
	"satellite/nets"
)

var tcpCmd = flag.NewFlagSet(CmdTcp, flag.ExitOnError)
var tcpIp string
var tcpPort string
var tcpMode string

func init() {
	tcpCmd.StringVar(&tcpIp, "ip", "127.0.0.1", "ip address: ipv4 address witch tcp server listen, such as \"127.0.0.1\"")
	tcpCmd.StringVar(&tcpPort, "port", "11514", "port: port number witch tcp server listen, such as \"11514\"")
	tcpCmd.StringVar(&tcpMode, "mode", "server", "mode: tcp mode choose, 's' or 'server' indicate tcp server, 'c' or 'client' indicate tcp client")
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
	handleCmdTcp(tcpIp, tcpPort, tcpMode)
}

func handleCmdTcp(ip string, port string, mode string) {
	switch mode {
	case "s", "server":
		nets.StartTcpServer(ip, port)
	case "c", "client":
		nets.StartTcpClient(ip, port)
	default:
		fmt.Println("Invalid Tcp Mode. You can input 'c' stand for 'client' or 's' for 'server'.")
	}
}
