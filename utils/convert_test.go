package utils

import "testing"

func TestSplitByte(t *testing.T) {
	data := []byte("hello,World!")
	size := 128
	_, err := SplitByte(data, size)
	if err != nil {
		t.Fatal("Error Split Byte:", err)
	}
}

func BenchmarkSplitByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := []byte("hello,World!")
		size := 128
		_, err := SplitByte(data, size)
		if err != nil {
			b.Fatal("Error Split Byte:", err)
		}
	}
}
