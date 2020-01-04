package nets

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHttpGet(t *testing.T) {
	url := "http://www.baidu.com"
	r, err := HttpGet(url)
	if err != nil {
		t.Fatal("Error http get:", err)
	}
	fmt.Println(string(r))
}

func BenchmarkHttpGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://www.baidu.com"
		_, err := HttpGet(url)
		if err != nil {
			b.Fatal("Error http get:", err)
		}
	}
}

func TestHttpHead(t *testing.T) {
	url := "http://www.baidu.com"
	code, err := HttpHead(url)
	if err != nil {
		t.Fatal("Error http head:", err)
	}
	if code != http.StatusOK {
		t.Fatal("Error http head response status code")
	}
}

func BenchmarkHttpHead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://www.baidu.com"
		code, err := HttpHead(url)
		if err != nil {
			b.Fatal("Error http head:", err)
		}
		if code != http.StatusOK {
			b.Fatal("Error http head response status code")
		}
	}
}
