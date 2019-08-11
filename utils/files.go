package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type FileInfos []os.FileInfo

func (s FileInfos) Len() int {
	return len(s)
}

func (s FileInfos) Less(i, j int) bool {
	return s[i].ModTime().Before(s[j].ModTime())
}

func (s FileInfos) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

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

func ListFiles(dir string, term string) (files []os.FileInfo, err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return err
		}
		if strings.Contains(info.Name(), term) {
			files = append(files, info)
		}
		return err
	})
	if err != nil {
		log.Println("Error list files:",  err)
	}
	return
}
