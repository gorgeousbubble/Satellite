package task

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

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
