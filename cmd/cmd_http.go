package cmd

import (
	"flag"
	"log"
	"os"
	. "satellite/global"
	"satellite/nets"
)

var httpCmd = flag.NewFlagSet(CmdHttp, flag.ExitOnError)
var httpIp string
var httpPort string

func init() {
	httpCmd.StringVar(&httpIp, "ip", "127.0.0.1", "ip address: ipv4 address witch http server listen, such as \"127.0.0.1\"")
	httpCmd.StringVar(&httpPort, "port", "14514", "port: port number witch http server listen, such as \"14514\"")
}

// ParseCmdHttp function
// this function will be called in main.go and parse and execute one http command
// ip address support both IPv4 and IPv6, you can send parameter like 'localhost', '192.168.1.100', etc.
// port is the integer number between 1~65535, please use the reserve port range (50000~65535), such as '55555'
// any failure or error will output print to screen and exit process
func ParseCmdHttp() {
	// check args number
	if len(os.Args) == 2 {
		httpCmd.Usage()
		os.Exit(1)
	}
	// parse command http
	err := httpCmd.Parse(os.Args[2:])
	if err != nil {
		log.Println("Error Parse Http Command.")
		os.Exit(1)
	}
	// handle command parameters
	handleCmdHttp(httpIp, httpPort)
}

// function handleCmdHttp
// this function mainly handle the main flow of command
// any error will break and exit
func handleCmdHttp(ip string, port string) {
	nets.StartHttpServer(ip, port)
}
