package decomp

import (
	"fmt"
	"testing"
)

// TestDeCompressZlib function
func TestDeCompressZlib(t *testing.T) {
	src := []byte{120, 156, 202, 72, 205, 201, 201, 215, 41, 207, 47, 202, 73, 81, 4, 4, 0, 0, 255, 255, 30, 221, 4, 138}
	fmt.Println("Before DeCompress Zlib:", string(src))
	dest, err := DeCompressZlib(src)
	if err != nil {
		t.Fatal("Error DeCompress Zlib:", err)
	}
	fmt.Println("After DeCompress Zlib:", string(dest))
}

// BenchmarkDeCompressZlib function
func BenchmarkDeCompressZlib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte{120, 156, 202, 72, 205, 201, 201, 215, 41, 207, 47, 202, 73, 81, 4, 4, 0, 0, 255, 255, 30, 221, 4, 138}
		_, err := DeCompressZlib(src)
		if err != nil {
			b.Fatal("Error DeCompress Zlib:", err)
		}
	}
}
