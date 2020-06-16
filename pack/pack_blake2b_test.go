package pack

import (
	"io/ioutil"
	"testing"
)

func TestBlake2bEncode128(t *testing.T) {
	src := "hello,world!"
	r := Blake2bEncode128(src)
	err := ioutil.WriteFile("../test/data/pack/file_blake2b_128.txt", []byte(r), 0644)
	if err != nil {
		t.Fatal("Error Write blake2b Encode:", err)
	}
}

func BenchmarkBlake2bEncode128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		r := Blake2bEncode128(src)
		err := ioutil.WriteFile("../test/data/pack/file_blake2b_128.txt", []byte(r), 0644)
		if err != nil {
			b.Fatal("Error Write blake2b Encode:", err)
		}
	}
}

func TestBlake2bEncode256(t *testing.T) {
	src := "hello,world!"
	r := Blake2bEncode256(src)
	err := ioutil.WriteFile("../test/data/pack/file_blake2b_256.txt", []byte(r), 0644)
	if err != nil {
		t.Fatal("Error Write blake2b Encode:", err)
	}
}

func BenchmarkBlake2bEncode256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		r := Blake2bEncode256(src)
		err := ioutil.WriteFile("../test/data/pack/file_blake2b_256.txt", []byte(r), 0644)
		if err != nil {
			b.Fatal("Error Write blake2b Encode:", err)
		}
	}
}

func TestBlake2bEncode512(t *testing.T) {
	src := "hello,world!"
	r := Blake2bEncode512(src)
	err := ioutil.WriteFile("../test/data/pack/file_blake2b_512.txt", []byte(r), 0644)
	if err != nil {
		t.Fatal("Error Write blake2b Encode:", err)
	}
}

func BenchmarkBlake2bEncode512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		r := Blake2bEncode512(src)
		err := ioutil.WriteFile("../test/data/pack/file_blake2b_512.txt", []byte(r), 0644)
		if err != nil {
			b.Fatal("Error Write blake2b Encode:", err)
		}
	}
}

func TestBlake2bCheck128(t *testing.T) {
	src := "hello,world!"
	dest := "1748e3d0f53508245851db4571424eee"
	r := Blake2bCheck128(src, dest)
	if !r {
		t.Fatal("Error check blake2b 128 encode:", r)
	}
}

func TestBlake2bCheck256(t *testing.T) {
	src := "hello,world!"
	dest := "8268578331a07de98347abd8cf11addf924a4ea0ac75e4aec1bf3fe6cb314553"
	r := Blake2bCheck256(src, dest)
	if !r {
		t.Fatal("Error check blake2b 256 encode:", r)
	}
}
