package parses

import (
	"bytes"
	"fmt"
	"testing"
)

func TestExtractOneElement(t *testing.T) {
	s := []byte("hello,1,[1,2,3],{simple, 12}")
	r, sub, err := extractOneElement(s)
	if err != nil {
		t.Fatal("Error extract one element:", err)
	}
	fmt.Println("extract return:", string(r))
	fmt.Println("extract substring:", string(sub))
	if !bytes.Equal(r, []byte("hello")) {
		t.Fatal("Error extract return")
	}
	if !bytes.Equal(sub, []byte("1,[1,2,3],{simple, 12}")) {
		t.Fatal("Error extract substring")
	}
}

func BenchmarkExtractOneElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("hello,1,[1,2,3],{simple, 12}")
		r, sub, err := extractOneElement(s)
		if err != nil {
			b.Fatal("Error extract one element:", err)
		}
		if !bytes.Equal(r, []byte("hello")) {
			b.Fatal("Error extract return")
		}
		if !bytes.Equal(sub, []byte("1,[1,2,3],{simple, 12}")) {
			b.Fatal("Error extract substring")
		}
	}
}
