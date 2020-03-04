package task

import (
	"fmt"
	"testing"
	"time"
)

func TestTaskTimer(t *testing.T) {
	ch := make(chan interface{})
	timer := NewTaskTimer(func() {
		fmt.Println("hello,world!")
	}, ch)
	timer.Start(time.Second)
	time.Sleep(time.Second * 5)
	timer.Stop()
}
