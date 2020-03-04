package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	. "satellite/global"
	"satellite/nets"
)

var udpCmd = flag.NewFlagSet(CmdUdp, flag.ExitOnError)
var udpIp string
var udpPort string
var udpMode string

func init() {
	udpCmd.StringVar(&udpIp, "ip", "127.0.0.1", "ip address: ipv4 address witch udp server listen, such as \"127.0.0.1\"")
	udpCmd.StringVar(&udpPort, "port", "12514", "port: port number witch udp server listen, such as \"12514\"")
	udpCmd.StringVar(&udpMode, "mode", "server", "mode: udp mode choose, 's' or 'server' indicate udp server, 'c' or 'client' indicate udp client")
}

func ParseCmdUdp() {
	// check args number
	if len(os.Args) == 2 {
		udpCmd.Usage()
		os.Exit(1)
	}
	// parse command udp
	err := udpCmd.Parse(os.Args[2:])
	if err != nil {
		log.Println("Error Parse Udp Command.")
		os.Exit(1)
	}
	// handle command parameters
	handleCmdUdp(udpIp, udpPort, udpMode)
}

func handleCmdUdp(ip string, port string, mode string) {
	switch mode {
	case "s", "server":
		nets.StartUdpServer(ip, port)
	case "c", "client":
		nets.StartUdpClient(ip, port)
	default:
		fmt.Println("Invalid Udp Mode. You can input 'c' stand for 'client' or 's' for 'server'.")
	}
}
