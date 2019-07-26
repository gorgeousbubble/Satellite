package pack

import (
	"io/ioutil"
	"sync"
	"testing"
)

func TestPackDESOneGo(t *testing.T) {
	var wg sync.WaitGroup
	var r []byte
	src := "../test/data/pack/file.txt"
	wg.Add(1)
	err := PackDESOneGo(src, &r, &wg)
	if err != nil {
		t.Fatal("Error Pack DES One Go:", err)
	}
	wg.Wait()
	err = ioutil.WriteFile("../test/data/pack/file_des.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write DES One Go:", err)
	}
}

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

func BenchmarkPackDESOneGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r []byte
		src := "../test/data/pack/file.txt"
		wg.Add(1)
		err := PackDESOneGo(src, &r, &wg)
		if err != nil {
			b.Fatal("Error Pack DES One Go:", err)
		}
		wg.Wait()
		err = ioutil.WriteFile("../test/data/pack/file_des.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write DES One Go:", err)
		}
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
