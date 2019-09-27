package exec

import (
	"bytes"
	"os/exec"
)

func ExecCmd(s string) (r string, err error) {
	var out bytes.Buffer
	cmd := exec.Command("cmd", "/C", s)
	cmd.Stdout = &out
	err = cmd.Run()
	return out.String(), err
}
