package nets

import (
	"fmt"
	"log"
	"net"
	"os"
)

func StartUdpClient(ip string, port string) {
	// resolve ip address
	fmt.Println("Start Udp Client")
	addr, err := net.ResolveUDPAddr("udp", ip+":"+port)
	if err != nil {
		fmt.Println("Error resolve ip address:", err)
		log.Println("Error resolve ip address:", err)
		os.Exit(1)
	}
	// connect to tcp server
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error connect udp:", err)
		log.Println("Error connect udp:", err)
		os.Exit(1)
	}
	fmt.Println("Connect udp success.")
	// connect send handler
	go connUdpRecvHandler(conn)
	connUdpSendHandler(conn)
}
