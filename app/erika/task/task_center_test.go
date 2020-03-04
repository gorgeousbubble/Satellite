package task

import "testing"

func TestRun(t *testing.T) {
	err := Run()
	if err != nil {
		t.Fatal("Error run task:", err)
	}
}
