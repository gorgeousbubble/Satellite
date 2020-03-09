package task

import (
	"errors"
	"fmt"
	"satellite/app/erika/logs"
	"satellite/app/erika/parse"
	"satellite/app/erika/proc"
	. "satellite/utils"
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

func (tc *TaskCenter) Discovery(dir string) (err error) {
	// discovery task list...
	files, err := ListFiles(dir, "json")
	if err != nil {
		logs.Error("Error walk dir file list:", err)
		return err
	}
	// initial error
	err = errors.New("no task found")
	// unmarshal file
	for _, v := range files {
		// prepare for data parse
		s := parse.TAutoExec{}
		path := dir + "/" + v.Name()
		// unmarshal data to struct
		err = parse.Unmarshal(path, &s)
		if err != nil {
			logs.Info("Error Unmarshal file:", err)
			continue
		}
		// create task process...
		var location string
		var commands []string
		location = s.Attribute.Location
		for _, v := range s.Task.Exec {
			commands = append(commands, v.Command)
		}
		tc.Create(s.Schedule, func() { proc.HandleAutoExec(location, commands) })
		err = nil
		break
	}
	return err
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
