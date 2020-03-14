package app

import (
	"errors"
	"fmt"
	"runtime"
	"satellite/app/erika/logs"
	"satellite/app/erika/nets"
	"satellite/app/erika/task"
	"satellite/exec"
	"strings"
)

func Start(ip string, port string, force bool) (err error) {
	// check if there any process running
	exist, err := isAlive()
	if err != nil {
		logs.Error("Error check process exist:", err)
		return err
	}
	if exist {
		if !force {
			logs.Error("Erika is already running. Please shutdown it first before start an instance.")
			err = errors.New("erika is already running")
			return err
		}
		err = kill()
		if err != nil {
			logs.Error("Error execute kill process:", err)
			return err
		}
	}
	// new task center instance...
	tc := task.NewTaskCenter()
	// create task job
	err = tc.Create("@hourly", func() { tc.Discovery("./task") })
	if err != nil {
		logs.Error("Error create job in task center:", err)
		return err
	}
	// automatic discovery task...
	err = tc.Discovery("./task")
	if err != nil {
		logs.Error("Error automatic discovery tasks:", err)
		return err
	}
	// start task...
	tc.Start()
	defer tc.Start()
	// erika tips
	fmt.Println("Erika Start...")
	// start http service
	err = nets.StartHttpServer(ip, port)
	if err != nil {
		logs.Error("Error start http server:", err)
		return err
	}
	return err
}

func Stop() (err error) {
	fmt.Println("Erika Stop...")
	err = kill()
	if err != nil {
		logs.Error("Error execute kill process:", err)
		return err
	}
	fmt.Println("Erika has successfully stopped.")
	return err
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
