package pack

import (
	"crypto/sha256"
	"io/ioutil"
	"sync"
	"testing"
)

func TestSHA256EncryptGo(t *testing.T) {
	var wg sync.WaitGroup
	var r [sha256.Size]byte
	src := []byte("hello,world!")
	wg.Add(1)
	go SHA256EncryptGo(src, &r, &wg)
	wg.Wait()
	err := ioutil.WriteFile("../test/data/pack/file_sha256.txt", r[:(sha256.Size-1)], 0644)
	if err != nil {
		t.Fatal("Error Write SHA256 Encrypt:", err)
	}
}

func TestSHA256Encrypt(t *testing.T) {
	src := []byte("hello,world!")
	r := SHA256Encrypt(src)
	err := ioutil.WriteFile("../test/data/pack/file_sha256.txt", r[:(sha256.Size-1)], 0644)
	if err != nil {
		t.Fatal("Error Write SHA256 Encrypt:", err)
	}
}

func BenchmarkSHA256EncryptGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r [sha256.Size]byte
		src := []byte("hello,world!")
		wg.Add(1)
		go SHA256EncryptGo(src, &r, &wg)
		wg.Wait()
		err := ioutil.WriteFile("../test/data/pack/file_sha256.txt", r[:(sha256.Size-1)], 0644)
		if err != nil {
			b.Fatal("Error Write SHA256 Encrypt:", err)
		}
	}
}

func BenchmarkSHA256Encrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte("hello,world!")
		r := SHA256Encrypt(src)
		err := ioutil.WriteFile("../test/data/pack/file_sha256.txt", r[:(sha256.Size-1)], 0644)
		if err != nil {
			b.Fatal("Error Write SHA256 Encrypt:", err)
		}
	}
}
