package pack

import (
	"io/ioutil"
	"sync"
	"testing"
)

func TestPackBase64OneGo(t *testing.T) {
	var wg sync.WaitGroup
	var r string
	src := "../test/data/pack/file.txt"
	wg.Add(1)
	go PackBase64OneGo(src, &r, &wg)
	wg.Wait()
	err := ioutil.WriteFile("../test/data/pack/file_base64.txt", []byte(r), 0644)
	if err != nil {
		t.Fatal("Error Write Base64 One:", err)
	}
}

func TestPackBase64One(t *testing.T) {
	src := "../test/data/pack/file.txt"
	r, err := PackBase64One(src)
	if err != nil {
		t.Fatal("Error Pack Base64 One:", err)
	}
	err = ioutil.WriteFile("../test/data/pack/file_base64.txt", []byte(r), 0644)
	if err != nil {
		t.Fatal("Error Write Base64 One:", err)
	}
}

func TestBase64EncryptGo(t *testing.T) {
	var wg sync.WaitGroup
	var r string
	src := "hello,world!"
	dest := "aGVsbG8sd29ybGQh"
	wg.Add(1)
	go Base64EncryptGo(src, &r, &wg)
	wg.Wait()
	if r != dest {
		t.Fatal("Error Encrypt Base64.")
	}
	err := ioutil.WriteFile("../test/data/pack/file_base64.txt", []byte(r), 0644)
	if err != nil {
		t.Fatal("Error Write Base64 One:", err)
	}
}

func TestBase64Encrypt(t *testing.T) {
	src := "hello,world!"
	dest := "aGVsbG8sd29ybGQh"
	r := Base64Encrypt(src)
	if r != dest {
		t.Fatal("Error Encrypt Base64.")
	}
	err := ioutil.WriteFile("../test/data/pack/file_base64.txt", []byte(r), 0644)
	if err != nil {
		t.Fatal("Error Write Base64 One:", err)
	}
}

func BenchmarkPackBase64OneGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r string
		src := "../test/data/pack/file.txt"
		wg.Add(1)
		go PackBase64OneGo(src, &r, &wg)
		wg.Wait()
		err := ioutil.WriteFile("../test/data/pack/file_base64.txt", []byte(r), 0644)
		if err != nil {
			b.Fatal("Error Write Base64 One:", err)
		}
	}
}

func BenchmarkPackBase64One(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/pack/file.txt"
		r, err := PackBase64One(src)
		if err != nil {
			b.Fatal("Error Pack Base64 One:", err)
		}
		err = ioutil.WriteFile("../test/data/pack/file_base64.txt", []byte(r), 0644)
		if err != nil {
			b.Fatal("Error Write Base64 One:", err)
		}
	}
}

func BenchmarkBase64EncryptGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r string
		src := "hello,world!"
		dest := "aGVsbG8sd29ybGQh"
		wg.Add(1)
		go Base64EncryptGo(src, &r, &wg)
		wg.Wait()
		if r != dest {
			b.Fatal("Error Encrypt Base64.")
		}
		err := ioutil.WriteFile("../test/data/pack/file_base64.txt", []byte(r), 0644)
		if err != nil {
			b.Fatal("Error Write Base64 One:", err)
		}
	}
}

func BenchmarkBase64Encrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		dest := "aGVsbG8sd29ybGQh"
		r := Base64Encrypt(src)
		if r != dest {
			b.Fatal("Error Encrypt Base64.")
		}
		err := ioutil.WriteFile("../test/data/pack/file_base64.txt", []byte(r), 0644)
		if err != nil {
			b.Fatal("Error Write Base64 One:", err)
		}
	}
}