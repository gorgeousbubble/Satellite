package nets

import (
	"fmt"
	"log"
	"net"
	"os"
)

func StartUdpServer(ip string, port string) {
	// resolve ip address
	fmt.Println("Start Udp Server")
	addr, err := net.ResolveUDPAddr("udp", ip+":"+port)
	if err != nil {
		fmt.Println("Error resolve ip address:", err)
		log.Println("Error resolve ip address:", err)
		os.Exit(1)
	}
	// start listen tcp
	fmt.Println("Listen Udp:", ip+":"+port)
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listen udp:", err)
		log.Println("Error listen udp:", err)
		os.Exit(1)
	}
	defer conn.Close()
	// loop waiting for send message...
	fmt.Println("Loop waiting for send message...")
	connUdpRecvHandler(conn)
}
