package pack

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGobEncode(t *testing.T) {
	type test struct {
		Name  string
		Value int
	}
	buf := bytes.Buffer{}
	in := test{Name: "test_gob", Value: 5}
	err := GobEncode(&buf, in)
	if err != nil {
		t.Fatal("Error gob encode:", err)
	}
	fmt.Println(buf.String())
}

func BenchmarkGobEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		type test struct {
			Name  string
			Value int
		}
		buf := bytes.Buffer{}
		in := test{Name: "test_gob", Value: 5}
		err := GobEncode(&buf, in)
		if err != nil {
			b.Fatal("Error gob encode:", err)
		}
	}
}

func TestGobEncodeTo(t *testing.T) {
	type test struct {
		Name  string
		Value int
	}
	src := "../test/data/pack/file_gob.gob"
	in := test{Name: "test_gob", Value: 5}
	err := GobEncodeTo(src, in)
	if err != nil {
		t.Fatal("Error gob encode to file:", err)
	}
}

func BenchmarkGobEncodeTo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		type test struct {
			Name  string
			Value int
		}
		src := "../test/data/pack/file_gob.gob"
		in := test{Name: "test_gob", Value: 5}
		err := GobEncodeTo(src, in)
		if err != nil {
			b.Fatal("Error gob encode to file:", err)
		}
	}
}
