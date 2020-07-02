package logging

import (
	"log"
	"testing"
)

// TestLoggingPlain function
// test logging package with simple text
func TestLoggingPlain(t *testing.T) {
	log.Println("hello,world!")
}

// TestLoggingWithParameters function
// test logging package with parameters
func TestLoggingWithParameters(t *testing.T) {
	a, b := 3, 5
	log.Println("a =", a)
	log.Println("b =", b)
	log.Printf("a + b = %v\n", a + b)
}
