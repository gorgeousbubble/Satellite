package parses

import (
	"bytes"
	"fmt"
	"testing"
)

func TestExtractOneElement(t *testing.T) {
	s := []byte("hello,1,[1,2,3],{simple,12}")
	r, sub, err := extractOneElement(s)
	if err != nil {
		t.Fatal("Error extract one element:", err)
	}
	fmt.Println("extract return:", string(r))
	fmt.Println("extract substring:", string(sub))
	if !bytes.Equal(r, []byte("hello")) {
		t.Fatal("Error extract return")
	}
	if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
		t.Fatal("Error extract substring")
	}
}

func BenchmarkExtractOneElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("hello,1,[1,2,3],{simple,12}")
		r, sub, err := extractOneElement(s)
		if err != nil {
			b.Fatal("Error extract one element:", err)
		}
		if !bytes.Equal(r, []byte("hello")) {
			b.Fatal("Error extract return")
		}
		if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
			b.Fatal("Error extract substring")
		}
	}
}

func TestExtractOneList(t *testing.T) {
	s := []byte("[apple,orange],1,[1,2,3],{simple,12}")
	r, sub, err := extractOneList(s)
	if err != nil {
		t.Fatal("Error extract one list:", err)
	}
	fmt.Println("extract return:", string(r))
	fmt.Println("extract substring:", string(sub))
	if !bytes.Equal(r, []byte("[apple,orange]")) {
		t.Fatal("Error extract return")
	}
	if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
		t.Fatal("Error extract substring")
	}
}

func BenchmarkExtractOneList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("[apple,orange],1,[1,2,3],{simple,12}")
		r, sub, err := extractOneList(s)
		if err != nil {
			b.Fatal("Error extract one list:", err)
		}
		if !bytes.Equal(r, []byte("[apple,orange]")) {
			b.Fatal("Error extract return")
		}
		if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
			b.Fatal("Error extract substring")
		}
	}
}

func TestExtractOneTuple(t *testing.T) {
	s := []byte("{stars,514},1,[1,2,3],{simple,12}")
	r, sub, err := extractOneTuple(s)
	if err != nil {
		t.Fatal("Error extract one tuple:", err)
	}
	fmt.Println("extract return:", string(r))
	fmt.Println("extract substring:", string(sub))
	if !bytes.Equal(r, []byte("{stars,514}")) {
		t.Fatal("Error extract return")
	}
	if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
		t.Fatal("Error extract substring")
	}
}

func BenchmarkExtractOneTuple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("{stars,514},1,[1,2,3],{simple,12}")
		r, sub, err := extractOneTuple(s)
		if err != nil {
			b.Fatal("Error extract one tuple:", err)
		}
		if !bytes.Equal(r, []byte("{stars,514}")) {
			b.Fatal("Error extract return")
		}
		if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
			b.Fatal("Error extract substring")
		}
	}
}
