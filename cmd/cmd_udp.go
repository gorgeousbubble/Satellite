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

// ParseCmdUdp function
// this function will be called in main.go and parse and execute one udp command
// ip address support both IPv4 and IPv6, you can send parameter like 'localhost', '192.168.1.100', etc.
// port is the integer number between 1~65535, please use the reserve port range (50000~65535), such as '55555'
// any failure or error will output print to screen and exit process
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

// function handleCmdTcp
// this function mainly handle the main flow of command
// any error will break and exit
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
