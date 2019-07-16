package utils

import "testing"

func TestGenRSAKey(t *testing.T) {
	err := GenRSAKey(1024)
	if err != nil {
		t.Errorf("Error Generate RSA Key.")
	}
}
