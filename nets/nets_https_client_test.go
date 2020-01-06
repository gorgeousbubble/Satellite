package nets

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHttpsGet(t *testing.T) {
	url := "https://github.com"
	r, err := HttpsGet(url)
	if err != nil {
		t.Fatal("Error https get:", err)
	}
	fmt.Println(string(r))
}

func BenchmarkHttpsGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "https://github.com"
		_, err := HttpsGet(url)
		if err != nil {
			b.Fatal("Error https get:", err)
		}
	}
}

func TestHttpsHead(t *testing.T) {
	url := "https://github.com"
	code, err := HttpsHead(url)
	if err != nil {
		t.Fatal("Error https head:", err)
	}
	if code != http.StatusOK {
		t.Fatal("Error https head response status code")
	}
}
