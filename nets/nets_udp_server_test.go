package nets

import "testing"

func TestStartUdpServer(t *testing.T) {
	ip := "127.0.0.1"
	port := "6002"
	t.Skip("Skip start udp server...")
	// start udp server goroutine
	go StartUdpServer(ip, port)
	// start udp client...
	StartUdpClient(ip, port)
}
