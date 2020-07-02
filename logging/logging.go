package logging

import (
	"fmt"
	"log"
	"os"
	. "satellite/global"
	. "satellite/utils"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Notice about logging package...
// This package is a simple multiple files logger which used to record parameters during debug process
// You can easy to use it by import package like this [import _ "satellite/logging"]
// Because you do not really import any function in this package, just need call 'init' function to redirect log package
// When you import package in your code, you can use log.Println, log.Printf, log.Print ... to record your log
func init() {
	// Get current time now
	t := time.Now()
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error get dir:", err)
		os.Exit(1)
	}
	dir = strings.ReplaceAll(dir, "\\", "/")
	path := dir + LogPath
	name := dir + LogPath + "/" + AppName + "_" + strconv.Itoa(t.Year()) + "_" + strconv.Itoa(int(t.Month())) + "_" + strconv.Itoa(t.Day()) + ".log"
	// Check whether path exist
	is, _ := PathExist(path)
	if !is {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Println("Error mkdir:", err)
			os.Exit(1)
		}
	}
	// Only keep 5 log files
	files, err := ListFiles(path, AppName)
	if err != nil {
		fmt.Println("Error open log file:", err)
		os.Exit(1)
	}
	if len(files) >= LogNumber {
		sort.Sort(FileInfos(files))
		del := path + "/" + files[len(files)-1].Name()
		err = os.Remove(del)
		if err != nil {
			fmt.Println("Error remove log file:", err)
			os.Exit(1)
		}
	}
	// Create logging file
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error open log file:", err)
		os.Exit(1)
	}
	//defer file.Close()
	// Set log settings
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(file)
}
