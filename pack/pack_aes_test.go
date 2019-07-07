package pack

import (
	"io/ioutil"
	. "satellite/utils"
	"sync"
	"testing"
)

func TestPackAES(t *testing.T) {
	src := []string{"test/file_1.txt", "test/file_2.txt", "test/file_3.txt", "test/file_4.txt", "test/file_5.txt"}
	dest := "test/file_aes.txt"
	err := PackAES(src, dest)
	if err != nil {
		t.Fatal("Error Pack AES:", err)
	}
}

func TestPackAESOneGo(t *testing.T) {
	var wg sync.WaitGroup
	var r []byte
	src := "test/file.txt"
	wg.Add(1)
	err := PackAESOneGo(src, &r, &wg)
	if err != nil {
		t.Fatal("Error Pack AES One Go:", err)
	}
	wg.Wait()
	err = ioutil.WriteFile("test/file_aes.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write AES One Go:", err)
	}
}

func TestPackAESOne(t *testing.T) {
	src := "test/file.txt"
	r, err := PackAESOne(src)
	if err != nil {
		t.Fatal("Error Pack AES One:", err)
	}
	err = ioutil.WriteFile("test/file_aes.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write AES One:", err)
	}
}

func TestAESEncrypt(t *testing.T) {
	src := []byte("hello,world!")
	key := []byte("Satellite-266414")
	r, err := AESEncrypt(src, key)
	if err != nil {
		t.Fatal("Error AES Encrypt:", err)
	}
	err = ioutil.WriteFile("test/file_aes.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write AES One:", err)
	}
}

func TestSplitByte(t *testing.T) {
	data := []byte("hello,World!")
	size := ConstAESBufferSize
	_, err := SplitByte(data, size)
	if err != nil {
		t.Fatal("Error Split Byte:", err)
	}
}

func BenchmarkPackAES(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []string{"test/file_1.txt", "test/file_2.txt", "test/file_3.txt", "test/file_4.txt", "test/file_5.txt"}
		dest := "test/file_aes.txt"
		err := PackAES(src, dest)
		if err != nil {
			b.Fatal("Error Pack AES:", err)
		}
	}
}

func BenchmarkPackAESOneGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r []byte
		src := "test/file.txt"
		wg.Add(1)
		err := PackAESOneGo(src, &r, &wg)
		if err != nil {
			b.Fatal("Error Pack AES One Go:", err)
		}
		wg.Wait()
		err = ioutil.WriteFile("test/file_aes.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write AES One Go:", err)
		}
	}
}

func BenchmarkPackAESOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "test/file.txt"
		r, err := PackAESOne(src)
		if err != nil {
			b.Fatal("Error Pack AES One:", err)
		}
		err = ioutil.WriteFile("test/file_aes.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write AES One:", err)
		}
	}
}

func BenchmarkAESEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte("hello,world!")
		key := []byte("Satellite-266414")
		_, err := AESEncrypt(src, key)
		if err != nil {
			b.Fatal("Error AES Encrypt:", err)
		}
	}
}

func BenchmarkSplitByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := []byte("hello,World!")
		size := ConstAESBufferSize
		_, err := SplitByte(data, size)
		if err != nil {
			b.Fatal("Error Split Byte:", err)
		}
	}
}