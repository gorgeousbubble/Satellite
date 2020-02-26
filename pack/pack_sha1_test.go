package pack

import (
	"crypto/sha1"
	"io/ioutil"
	"sync"
	"testing"
)

// TestSHA1Check function
func TestSHA1Check(t *testing.T) {
	src := "hello,world!"
	dest := "4518135c05e0706c0a34168996517bb3f28d94b5"
	r := SHA1Check(src, dest)
	if !r {
		t.Fatal("Error Check SHA1:")
	}
}

// TestSHA1Check2 function
func TestSHA1Check2(t *testing.T) {
	src := "Nice to meet you~"
	dest := "4518135c05e0706c0a34168996517bb3f28d94b5"
	r := SHA1Check(src, dest)
	if r {
		t.Fatal("Error Check SHA1:")
	}
}

// TestSHA1Encode function
func TestSHA1Encode(t *testing.T) {
	src := "hello,world!"
	r := SHA1Encode(src)
	err := ioutil.WriteFile("../test/data/pack/file_sha1.txt", []byte(r), 0644)
	if err != nil {
		t.Fatal("Error Write SHA1 Encode:", err)
	}
}

// TestSHA1EncryptGo function
func TestSHA1EncryptGo(t *testing.T) {
	var wg sync.WaitGroup
	var r [sha1.Size]byte
	src := []byte("hello,world!")
	wg.Add(1)
	go SHA1EncryptGo(src, &r, &wg)
	wg.Wait()
	err := ioutil.WriteFile("../test/data/pack/file_sha1.txt", r[:sha1.Size], 0644)
	if err != nil {
		t.Fatal("Error Write SHA1 Encrypt:", err)
	}
}

// TestSHA1Encrypt function
func TestSHA1Encrypt(t *testing.T) {
	src := []byte("hello,world!")
	r := SHA1Encrypt(src)
	err := ioutil.WriteFile("../test/data/pack/file_sha1.txt", r[:sha1.Size], 0644)
	if err != nil {
		t.Fatal("Error Write SHA1 Encrypt:", err)
	}
}

func BenchmarkSHA1Check(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		dest := "4518135c05e0706c0a34168996517bb3f28d94b5"
		r := SHA1Check(src, dest)
		if !r {
			b.Fatal("Error Check SHA1:")
		}
	}
}

func BenchmarkSHA1Check2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "Nice to meet you~"
		dest := "4518135c05e0706c0a34168996517bb3f28d94b5"
		r := SHA1Check(src, dest)
		if r {
			b.Fatal("Error Check SHA1:")
		}
	}
}

func BenchmarkSHA1Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		r := SHA1Encode(src)
		err := ioutil.WriteFile("../test/data/pack/file_sha1.txt", []byte(r), 0644)
		if err != nil {
			b.Fatal("Error Write SHA1 Encode:", err)
		}
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
		err := ioutil.WriteFile("../test/data/pack/file_sha1.txt", r[:sha1.Size], 0644)
		if err != nil {
			b.Fatal("Error Write SHA1 Encrypt:", err)
		}
	}
}

func BenchmarkSHA1Encrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte("hello,world!")
		r := SHA1Encrypt(src)
		err := ioutil.WriteFile("../test/data/pack/file_sha1.txt", r[:sha1.Size], 0644)
		if err != nil {
			b.Fatal("Error Write SHA1 Encrypt:", err)
		}
	}
}
