package task

import "time"

type TaskTimer struct {
	F  func()
	Ch chan interface{}
}

func NewTaskTimer(f func(), ch chan interface{}) *TaskTimer {
	timer := &TaskTimer{
		F:  f,
		Ch: ch,
	}
	return timer
}

func (t *TaskTimer) Start(d time.Duration) {
	ticker := time.NewTicker(d)
	go func() {
		for range ticker.C {
			select {
			case <-t.Ch:
				break
			default:
				t.F()
			}
		}
	}()
}

func (t *TaskTimer) Stop() {
	t.Ch <- struct{}{}
}
