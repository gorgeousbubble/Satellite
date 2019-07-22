package pack

import (
	"io/ioutil"
	"testing"
)

func TestSHA1Encrypt(t *testing.T) {
	src := []byte("hello,world!")
	r := SHA1Encrypt(src)
	err := ioutil.WriteFile("../test/data/pack/file_sha1.txt", r[:19], 0644)
	if err != nil {
		t.Fatal("Error Write SHA1 Encrypt:", err)
	}
}

func BenchmarkSHA1Encrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte("hello,world!")
		r := SHA1Encrypt(src)
		err := ioutil.WriteFile("../test/data/pack/file_sha1.txt", r[:19], 0644)
		if err != nil {
			b.Fatal("Error Write SHA1 Encrypt:", err)
		}
	}
}
