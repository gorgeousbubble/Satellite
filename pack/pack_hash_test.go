package pack

import (
	"fmt"
	"testing"
)

func TestPackHashEncodeMD5(t *testing.T) {
	src := "hello,world!"
	algorithm := "md5"
	dest, err := PackHashEncode(src, algorithm)
	if err != nil {
		t.Fatal("Error pack hash encode:", err)
	}
	fmt.Println("Hash md5 encode:", dest)
}

func BenchmarkPackHashEncodeMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		algorithm := "md5"
		_, err := PackHashEncode(src, algorithm)
		if err != nil {
			b.Fatal("Error pack hash encode:", err)
		}
	}
}

func TestPackHashEncodeSHA1(t *testing.T) {
	src := "hello,world!"
	algorithm := "sha1"
	dest, err := PackHashEncode(src, algorithm)
	if err != nil {
		t.Fatal("Error pack hash encode:", err)
	}
	fmt.Println("Hash sha1 encode:", dest)
}

func BenchmarkPackHashEncodeSHA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		algorithm := "sha1"
		_, err := PackHashEncode(src, algorithm)
		if err != nil {
			b.Fatal("Error pack hash encode:", err)
		}
	}
}

func TestPackHashEncodeSHA256(t *testing.T) {
	src := "hello,world!"
	algorithm := "sha256"
	dest, err := PackHashEncode(src, algorithm)
	if err != nil {
		t.Fatal("Error pack hash encode:", err)
	}
	fmt.Println("Hash sha256 encode:", dest)
}

func BenchmarkPackHashEncodeSHA256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		algorithm := "sha256"
		_, err := PackHashEncode(src, algorithm)
		if err != nil {
			b.Fatal("Error pack hash encode:", err)
		}
	}
}

func TestPackHashEncodeSHA512(t *testing.T) {
	src := "hello,world!"
	algorithm := "sha512"
	dest, err := PackHashEncode(src, algorithm)
	if err != nil {
		t.Fatal("Error pack hash encode:", err)
	}
	fmt.Println("Hash sha512 encode:", dest)
}

func BenchmarkPackHashEncodeSHA512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		algorithm := "sha512"
		_, err := PackHashEncode(src, algorithm)
		if err != nil {
			b.Fatal("Error pack hash encode:", err)
		}
	}
}

func TestPackHashEncodeBlake2b128(t *testing.T) {
	src := "hello,world!"
	algorithm := "blake2b128"
	dest, err := PackHashEncode(src, algorithm)
	if err != nil {
		t.Fatal("Error pack hash encode:", err)
	}
	fmt.Println("Hash blake2b encode:", dest)
}

func BenchmarkPackHashEncodeBlake2b128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		algorithm := "blake2b128"
		_, err := PackHashEncode(src, algorithm)
		if err != nil {
			b.Fatal("Error pack hash encode:", err)
		}
	}
}

func TestPackHashEncodeBlake2b256(t *testing.T) {
	src := "hello,world!"
	algorithm := "blake2b256"
	dest, err := PackHashEncode(src, algorithm)
	if err != nil {
		t.Fatal("Error pack hash encode:", err)
	}
	fmt.Println("Hash blake2b encode:", dest)
}

func BenchmarkPackHashEncodeBlake2b256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		algorithm := "blake2b256"
		_, err := PackHashEncode(src, algorithm)
		if err != nil {
			b.Fatal("Error pack hash encode:", err)
		}
	}
}

func TestPackHashEncodeBlake2b512(t *testing.T) {
	src := "hello,world!"
	algorithm := "blake2b512"
	dest, err := PackHashEncode(src, algorithm)
	if err != nil {
		t.Fatal("Error pack hash encode:", err)
	}
	fmt.Println("Hash blake2b encode:", dest)
}

func BenchmarkPackHashEncodeBlake2b512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "hello,world!"
		algorithm := "blake2b512"
		_, err := PackHashEncode(src, algorithm)
		if err != nil {
			b.Fatal("Error pack hash encode:", err)
		}
	}
}
