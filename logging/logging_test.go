package logging

import (
	"log"
	"testing"
)

func TestLoggingPlain(t *testing.T) {
	log.Println("hello,world!")
}

func TestLoggingWithParameters(t *testing.T) {
	a, b := 3, 5
	log.Println("a =", a)
	log.Println("b =", b)
	log.Printf("a + b = %v\n", a + b)
}
