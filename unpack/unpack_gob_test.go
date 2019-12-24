package unpack

import (
	"testing"
)

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
