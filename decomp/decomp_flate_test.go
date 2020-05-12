package decomp

import (
	"fmt"
	"testing"
)

func TestDeCompressFlate(t *testing.T) {
	src := []byte{202, 72, 205, 201, 201, 215, 41, 207, 47, 202, 73, 81, 4, 0, 0, 0, 255, 255}
	dest, err := DeCompressFlate(src)
	if err != nil {
		t.Fatal("Error decompress flate:", err)
	}
	fmt.Println("The src:", string(src))
	fmt.Println("The dest:", string(dest))
}
