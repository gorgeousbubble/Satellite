package app

import (
	"fmt"
	"os"
	"runtime"
	"satellite/app/erika/logs"
	"satellite/app/erika/nets"
	"satellite/app/erika/task"
	"satellite/exec"
	"strings"
)

func Run(ip string, port string) {
	// check if there any process running
	exist, err := isAlive()
	if err != nil {
		logs.Error("Error check process exist:", err)
	}
	if exist {
		fmt.Println("Erika is already running. Please shutdown it first before start an instance.")
		os.Exit(1)
	}
	// new task center instance...
	tc := task.NewTaskCenter()
	// create task job
	err = tc.Create("@hourly", func() { logs.Info("hello,World!") })
	if err != nil {
		logs.Error("Error create job in task center:", err)
	}
	// start task...
	tc.Start()
	defer tc.Start()
	// erika tips
	fmt.Println("Erika Start...")
	// start http service
	nets.StartHttpServer(ip, port)
}

func isAlive() (exist bool, err error) {
	var s string
	var r string
	exist = true
	if runtime.GOOS == "windows" {
		s = "tasklist | findstr erika"
	} else {
		s = "ps -e | grep erika"
	}
	r, err = exec.Exec(s)
	if err != nil {
		logs.Error("Error execute command:", err)
		return exist, err
	}
	logs.Info("Execute result:", r)
	// check alive status...
	n := strings.Count(r, "erika")
	if n == 1 {
		exist = false
	}
	logs.Debug("Process number:", n)
	return exist, err
}

func kill() (err error) {
	var s string
	var r string
	if runtime.GOOS == "windows" {
		s = "taskkill /f /t /im erika.exe"
	} else {
		s = "killall erika"
	}
	r, err = exec.Exec(s)
	if err != nil {
		logs.Error("Error execute command:", err)
		return err
	}
	logs.Info("Execute result:", r)
	return err
}
