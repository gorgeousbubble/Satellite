package nets

import "testing"

func TestStartTcpServer(t *testing.T) {
	ip := "127.0.0.1"
	port := "11514"
	// skip this test case
	t.Skip("IGNORE: TestStartTcpServer")
	// start tcp server goroutine
	go StartTcpServer(ip, port)
	// start tcp client...
	StartTcpClient(ip, port)
}
