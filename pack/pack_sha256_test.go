package pack

import (
	"crypto/sha256"
	"io/ioutil"
	"sync"
	"testing"
)

// TestSHA256Check function
func TestSHA256Check(t *testing.T) {
	src := "hello,world!"
	dest := "ec1e0bd875226943ad0e8877bdba4ca449c4cb8591a5363921c9f1ee20084c34"
	r := SHA256Check(src, dest)
	if !r {
		t.Fatal("Error Check SHA256:")
	}
}

// TestSHA256Check2 function
func TestSHA256Check2(t *testing.T) {
	src := "Nice to meet you~"
	dest := "ec1e0bd875226943ad0e8877bdba4ca449c4cb8591a5363921c9f1ee20084c34"
	r := SHA256Check(src, dest)
	if r {
		t.Fatal("Error Check SHA256:")
	}
}

// TestSHA256Encode function
func TestSHA256Encode(t *testing.T) {
	src := "hello,world!"
	r := SHA256Encode(src)
	err := ioutil.WriteFile("../test/data/pack/file_sha256.txt", []byte(r), 0644)
	if err != nil {
		t.Fatal("Error Write SHA256 Encode:", err)
	}
}

// TestSHA256EncryptGo function
func TestSHA256EncryptGo(t *testing.T) {
	var wg sync.WaitGroup
	var r [sha256.Size]byte
	src := []byte("hello,world!")
	wg.Add(1)
	go SHA256EncryptGo(src, &r, &wg)
	wg.Wait()
	err := ioutil.WriteFile("../test/data/pack/file_sha256.txt", r[:sha256.Size], 0644)
	if err != nil {
		t.Fatal("Error Write SHA256 Encrypt:", err)
	}
}

// TestSHA256Encrypt function
func TestSHA256Encrypt(t *testing.T) {
	src := []byte("hello,world!")
	r := SHA256Encrypt(src)
	err := ioutil.WriteFile("../test/data/pack/file_sha256.txt", r[:sha256.Size], 0644)
	if err != nil {
		t.Fatal("Error Write SHA256 Encrypt:", err)
	}
}

// BenchmarkSHA256Check function
func BenchmarkSHA256Check(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		dest := "ec1e0bd875226943ad0e8877bdba4ca449c4cb8591a5363921c9f1ee20084c34"
		r := SHA256Check(src, dest)
		if !r {
			b.Fatal("Error Check SHA256:")
		}
	}
}

// BenchmarkSHA256Check2 function
func BenchmarkSHA256Check2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "Nice to meet you~"
		dest := "ec1e0bd875226943ad0e8877bdba4ca449c4cb8591a5363921c9f1ee20084c34"
		r := SHA256Check(src, dest)
		if r {
			b.Fatal("Error Check SHA256:")
		}
	}
}

// BenchmarkSHA256Encode function
func BenchmarkSHA256Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		r := SHA256Encode(src)
		err := ioutil.WriteFile("../test/data/pack/file_sha256.txt", []byte(r), 0644)
		if err != nil {
			b.Fatal("Error Write SHA256 Encode:", err)
		}
	}
}

// BenchmarkSHA256EncryptGo function
func BenchmarkSHA256EncryptGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r [sha256.Size]byte
		src := []byte("hello,world!")
		wg.Add(1)
		go SHA256EncryptGo(src, &r, &wg)
		wg.Wait()
		err := ioutil.WriteFile("../test/data/pack/file_sha256.txt", r[:sha256.Size], 0644)
		if err != nil {
			b.Fatal("Error Write SHA256 Encrypt:", err)
		}
	}
}

// BenchmarkSHA256Encrypt function
func BenchmarkSHA256Encrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte("hello,world!")
		r := SHA256Encrypt(src)
		err := ioutil.WriteFile("../test/data/pack/file_sha256.txt", r[:sha256.Size], 0644)
		if err != nil {
			b.Fatal("Error Write SHA256 Encrypt:", err)
		}
	}
}
