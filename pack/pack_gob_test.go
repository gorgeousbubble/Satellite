package pack

import "testing"

func TestGobEncodeTo(t *testing.T) {
	type test struct {
		Name string
		Value int
	}
	src := "../test/data/pack/file_gob.gob"
	in := test{Name:"test_gob", Value:5}
	err := GobEncodeTo(src, in)
	if err != nil {
		t.Fatal("Error gob encode to file:", err)
	}
}
