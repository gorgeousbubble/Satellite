package comp

import (
	"fmt"
	"testing"
)

func TestCompressFlate(t *testing.T) {
	src := []byte("hello,world!")
	dest, err := CompressFlate(src)
	if err != nil {
		t.Fatal("Error compress flate:", err)
	}
	fmt.Println("The src:", string(src))
	fmt.Println("The dest:", string(dest))
}

func BenchmarkCompressFlate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte("hello,world!")
		_, err := CompressFlate(src)
		if err != nil {
			b.Fatal("Error compress flate:", err)
		}
	}
}
