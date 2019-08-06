package comp

import (
	"fmt"
	"testing"
)

func TestCompressZlib(t *testing.T) {
	src := "hello,world!"
	fmt.Println("Before Compress Zlib:", src)
	dest, err := CompressZlib([]byte(src))
	if err != nil {
		t.Fatal("Error Compress Zlib:", err)
	}
	fmt.Println("After Compress Zlib:", dest)
}

func BenchmarkCompressZlib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		_, err := CompressZlib([]byte(src))
		if err != nil {
			b.Fatal("Error Compress Zlib:", err)
		}
	}
}
