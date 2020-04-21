package decomp

import (
	"compress/bzip2"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"satellite/routes"
	"satellite/utils"
)

// DeCompressBzip2 function
// input src file list, output dest file path and algorithm which used in unpack, return error info
// this function will base on algorithm to call correspond function
// src file support both absolute and relative paths, like 'C:\\file.bz2' or '../test/data/file.bz2'
// dest file also support both absolute and relative paths, like 'C:\\' or '../test/data/'
// return err indicate the success or failure function execute
func DeCompressBzip2(src []string, dest string) (err error) {
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
	// loop decompress src list files
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
			// apply one bzip2 reader to read file
			br := bzip2.NewReader(data)
			// read the extend of file name
			target := routes.TrimSuffixSlash(dest) + "/" + routes.TrimSuffixPoint(info.Name())
			//...
			// create the dest tar file...
			file, err := os.Create(target)
			if err != nil {
				log.Println("Error create file:", err)
				return err
			}
			defer file.Close()
			// read compress data into file
			_, err = io.Copy(file, br)
			if err != nil {
				log.Println("Error read decompress data into file:", err)
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
