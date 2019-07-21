package pack

import (
	"io/ioutil"
	. "satellite/utils"
	"sync"
	"testing"
)

func TestPackRSA(t *testing.T) {
	src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
	dest := "../test/data/pack/file_rsa.txt"
	err := PackRSA(src, dest)
	if err != nil {
		t.Fatal("Error Pack RSA:", err)
	}
}

func TestPackRSAOneGo(t *testing.T) {
	var wg sync.WaitGroup
	var r []byte
	src := "../test/data/pack/file.txt"
	wg.Add(1)
	err := PackRSAOneGo(src, &r, &wg)
	if err != nil {
		t.Fatal("Error Pack RSA One Go:", err)
	}
	wg.Wait()
	err = ioutil.WriteFile("../test/data/pack/file_rsa.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write RSA One Go:", err)
	}
}

func TestPackRSAOne(t *testing.T) {
	src := "../test/data/pack/file.txt"
	r, err := PackRSAOne(src)
	if err != nil {
		t.Fatal("Error Pack RSA One:", err)
	}
	err = ioutil.WriteFile("../test/data/pack/file_rsa.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write RSA One:", err)
	}
}

func TestRSAEncryptGo(t *testing.T) {
	var wg sync.WaitGroup
	var r []byte
	src := []byte("hello,world!")
	wg.Add(1)
	go RSAEncryptGo(src, ConstRSAPublicKey, &r, &wg)
	wg.Wait()
	err := ioutil.WriteFile("../test/data/pack/file_rsa.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write RSA Encrypt:", err)
	}
}

func TestRSAEncrypt(t *testing.T) {
	src := []byte("hello,world!")
	r, err := RSAEncrypt(src, ConstRSAPublicKey)
	if err != nil {
		t.Fatal("Error RSA Encrypt:", err)
	}
	err = ioutil.WriteFile("../test/data/pack/file_rsa.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write RSA Encrypt:", err)
	}
}

func BenchmarkPackRSA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []string{"../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"}
		dest := "../test/data/pack/file_rsa.txt"
		err := PackRSA(src, dest)
		if err != nil {
			b.Fatal("Error Pack RSA:", err)
		}
	}
}

func BenchmarkPackRSAOneGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r []byte
		src := "../test/data/pack/file.txt"
		wg.Add(1)
		err := PackRSAOneGo(src, &r, &wg)
		if err != nil {
			b.Fatal("Error Pack RSA One Go:", err)
		}
		wg.Wait()
		err = ioutil.WriteFile("../test/data/pack/file_rsa.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write RSA One Go:", err)
		}
	}
}

func BenchmarkPackRSAOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/pack/file.txt"
		r, err := PackRSAOne(src)
		if err != nil {
			b.Fatal("Error Pack RSA One:", err)
		}
		err = ioutil.WriteFile("../test/data/pack/file_rsa.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write RSA One:", err)
		}
	}
}

func BenchmarkRSAEncryptGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r []byte
		src := []byte("hello,world!")
		wg.Add(1)
		go RSAEncryptGo(src, ConstRSAPublicKey, &r, &wg)
		wg.Wait()
		err := ioutil.WriteFile("../test/data/pack/file_rsa.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write RSA Encrypt:", err)
		}
	}
}

func BenchmarkRSAEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte("hello,world!")
		r, err := RSAEncrypt(src, ConstRSAPublicKey)
		if err != nil {
			b.Fatal("Error RSA Encrypt:", err)
		}
		err = ioutil.WriteFile("../test/data/pack/file_rsa.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write RSA Encrypt:", err)
		}
	}
}
