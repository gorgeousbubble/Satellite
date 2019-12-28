package comp

import (
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func CompressGzip(src []string, dest string) (err error) {
	// create the dest tar file...
	file, err := os.Create(dest)
	if err != nil {
		log.Println("Error create file:", err)
		return err
	}
	defer file.Close()
	// apply one gzip writer to write file
	gw := gzip.NewWriter(file)
	defer gw.Close()
	// loop compress src list files
	for _, v := range src {
		err = filepath.Walk(v, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println("Error compress file:", err)
				return err
			}
			// open the src file...
			data, err := os.Open(path)
			if err != nil {
				log.Println("Error open file:", err)
				return err
			}
			defer data.Close()
			// write gzip writer info
			gw.Name = info.Name()
			gw.Comment = "gzip compress by satellite"
			gw.ModTime = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.UTC)
			// write compress data into file
			_, err = io.Copy(gw, data)
			if err != nil {
				log.Println("Error write compress data into file:", err)
				return err
			}
			return err
		})
		if err != nil {
			log.Println("Error compress file:", err)
			return err
		}
	}
	return err
}
