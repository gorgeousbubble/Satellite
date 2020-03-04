package unpack

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGobDecode(t *testing.T) {
	type test struct {
		Name  string
		Value int
	}
	buf := bytes.Buffer{}
	out := test{}
	buf.Write([]byte{37, 255, 129, 3, 1, 1, 4, 116, 101, 115, 116, 1, 255, 130, 0, 1, 2, 1, 4, 78, 97, 109, 101, 1, 12, 0, 1, 5, 86, 97, 108, 117, 101, 1, 4, 0, 0, 0, 15, 255, 130, 1, 8, 116, 101, 115, 116, 95, 103, 111, 98, 1, 10, 0})
	err := GobDecode(&buf, &out)
	if err != nil {
		t.Fatal("Error gob encode:", err)
	}
	if out.Name != "test_gob" {
		t.Fatal("Error parse name")
	}
	if out.Value != 5 {
		t.Fatal("Error parse value")
	}
	fmt.Println(out)
}

func BenchmarkGobDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		type test struct {
			Name  string
			Value int
		}
		buf := bytes.Buffer{}
		out := test{}
		buf.Write([]byte{37, 255, 129, 3, 1, 1, 4, 116, 101, 115, 116, 1, 255, 130, 0, 1, 2, 1, 4, 78, 97, 109, 101, 1, 12, 0, 1, 5, 86, 97, 108, 117, 101, 1, 4, 0, 0, 0, 15, 255, 130, 1, 8, 116, 101, 115, 116, 95, 103, 111, 98, 1, 10, 0})
		err := GobDecode(&buf, &out)
		if err != nil {
			b.Fatal("Error gob encode:", err)
		}
		if out.Name != "test_gob" {
			b.Fatal("Error parse name")
		}
		if out.Value != 5 {
			b.Fatal("Error parse value")
		}
	}
}

func TestGobDecodeFrom(t *testing.T) {
	type test struct {
		Name  string
		Value int
	}
	src := "../test/data/unpack/file_gob.gob"
	out := test{}
	err := GobDecodeFrom(src, &out)
	if err != nil {
		t.Fatal("Error gob decode from file:", err)
	}
}

func BenchmarkGobDecodeFrom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		type test struct {
			Name  string
			Value int
		}
		src := "../test/data/unpack/file_gob.gob"
		out := test{}
		err := GobDecodeFrom(src, &out)
		if err != nil {
			b.Fatal("Error gob decode from file:", err)
		}
	}
}
