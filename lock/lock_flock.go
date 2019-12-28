// +build !windows

package lock

import (
	"fmt"
	"os"
	"sync"
	"syscall"
)

type FileLock struct {
	file *os.File
	path string
}

func New(path string) *FileLock {
	return &FileLock{path: path}
}

func (flock *FileLock) Lock() (err error) {
	file, err := os.Open(flock.path)
	if err != nil {
		return err
	}
	flock.file = file
	err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		return err
	}
	return err
}

func (flock *FileLock) Unlock() (err error) {
	defer flock.file.Close()
	return syscall.Flock(int(flock.file.Fd()), syscall.LOCK_UN)
}
