package comp

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
)

// CompressTar function
// input src file list, output dest file path, return error info
// this function will use tar algorithm to compress file list
// src file support both absolute and relative paths, like 'C:\\file.txt' or '../test/data/file.txt'
// dest file also support both absolute and relative paths, like 'C:\\package.pak' or '../test/data/package.pak'
// return err indicate the success or failure function execute
func CompressTar(src []string, dest string) (err error) {
	// create the dest tar file...
	file, err := os.Create(dest)
	if err != nil {
		log.Println("Error create file:", err)
		return err
	}
	defer file.Close()
	// apply one tar writer to write file
	tw := tar.NewWriter(file)
	defer tw.Close()
	// loop compress src list files
	for _, v := range src {
		err = filepath.Walk(v, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println("Error compress file:", err)
				return err
			}
			// get file information header
			header, err := tar.FileInfoHeader(info, "")
			if err != nil {
				log.Println("Error get file info header:", err)
				return err
			}
			// assign compress method
			err = tw.WriteHeader(header)
			if err != nil {
				log.Println("Error write compress file header:", err)
				return err
			}
			// open the src file...
			data, err := os.Open(path)
			if err != nil {
				log.Println("Error open file:", err)
				return err
			}
			defer data.Close()
			// write compress data into file
			_, err = io.Copy(tw, data)
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

func CompressTarGz(src []string, dest string) (err error) {
	// create the dest tar.gz file...
	file, err := os.Create(dest)
	if err != nil {
		log.Println("Error create file:", err)
		return err
	}
	defer file.Close()
	// apply one gzip writer to write file
	gw := gzip.NewWriter(file)
	defer gw.Close()
	// apply one tar writer to write file
	tw := tar.NewWriter(gw)
	defer tw.Close()
	// loop compress src list files
	for _, v := range src {
		err = filepath.Walk(v, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println("Error compress file:", err)
				return err
			}
			// get file information header
			header, err := tar.FileInfoHeader(info, "")
			if err != nil {
				log.Println("Error get file info header:", err)
				return err
			}
			// assign compress method
			err = tw.WriteHeader(header)
			if err != nil {
				log.Println("Error write compress file header:", err)
				return err
			}
			// open the src file...
			data, err := os.Open(path)
			if err != nil {
				log.Println("Error open file:", err)
				return err
			}
			defer data.Close()
			// write compress data into file
			_, err = io.Copy(tw, data)
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
