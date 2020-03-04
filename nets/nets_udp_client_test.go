package nets

import "testing"

func TestStartUdpClient(t *testing.T) {
	ip := "127.0.0.1"
	port := "12515"
	// skip this test case
	t.Skip("IGNORE: TestStartUdpClient")
	// start udp server goroutine
	go StartUdpServer(ip, port)
	// start udp client...
	StartUdpClient(ip, port)
}
