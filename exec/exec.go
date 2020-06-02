package exec

import "runtime"

// Exec function
// this function is mainly used to execute command in os system support both Windows and Linux
// support command in Windows, Linux, etc.
// input command string which will be executed in target os system
// return execute result and err indicate the success or failure function execute
func Exec(s string) (r string, err error) {
	if runtime.GOOS == "windows" {
		r, err = ExecCmd(s)
	} else {
		r, err = ExecShell(s)
	}
	return r, err
}
