package pack

import (
	"crypto/sha1"
	"io/ioutil"
	"sync"
	"testing"
)

func TestSHA1EncryptGo(t *testing.T) {
	var wg sync.WaitGroup
	var r [sha1.Size]byte
	src := []byte("hello,world!")
	wg.Add(1)
	go SHA1EncryptGo(src, &r, &wg)
	wg.Wait()
	err := ioutil.WriteFile("../test/data/pack/file_sha1.txt", r[:(sha1.Size)], 0644)
	if err != nil {
		t.Fatal("Error Write SHA1 Encrypt:", err)
	}
}

func TestSHA1Encrypt(t *testing.T) {
	src := []byte("hello,world!")
	r := SHA1Encrypt(src)
	err := ioutil.WriteFile("../test/data/pack/file_sha1.txt", r[:(sha1.Size)], 0644)
	if err != nil {
		t.Fatal("Error Write SHA1 Encrypt:", err)
	}
}

func BenchmarkSHA1EncryptGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r [sha1.Size]byte
		src := []byte("hello,world!")
		wg.Add(1)
		go SHA1EncryptGo(src, &r, &wg)
		wg.Wait()
		err := ioutil.WriteFile("../test/data/pack/file_sha1.txt", r[:(sha1.Size)], 0644)
		if err != nil {
			b.Fatal("Error Write SHA1 Encrypt:", err)
		}
	}
}

func BenchmarkSHA1Encrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte("hello,world!")
		r := SHA1Encrypt(src)
		err := ioutil.WriteFile("../test/data/pack/file_sha1.txt", r[:(sha1.Size)], 0644)
		if err != nil {
			b.Fatal("Error Write SHA1 Encrypt:", err)
		}
	}
}
