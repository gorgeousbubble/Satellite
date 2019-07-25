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

func TestHMAC_SHA256(t *testing.T) {
	src := "hello,world!"
	key := "Alopex6414"
	dest := "8cd5d33d63218ab784b2ad2585adaee7f33493f12908b7b7b8d67a842905f45d"
	r := HMAC_SHA256(src, key)
	if r != dest {
		t.Fatal("Error HMAC SHA256:")
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

func BenchmarkHMAC_SHA256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		key := "Alopex6414"
		dest := "8cd5d33d63218ab784b2ad2585adaee7f33493f12908b7b7b8d67a842905f45d"
		r := HMAC_SHA256(src, key)
		if r != dest {
			b.Fatal("Error HMAC SHA256:")
		}
	}
}
