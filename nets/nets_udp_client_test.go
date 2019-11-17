package nets

import "testing"

func TestStartUdpClient(t *testing.T) {
	ip := "127.0.0.1"
	port := "6003"
	t.Skip("Skip start udp client...")
	// start udp server goroutine
	go StartUdpServer(ip, port)
	// start udp client...
	StartUdpClient(ip, port)
}
