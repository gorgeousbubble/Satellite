package pack

import (
	"crypto/sha512"
	"io/ioutil"
	"testing"
)

func TestSHA512Encrypt(t *testing.T) {
	src := []byte("hello,world!")
	r := SHA512Encrypt(src)
	err := ioutil.WriteFile("../test/data/pack/file_sha512.txt", r[:sha512.Size], 0644)
	if err != nil {
		t.Fatal("Error Write SHA512 Encrypt:", err)
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
