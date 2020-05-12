package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	. "satellite/global"
	"satellite/nets"
)

var rpcCmd = flag.NewFlagSet(CmdRpc, flag.ExitOnError)
var rpcIp string
var rpcPort string
var rpcProtocol string

func init() {
	rpcCmd.StringVar(&rpcIp, "ip", "127.0.0.1", "ip address: ipv4 address witch rpc server listen, such as \"127.0.0.1\"")
	rpcCmd.StringVar(&rpcPort, "port", "13514", "port: port number witch rpc server listen, such as \"13514\"")
	rpcCmd.StringVar(&rpcProtocol, "protocol", "tcp", "protocol: rpc realize protocol, you can choose one from ['tcp','http']")
}

// ParseCmdRpc function
// this function will be called in main.go and parse and execute one qrcode command
// ip address support both IPv4 and IPv6, you can send parameter like 'localhost', '192.168.1.100', etc.
// port is the integer number between 1~65535, please use the reserve port range (50000~65535), such as '55555'
// protocol is the api name which you want to call
// any failure or error will output print to screen and exit process
func ParseCmdRpc() {
	// check args number
	if len(os.Args) == 2 {
		rpcCmd.Usage()
		os.Exit(1)
	}
	// parse command rpc
	err := rpcCmd.Parse(os.Args[2:])
	if err != nil {
		log.Println("Error Parse Rpc Command.")
		os.Exit(1)
	}
	// handle command parameters
	handleCmdRpc(rpcIp, rpcPort, rpcProtocol)
}

// handleCmdRpc function
// this function mainly handle the main flow of command
// any error will break and exit
func handleCmdRpc(ip string, port string, protocol string) {
	switch protocol {
	case "tcp":
		nets.StartRpcTcpServer(ip, port)
	case "http":
		nets.StartRpcHttpServer(ip, port)
	default:
		fmt.Println("Invalid Rpc protocol. You can choose one from ['tcp','http']")
	}
}
