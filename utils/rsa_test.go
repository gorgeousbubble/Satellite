package utils

import "testing"

func TestGenRSAKey2File(t *testing.T) {
	err := GenRSAKey2File(1024)
	if err != nil {
		t.Errorf("Error Generate RSA Key to File.")
	}
}

func TestGenRSAKey2Memory(t *testing.T) {
	var private []byte
	var public []byte
	err := GenRSAKey2Memory(&private, &public, 1024)
	if err != nil {
		t.Errorf("Error Generate RSA Key to Memory.")
	}
}
