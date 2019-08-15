package nets

import (
	"fmt"
	"log"
	"net"
	"os"
	. "satellite/global"
	"strings"
)

func StartTcpServer(ip string, port string) {
	// resolve ip address
	addr, err := net.ResolveTCPAddr("tcp", ip + ":" + port)
	if err != nil {
		log.Println("Error resolve ip address:", err)
		os.Exit(1)
	}
	// start listen tcp
	socket, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Println("Error listen tcp:", err)
		os.Exit(1)
	}
	// loop waiting for connect...
	for {
		conn, err := socket.Accept()
		if err != nil {
			log.Println("Error accept connect:", err)
			continue
		}
		go connectHandler(conn)
	}
}

func connectHandler(c net.Conn) {
	// invalid socket
	if c == nil {
		log.Panicln("Invalid socket connect.")
	}
	// create slice to receive data
	buf := make([]byte, ConstTCPBufferSize)
	// loop receive data stream
	for {
		// start read data stream from net
		n, err := c.Read(buf)
		if n == 0 || err != nil {
			if err != nil {
				log.Println("Error read data stream:", err)
			}
			c.Close()
			break
		}
		// handle data stream(just for demo...)
		str := strings.TrimSpace(string(buf[0:n]))
		fmt.Println(str)
	}
}
