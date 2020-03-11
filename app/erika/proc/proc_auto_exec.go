package proc

import (
	"satellite/app/erika/logs"
	"satellite/exec"
)

func HandleAutoExec(location string, commands []string) (err error) {
	var s string
	var r string
	// execute all commands...
	for _, v := range commands {
		// combine command and record it
		s = "cd " + location + ";" + v
		logs.Info("Execute command:", s)
		// execute one command
		r, err = exec.Exec(s)
		if err != nil {
			logs.Error("Error execute command:", err)
			logs.Error("Result:", r)
			return err
		}
		logs.Info("Result:", r)
	}
	return err
}
