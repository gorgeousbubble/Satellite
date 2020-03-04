package nets

import "testing"

func TestStartUdpServer(t *testing.T) {
	ip := "127.0.0.1"
	port := "12514"
	// skip this test case
	t.Skip("IGNORE: TestStartUdpServer")
	// start udp server goroutine
	go StartUdpServer(ip, port)
	// start udp client...
	StartUdpClient(ip, port)
}
