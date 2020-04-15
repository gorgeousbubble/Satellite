package comp

import (
	"compress/gzip"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"satellite/routes"
	"satellite/utils"
	"time"
)

// CompressGzip function
// input src file list, output dest file path, return error info
// this function will use gzip algorithm to compress file list
// src file support both absolute and relative paths, like 'C:\\file.txt' or '../test/data/file.txt'
// dest file also support both absolute and relative paths, like 'C:\\package.gz' or '../test/data/package.gz'
// return err indicate the success or failure function execute
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
			target := routes.TrimSuffixSlash(dest) + "/" + routes.TrimSuffixPoint(info.Name()) + ".gz"
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
			gw.Extra = []byte(routes.GetSuffixPoint(info.Name()))
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
