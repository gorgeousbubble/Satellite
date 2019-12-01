package shell

import "testing"

func TestShellUpx(t *testing.T) {
	src := "../test/data/shell/package.exe"
	dest := "../test/data/shell/package_upx.exe"
	err := shellUpx(src, dest)
	if err != nil {
		t.Fatal("Error shell upx:", err)
	}
}
