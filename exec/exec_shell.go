package exec

import (
	"bytes"
	"os/exec"
)

func ExecShell(s string) (r string, err error) {
	var out bytes.Buffer
	cmd := exec.Command("/bin/bash", "-c", s)
	cmd.Stdout = &out
	err = cmd.Run()
	return out.String(), err
}
