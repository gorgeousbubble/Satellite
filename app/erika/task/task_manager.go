package task

import "errors"

type TaskManager struct {
	M map[string]interface{}
}

func NewTaskManager() *TaskManager {
	// new task manager...
	tm := &TaskManager{}
	tm.M = make(map[string]interface{})
	return tm
}

func (tm *TaskManager) AppendTask(id string, resource interface{}) (err error) {
	// traversal the task manager list...
	for k := range tm.M {
		// key value already exist
		if id == k {
			err = errors.New("task already exist")
			return err
		}
	}
	// append new key value
	tm.M[id] = resource
	return err
}

func (tm *TaskManager) RemoveTask(id string) (err error) {
	var exist = false
	// traversal the task manager list...
	for k := range tm.M {
		if id == k {
			exist = true
			break
		}
	}
	// key value not exist
	if !exist {
		err = errors.New("task not exist")
		return err
	}
	// remove key value in map
	delete(tm.M, id)
	return err
}
