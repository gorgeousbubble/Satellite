package nets

import (
	"net"
	"testing"
)

func TestStartTcpServer(t *testing.T) {
	ip := "127.0.0.1"
	port := "6000"
	// start tcp server goroutine
	go StartTcpServer(ip, port)
	// connect to tcp server
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		t.Errorf("Error connect to server: %v", err)
	}
	if conn == nil {
		t.Error("Invalid socket connect.")
	}
	_, err = conn.Write([]byte("hello,world!"))
	if err != nil {
		t.Errorf("Error write data stream: %v", err)
	}
	conn.Close()
}
