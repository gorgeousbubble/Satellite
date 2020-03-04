package nets

import (
	"net/http"
	"testing"
)

func TestStartFtpServer(t *testing.T) {
	ip := "127.0.0.1"
	port := "10021"
	// skip this test case
	t.Skip("IGNORE: TestStartFtpServer")
	// start ftp server
	go StartFtpServer(ip, port)
	// start ftp client
	r, err := http.Get("http://" + ip + ":" + port)
	if err != nil {
		t.Errorf("Error GET ftp server url:%v", err)
	}
	if r.StatusCode != http.StatusOK {
		t.Errorf("Error GET ftp server response")
	}
}
