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
		go connectHandler(conn)
	}
}

func connectHandler(c net.Conn) {
	// invalid socket
	if c == nil {
		fmt.Println("Invalid socket connect.")
		log.Println("Invalid socket connect.")
		return
	}
	// create slice to receive data
	buf := make([]byte, ConstTCPBufferSize)
	// loop receive data stream
	for {
		// start read data stream from net
		n, err := c.Read(buf)
		if n == 0 || err != nil {
			if err != nil {
				fmt.Println("Error read data stream:", err)
				log.Println("Error read data stream:", err)
			}
			c.Close()
			break
		}
		// handle data stream
		str := strings.TrimSpace(string(buf[0:n]))
		fmt.Println("[" + c.RemoteAddr().String() + "]:", str)
	}
}
