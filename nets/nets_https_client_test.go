package nets

import (
	"fmt"
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
