package nets

import (
	"fmt"
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
