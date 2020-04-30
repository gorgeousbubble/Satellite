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

func TestPackHashEncodeSHA1(t *testing.T) {
	src := "hello,world!"
	algorithm := "sha1"
	dest, err := PackHashEncode(src, algorithm)
	if err != nil {
		t.Fatal("Error pack hash encode:", err)
	}
	fmt.Println("Hash sha1 encode:", dest)
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

func TestPackHashEncodeSHA512(t *testing.T) {
	src := "hello,world!"
	algorithm := "sha512"
	dest, err := PackHashEncode(src, algorithm)
	if err != nil {
		t.Fatal("Error pack hash encode:", err)
	}
	fmt.Println("Hash sha512 encode:", dest)
}