package pack

import (
	"crypto/md5"
	"io/ioutil"
	"sync"
	"testing"
)

func TestMD5Check(t *testing.T) {
	src := "hello,world!"
	dest := "c0e84e870874dd37ed0d164c7986f03a"
	b := MD5Check(src, dest)
	if !b {
		t.Fatal("Error Check MD5:")
	}
}

func TestMD5Check2(t *testing.T) {
	src := "Nice to meet you~"
	dest := "c0e84e870874dd37ed0d164c7986f03a"
	b := MD5Check(src, dest)
	if b {
		t.Fatal("Error Check MD5:")
	}
}

func TestMD5Encode(t *testing.T) {
	src := "hello,world!"
	r := MD5Encode(src)
	err := ioutil.WriteFile("../test/data/pack/file_md5.txt", []byte(r), 0644)
	if err != nil {
		t.Fatal("Error Write MD5 Encode:", err)
	}
}

func TestMD5EncryptGo(t *testing.T) {
	var wg sync.WaitGroup
	var r [md5.Size]byte
	src := []byte("hello,world!")
	wg.Add(1)
	go MD5EncryptGo(src, &r, &wg)
	wg.Wait()
	err := ioutil.WriteFile("../test/data/pack/file_md5.txt", r[:md5.Size], 0644)
	if err != nil {
		t.Fatal("Error Write MD5 Encrypt:", err)
	}
}

func TestMD5Encrypt(t *testing.T) {
	src := []byte("hello,world!")
	r := MD5Encrypt(src)
	err := ioutil.WriteFile("../test/data/pack/file_md5.txt", r[:md5.Size], 0644)
	if err != nil {
		t.Fatal("Error Write MD5 Encrypt:", err)
	}
}

func BenchmarkMD5Check(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		dest := "c0e84e870874dd37ed0d164c7986f03a"
		f := MD5Check(src, dest)
		if !f {
			b.Fatal("Error Check MD5:")
		}
	}
}

func BenchmarkMD5Check2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "Nice to meet you~"
		dest := "c0e84e870874dd37ed0d164c7986f03a"
		f := MD5Check(src, dest)
		if f {
			b.Fatal("Error Check MD5:")
		}
	}
}

func BenchmarkMD5Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		r := MD5Encode(src)
		err := ioutil.WriteFile("../test/data/pack/file_md5.txt", []byte(r), 0644)
		if err != nil {
			b.Fatal("Error Write MD5 Encode:", err)
		}
	}
}

func BenchmarkMD5EncryptGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r [md5.Size]byte
		src := []byte("hello,world!")
		wg.Add(1)
		go MD5EncryptGo(src, &r, &wg)
		wg.Wait()
		err := ioutil.WriteFile("../test/data/pack/file_md5.txt", r[:md5.Size], 0644)
		if err != nil {
			b.Fatal("Error Write MD5 Encrypt:", err)
		}
	}
}

func BenchmarkMD5Encrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte("hello,world!")
		r := MD5Encrypt(src)
		err := ioutil.WriteFile("../test/data/pack/file_md5.txt", r[:md5.Size], 0644)
		if err != nil {
			b.Fatal("Error Write MD5 Encrypt:", err)
		}
	}
}
