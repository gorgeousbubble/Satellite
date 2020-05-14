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

// ParseCmdTcp function
// this function will be called in main.go and parse and execute one tcp command
// ip address support both IPv4 and IPv6, you can send parameter like 'localhost', '192.168.1.100', etc.
// port is the integer number between 1~65535, please use the reserve port range (50000~65535), such as '55555'
// any failure or error will output print to screen and exit process
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

// function handleCmdTcp
// this function mainly handle the main flow of command
// any error will break and exit
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
