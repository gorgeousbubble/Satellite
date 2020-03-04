package task

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

type TaskCenter struct {
	C *cron.Cron
}

func NewTaskCenter() *TaskCenter {
	// new task center...
	c := cron.New()
	tc := &TaskCenter{
		C: c,
	}
	return tc
}

func (tc *TaskCenter) Create(spec string, cmd func()) (err error) {
	// create a task...
	_, err = tc.C.AddFunc(spec, cmd)
	return err
}

func (tc *TaskCenter) Start() {
	tc.C.Start()
}

func (tc *TaskCenter) Stop() {
	tc.C.Stop()
}

func Run() (err error) {
	// create a new cron instance...
	c := cron.New()
	// add task function to task center, set time schedule...
	_, err = c.AddFunc("* * * * *", func() { fmt.Println(time.Now().Format("2006-01-02 15:04:05")) })
	if err != nil {
		return err
	}
	// start run task...
	c.Start()
	defer c.Stop()
	// do something else at during that time...
	// todo
	time.Sleep(time.Minute * 5)
	return err
}
