package pack

import (
	"crypto/md5"
	"io/ioutil"
	"sync"
	"testing"
)

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