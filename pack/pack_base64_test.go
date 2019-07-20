package pack

import (
	"io/ioutil"
	"testing"
)

func TestBase64Encrypt(t *testing.T) {
	src := "hello,world!"
	r := Base64Encrypt(src)
	err := ioutil.WriteFile("../test/data/file_base64.txt", []byte(r), 0644)
	if err != nil {
		t.Fatal("Error Write Base64 One:", err)
	}
}

func BenchmarkBase64Encrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		r := Base64Encrypt(src)
		err := ioutil.WriteFile("test/file_base64.txt", []byte(r), 0644)
		if err != nil {
			b.Fatal("Error Write Base64 One:", err)
		}
	}
}