package proc

import (
	"satellite/app/erika/logs"
	"satellite/exec"
	"strings"
	"time"
)

func HandleAutoMonitor(location, period, format, width, height string) (err error) {
	var s string
	var r string
	// get time now
	t := time.Now().Format("2006-01-02 15:04:05")
	t = strings.ReplaceAll(t, "-", "")
	t = strings.ReplaceAll(t, ":", "")
	t = strings.ReplaceAll(t, " ", "")
	// combine command for monitor
	s = "raspivid -o " + location + "/rasp_" + t + "." + format + " -t " + period + " -w " + width + " -h " + height
	logs.Info("Execute command:", s)
	// execute one command
	r, err = exec.Exec(s)
	if err != nil {
		logs.Error("Error execute command:", err)
		logs.Error("Result:", r)
		return err
	}
	logs.Info("Result:", r)
	return err
}
