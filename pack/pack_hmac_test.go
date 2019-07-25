package pack

import (
	"testing"
)

func TestHMAC_SHA1(t *testing.T) {
	src := "hello,world!"
	key := "Alopex6414"
	dest := "e6a6dd9149277e8a5a189e2041a4e5ef50b5bca0"
	r := HMAC_SHA1(src, key)
	if r != dest {
		t.Fatal("Error HMAC SHA1:")
	}
}

func BenchmarkHMAC_SHA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		key := "Alopex6414"
		dest := "e6a6dd9149277e8a5a189e2041a4e5ef50b5bca0"
		r := HMAC_SHA1(src, key)
		if r != dest {
			b.Fatal("Error HMAC SHA1:")
		}
	}
}
