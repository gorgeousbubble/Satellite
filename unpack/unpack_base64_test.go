package unpack

import (
	"io/ioutil"
	"testing"
)

func TestBase64Decrypt(t *testing.T) {
	src := "aGVsbG8sd29ybGQh"
	dest := "hello,world!"
	r := Base64Decrypt(src)
	if r != dest {
		t.Errorf("Error Decrypt Base64.")
	}
	err := ioutil.WriteFile("../test/data/unpack/file.txt", []byte(r), 0644)
	if err != nil {
		t.Fatal("Error Write Base64 One:", err)
	}
}

func BenchmarkBase64Decrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "aGVsbG8sd29ybGQh"
		dest := "hello,world!"
		r := Base64Decrypt(src)
		if r != dest {
			b.Errorf("Error Decrypt Base64.")
		}
		err := ioutil.WriteFile("../test/data/unpack/file.txt", []byte(r), 0644)
		if err != nil {
			b.Fatal("Error Write Base64 One:", err)
		}
	}
}