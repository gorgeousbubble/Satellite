package exec

import "runtime"

func Exec(s string) (r string, err error) {
	if runtime.GOOS == "windows" {
		r, err = ExecCmd(s)
	} else {
		r, err = ExecShell(s)
	}
	return r, err
}
