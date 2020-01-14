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
	rpcCmd.StringVar(&rpcPort, "port", "2514", "port: port number witch rpc server listen, such as \"2514\"")
	rpcCmd.StringVar(&rpcProtocol, "protocol", "tcp", "protocol: rpc realize protocol, you can choose one from ['tcp','http']")
}

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

func handleCmdRpc(ip string, port string, protocol string) {
	switch protocol {
	case "tcp":
		err := nets.StartRpcTcpServer(ip, port)
		if err != nil {
			fmt.Println("Error start rpc tcp server:", err)
			log.Println("Error start rpc tcp server:", err)
			os.Exit(1)
		}
	case "http":
		err := nets.StartRpcHttpServer(ip, port)
		if err != nil {
			fmt.Println("Error start rpc http server:", err)
			log.Println("Error start rpc http server:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid Rpc protocol. You can choose one from ['tcp','http']")
	}
}
