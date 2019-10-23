package nets

import (
	"testing"
)

func TestStartTcpServer(t *testing.T) {
	ip := "127.0.0.1"
	port := "6000"
	// start tcp server goroutine
	go StartTcpServer(ip, port)
	// start tcp client...
	StartTcpClient(ip, port)
}
