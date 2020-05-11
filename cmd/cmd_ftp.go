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

// ParseCmdFtp function
// this function will be called in main.go and parse and execute one ftp command
// ip address support both IPv4 and IPv6, you can send parameter like 'localhost', '192.168.1.100', etc.
// port is the integer number between 1~65535, please use the reserve port range (50000~65535), such as '55555'
// any failure or error will output print to screen and exit process
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

// function handleCmdFtp
// this function mainly handle the main flow of command
// any error will break and exit
func handleCmdFtp(ip string, port string) {
	nets.StartFtpServer(ip, port)
}
