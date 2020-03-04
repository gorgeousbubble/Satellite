package pack

import (
	"testing"
)

// TestHMAC_SHA1 function
func TestHMAC_SHA1(t *testing.T) {
	src := "hello,world!"
	key := "Alopex6414"
	dest := "e6a6dd9149277e8a5a189e2041a4e5ef50b5bca0"
	r := HMAC_SHA1(src, key)
	if r != dest {
		t.Fatal("Error HMAC SHA1:")
	}
}

// TestHMAC_SHA256 function
func TestHMAC_SHA256(t *testing.T) {
	src := "hello,world!"
	key := "Alopex6414"
	dest := "8cd5d33d63218ab784b2ad2585adaee7f33493f12908b7b7b8d67a842905f45d"
	r := HMAC_SHA256(src, key)
	if r != dest {
		t.Fatal("Error HMAC SHA256:")
	}
}

// TestHMAC_SHA512 function
func TestHMAC_SHA512(t *testing.T) {
	src := "hello,world!"
	key := "Alopex6414"
	dest := "3f84207c9a08114617659afe683589ceade646f2e1d9a20d96c916d64a1e1417633b02b1e59525821008e54d27e7776a5b3907ff0d424ec82add59c1c35c36c4"
	r := HMAC_SHA512(src, key)
	if r != dest {
		t.Fatal("Error HMAC SHA512:")
	}
}

// BenchmarkHMAC_SHA1 function
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

// BenchmarkHMAC_SHA256 function
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

// BenchmarkHMAC_SHA512 function
func BenchmarkHMAC_SHA512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		key := "Alopex6414"
		dest := "3f84207c9a08114617659afe683589ceade646f2e1d9a20d96c916d64a1e1417633b02b1e59525821008e54d27e7776a5b3907ff0d424ec82add59c1c35c36c4"
		r := HMAC_SHA512(src, key)
		if r != dest {
			b.Fatal("Error HMAC SHA512:")
		}
	}
}
