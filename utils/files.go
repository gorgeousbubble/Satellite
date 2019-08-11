package utils

import (
	"log"
	"os"
	"path/filepath"
)

func PathExist(dir string) (is bool, err error) {
	_, err = os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Path not exist")
		} else {
			log.Println("Error status:", err)
		}
		return false, err
	}
	return true, err
}

func ListFiles(dir string) (files []os.FileInfo, err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return err
		}
		files = append(files, info)
		return err
	})
	if err != nil {
		log.Println("Error list files:",  err)
	}
	return
}
