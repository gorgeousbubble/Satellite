package app

import (
	"fmt"
	"satellite/app/erika/logs"
	"satellite/app/erika/nets"
	"satellite/app/erika/task"
)

func Run(ip string, port string) {
	// new task center instance...
	tc := task.NewTaskCenter()
	// create task job
	err := tc.Create("@hourly", func() { logs.Info("hello,World!") })
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
