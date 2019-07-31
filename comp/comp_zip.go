package comp

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func CompressZipOneGo(src string, dest string, wg *sync.WaitGroup) (err error) {
	err = CompressZipOne(src, dest)
	if err != nil {
		log.Println("Error Compress Zip One:", err)
	}
	wg.Done()
	return err
}

func CompressZipOne(src string, dest string) (err error) {
	// create the dest zip file...
	file, err := os.Create(dest)
	if err != nil {
		log.Println("Error create file:", err)
		return err
	}
	defer file.Close()
	// apply one zip writer to write file
	archive := zip.NewWriter(file)
	defer archive.Close()
	// compress zip files...
	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("Error compress file:", err)
			return err
		}
		// get file information header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			log.Println("Error get file info header:", err)
			return err
		}
		// assign compress method
		header.Method = zip.Deflate
		writer, err := archive.CreateHeader(header)
		if err != nil {
			log.Println("Error create compress file header:", err)
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
		_, err = io.Copy(writer, data)
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
	return err
}
