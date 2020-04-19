package graphics

import "testing"

func TestDX9GraphicsCreate(t *testing.T) {
	dx9 := &DX9Graphics{}
	defer dx9.Release()
	err := dx9.Create()
	if err != nil {
		t.Fatal("Error create dx9 graphics:", err)
	}
}
