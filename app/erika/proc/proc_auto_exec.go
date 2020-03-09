package proc

import (
	"satellite/app/erika/logs"
	"satellite/exec"
)

func HandleAutoExec(location string, commands []string) (err error) {
	var s string
	var r string
	// combine command and record it
	s = "cd " + location
	logs.Info("Execute command:", s)
	// enter location which will execute command
	r, err = exec.Exec(s)
	if err != nil {
		logs.Error("Error execute command:", err)
		logs.Error("Result:", r)
		return err
	}
	logs.Info("Result:", r)
	// execute all commands...
	for _, v := range commands {
		logs.Info("Execute command:", v)
		// execute one command
		r, err = exec.Exec(v)
		if err != nil {
			logs.Error("Error execute command:", err)
			logs.Error("Result:", r)
			return err
		}
		logs.Info("Result:", r)
	}
	return err
}
