package nets

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	. "satellite/global"
	"strconv"
	"strings"
	"time"
)

func StartTcpClient(ip string, port string) {
	// connect to tcp server
	conn, err := net.Dial("tcp", ip + ":" + port)
	if err != nil {
		fmt.Println("Error connect to server:", err)
		log.Println("Error connect to server:", err)
		os.Exit(1)
	}
	fmt.Println("Connect to server success.")
	// connect send handler
	go connRecvHandler(conn)
	connSendHandler(conn)
}

func connSendHandler(c net.Conn) {
	// create stdin reader
	rd := bufio.NewReader(os.Stdin)
	// loop waiting for send message...
	for {
		// read stdin data stream
		in, err := rd.ReadString('\n')
		if err != nil {
			fmt.Println("Error read stdin data stream:", err)
			log.Println("Error read stdin data stream:", err)
			break
		}
		// handle data stream
		in = strings.TrimSpace(in)
		// send data stream
		_, err = c.Write([]byte(in))
		if err != nil {
			fmt.Println("Error write data stream:", err)
			log.Println("Error write data stream:", err)
			break
		}
		t := time.Now()
		fmt.Println("[" + c.LocalAddr().String() + "] ", strconv.Itoa(t.Hour()) + ":" + strconv.Itoa(t.Minute()) + ":" + strconv.Itoa(t.Second()))
		fmt.Println("Remote<-Local:", in)
	}
}

func connRecvHandler(c net.Conn) {
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
				fmt.Println("Remote host forcibly closed connect:", c.RemoteAddr().String())
				log.Println("Error read data stream:", err)
			}
			err = c.Close()
			if err != nil {
				fmt.Println("Error close socket:", err)
				log.Println("Error close socket:", err)
			}
			break
		}
		// handle data stream
		str := strings.TrimSpace(string(buf[0:n]))
		t := time.Now()
		fmt.Println("[" + c.RemoteAddr().String() + "] ", strconv.Itoa(t.Hour()) + ":" + strconv.Itoa(t.Minute()) + ":" + strconv.Itoa(t.Second()))
		fmt.Println("Remote->Local:", str)
	}
}
