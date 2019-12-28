package comp

import (
	"compress/gzip"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"satellite/utils"
	"strings"
	"time"
)

func CompressGzip(src []string, dest string) (err error) {
	// check the dest whether dir or not...
	is, err := utils.IsDir(dest)
	if err != nil {
		log.Println("Error check dest dir:", err)
		return err
	}
	if !is {
		err = errors.New("dest path is not dir")
		log.Println(err)
		return err
	}
	// loop compress src list files
	for _, v := range src {
		err = filepath.Walk(v, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println("Error compress file:", err)
				return err
			}
			target := trimSuffixSlash(dest) + "/" + trimSuffixPoint(info.Name()) + ".gz"
			//...
			// create the dest tar file...
			file, err := os.Create(target)
			if err != nil {
				log.Println("Error create file:", err)
				return err
			}
			defer file.Close()
			// open the src file...
			data, err := os.Open(path)
			if err != nil {
				log.Println("Error open file:", err)
				return err
			}
			defer data.Close()
			// read the data...
			buf, err := ioutil.ReadAll(data)
			if err != nil {
				log.Println("Error read file:", err)
				return err
			}
			// apply one gzip writer to write file
			gw := gzip.NewWriter(file)
			defer gw.Close()
			// write gzip writer info
			gw.Name = info.Name()
			gw.Comment = "gzip compress by satellite"
			gw.ModTime = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.UTC)
			// write compress data into file
			_, err = gw.Write(buf)
			if err != nil {
				log.Println("Error write compress data into file:", err)
				return err
			}
			// flush gzip writer
			err = gw.Flush()
			if err != nil {
				log.Println("Error flush gzip writer:", err)
				return err
			}
			/*// close gzip writer
			if err = gw.Close(); err != nil {
				log.Println("Error close gzip writer:", err)
				return err
			}
			// reset gzip writer
			gw.Reset(file)*/
			return err
		})
		if err != nil {
			log.Println("Error compress file:", err)
			return err
		}
	}
	return err
}

func trimSuffixPoint(name string) string {
	n := strings.Index(name, ".")
	if n == -1 {
		return name
	}
	short := name[0:n]
	return short
}

func trimSuffixSlash(name string) string {
	if name[len(name)-1] == '/' {
		return name[0 : len(name)-1]
	} else if name[len(name)-1] == '\\' {
		return name[0 : len(name)-2]
	}
	return name
}
