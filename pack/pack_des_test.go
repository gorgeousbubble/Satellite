package pack

import (
	"io/ioutil"
	"sync"
	"testing"
)

func TestPackDESOne(t *testing.T) {
	src := "../test/data/pack/file.txt"
	r, err := PackDESOne(src)
	if err != nil {
		t.Fatal("Error Pack DES One:", err)
	}
	err = ioutil.WriteFile("../test/data/pack/file_des.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write DES One:", err)
	}
}

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

func BenchmarkPackDESOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/pack/file.txt"
		r, err := PackDESOne(src)
		if err != nil {
			b.Fatal("Error Pack DES One:", err)
		}
		err = ioutil.WriteFile("../test/data/pack/file_des.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write DES One:", err)
		}
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
