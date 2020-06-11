package rasp

import (
	"fmt"
	"log"
	"satellite/exec"
	"strings"
	"time"
)

type RaspVideo struct {
	Type string
	Time string
}

func (rv *RaspVideo) Start() (err error) {
	// get time now
	t := time.Now().Format("2006-01-02 15:04:05")
	s := strings.ReplaceAll(t, "-", "")
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, " ", "")
	// combine command for video
	name := fmt.Sprintf("Rasp%v.%v", s, rv.Type)
	command := fmt.Sprintf("raspivid -o %v -t %v", name, rv.Time)
	// start execute command
	_, err = exec.Exec(command)
	if err != nil {
		log.Println("Error exec raspberry video:", err)
		return err
	}
	return err
}
