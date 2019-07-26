package pack

import (
	"io/ioutil"
	"sync"
	"testing"
)

func TestDESEncryptGo(t *testing.T) {
	var wg sync.WaitGroup
	var r []byte
	src := []byte("hello,world!")
	key := []byte("hyacinth")
	wg.Add(1)
	go DESEncryptGo(src, key, &r, &wg)
	wg.Wait()
	err := ioutil.WriteFile("../test/data/pack/file_des.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write DES Encrypt:", err)
	}
}

func TestDESEncrypt(t *testing.T) {
	src := []byte("hello,world!")
	key := []byte("hyacinth")
	r, err := DESEncrypt(src, key)
	if err != nil {
		t.Fatal("Error DES Encrypt:", err)
	}
	err = ioutil.WriteFile("../test/data/pack/file_des.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write DES Encrypt:", err)
	}
}

func BenchmarkDESEncryptGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r []byte
		src := []byte("hello,world!")
		key := []byte("hyacinth")
		wg.Add(1)
		go DESEncryptGo(src, key, &r, &wg)
		wg.Wait()
		err := ioutil.WriteFile("../test/data/pack/file_des.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write DES Encrypt:", err)
		}
	}
}

func BenchmarkDESEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte("hello,world!")
		key := []byte("hyacinth")
		r, err := DESEncrypt(src, key)
		if err != nil {
			b.Fatal("Error DES Encrypt:", err)
		}
		err = ioutil.WriteFile("../test/data/pack/file_des.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write DES Encrypt:", err)
		}
	}
}
