package nets

import (
	"fmt"
	"log"
	"net"
	"os"
)

func StartTcpClient(ip string, port string) {
	// connect to tcp server
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Error connect to server:", err)
		log.Println("Error connect to server:", err)
		os.Exit(1)
	}
	fmt.Println("Connect to server success.")
	// connect send handler
	go connTcpRecvHandler(conn)
	connTcpSendHandler(conn)
}
