package pack

import (
	"testing"
)

func TestPackHashEncodeMD5(t *testing.T) {
	src := "hello,world!"
	algorithm := "md5"
	dest, err := PackHashEncode(src, algorithm)
	if err != nil {
		t.Fatal("Error pack hash encode:", err)
	}
	if dest != "c0e84e870874dd37ed0d164c7986f03a" {
		t.Fatal("Error encode result")
	}
}