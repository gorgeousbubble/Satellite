package pack

import (
	"crypto/sha512"
	"io/ioutil"
	"sync"
	"testing"
)

// TestSHA512Check function
func TestSHA512Check(t *testing.T) {
	src := "hello,world!"
	dest := "fa9edcfdaab7a4165f8d2593f04077d46aca957493e7e181ca479582d519a299d967305294d5d46fb5e0158240441b94cd96510c2311bdc86870e5ebf3efe60c"
	r := SHA512Check(src, dest)
	if !r {
		t.Fatal("Error Check SHA512:")
	}
}

// TestSHA512Check2 function
func TestSHA512Check2(t *testing.T) {
	src := "Nice to meet you~"
	dest := "fa9edcfdaab7a4165f8d2593f04077d46aca957493e7e181ca479582d519a299d967305294d5d46fb5e0158240441b94cd96510c2311bdc86870e5ebf3efe60c"
	r := SHA512Check(src, dest)
	if r {
		t.Fatal("Error Check SHA512:")
	}
}

// TestSHA512Encode function
func TestSHA512Encode(t *testing.T) {
	src := "hello,world!"
	r := SHA512Encode(src)
	err := ioutil.WriteFile("../test/data/pack/file_sha512.txt", []byte(r), 0644)
	if err != nil {
		t.Fatal("Error Write SHA512 Encode:", err)
	}
}

// TestSHA512EncryptGo function
func TestSHA512EncryptGo(t *testing.T) {
	var wg sync.WaitGroup
	var r [sha512.Size]byte
	src := []byte("hello,world!")
	wg.Add(1)
	go SHA512EncryptGo(src, &r, &wg)
	wg.Wait()
	err := ioutil.WriteFile("../test/data/pack/file_sha512.txt", r[:sha512.Size], 0644)
	if err != nil {
		t.Fatal("Error Write SHA512 Encrypt:", err)
	}
}

// TestSHA512Encrypt
func TestSHA512Encrypt(t *testing.T) {
	src := []byte("hello,world!")
	r := SHA512Encrypt(src)
	err := ioutil.WriteFile("../test/data/pack/file_sha512.txt", r[:sha512.Size], 0644)
	if err != nil {
		t.Fatal("Error Write SHA512 Encrypt:", err)
	}
}

func BenchmarkSHA512Check(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		dest := "fa9edcfdaab7a4165f8d2593f04077d46aca957493e7e181ca479582d519a299d967305294d5d46fb5e0158240441b94cd96510c2311bdc86870e5ebf3efe60c"
		r := SHA512Check(src, dest)
		if !r {
			b.Fatal("Error Check SHA512:")
		}
	}
}

func BenchmarkSHA512Check2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "Nice to meet you~"
		dest := "fa9edcfdaab7a4165f8d2593f04077d46aca957493e7e181ca479582d519a299d967305294d5d46fb5e0158240441b94cd96510c2311bdc86870e5ebf3efe60c"
		r := SHA512Check(src, dest)
		if r {
			b.Fatal("Error Check SHA512:")
		}
	}
}

func BenchmarkSHA512Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		r := SHA512Encode(src)
		err := ioutil.WriteFile("../test/data/pack/file_sha512.txt", []byte(r), 0644)
		if err != nil {
			b.Fatal("Error Write SHA512 Encode:", err)
		}
	}
}

func BenchmarkSHA512EncryptGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r [sha512.Size]byte
		src := []byte("hello,world!")
		wg.Add(1)
		go SHA512EncryptGo(src, &r, &wg)
		wg.Wait()
		err := ioutil.WriteFile("../test/data/pack/file_sha512.txt", r[:sha512.Size], 0644)
		if err != nil {
			b.Fatal("Error Write SHA512 Encrypt:", err)
		}
	}
}

func BenchmarkSHA512Encrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte("hello,world!")
		r := SHA512Encrypt(src)
		err := ioutil.WriteFile("../test/data/pack/file_sha512.txt", r[:sha512.Size], 0644)
		if err != nil {
			b.Fatal("Error Write SHA512 Encrypt:", err)
		}
	}
}
