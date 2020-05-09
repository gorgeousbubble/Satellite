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
	httpsCmd.StringVar(&httpsPort, "port", "15514", "port: port number witch http server listen, such as \"15514\"")
}

// ParseCmdHttps function
// this function will be called in main.go and parse and execute one https command
// ip address support both IPv4 and IPv6, you can send parameter like 'localhost', '192.168.1.100', etc.
// port is the integer number between 1~65535, please use the reserve port range (50000~65535), such as '55555'
// any failure or error will output print to screen and exit process
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

// function handleCmdHttps
// this function mainly handle the main flow of command
// any error will break and exit
func handleCmdHttps(ip string, port string) {
	nets.StartHttpsServer(ip, port)
}
