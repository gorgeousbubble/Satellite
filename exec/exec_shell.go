package exec

import (
	"bytes"
	"os/exec"
)

// ExecShell function
// this function is mainly used to execute command in Linux
// input command string which will be executed in target os system
// return execute result and err indicate the success or failure function execute
func ExecShell(s string) (r string, err error) {
	var out bytes.Buffer
	cmd := exec.Command("/bin/bash", "-c", s)
	cmd.Stdout = &out
	err = cmd.Run()
	return out.String(), err
}
