package nets

import (
	"fmt"
	"log"
	"net"
	"os"
)

func StartTcpServer(ip string, port string) {
	// resolve ip address
	fmt.Println("Start Tcp Server")
	addr, err := net.ResolveTCPAddr("tcp", ip + ":" + port)
	if err != nil {
		fmt.Println("Error resolve ip address:", err)
		log.Println("Error resolve ip address:", err)
		os.Exit(1)
	}
	// start listen tcp
	fmt.Println("Listen Tcp:", ip + ":" + port)
	socket, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println("Error listen tcp:", err)
		log.Println("Error listen tcp:", err)
		os.Exit(1)
	}
	// loop waiting for connect...
	fmt.Println("Loop waiting for connect...")
	for {
		conn, err := socket.Accept()
		if err != nil {
			fmt.Println("Error accept connect:", err)
			log.Println("Error accept connect:", err)
			continue
		}
		fmt.Println("Success accept client:", conn.RemoteAddr().String())
		go connRecvHandler(conn)
	}
}
