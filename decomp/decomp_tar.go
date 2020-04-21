package decomp

import (
	"archive/tar"
	"compress/bzip2"
	"compress/gzip"
	"io"
	"log"
	"os"
	"strings"
)

// DeCompressTar function
// input src file list, output dest file path and algorithm which used in unpack, return error info
// this function will base on algorithm to call correspond function
// src file support both absolute and relative paths, like 'C:\\file.tar' or '../test/data/file.tar'
// dest file also support both absolute and relative paths, like 'C:\\' or '../test/data/'
// return err indicate the success or failure function execute
func DeCompressTar(src string, dest string) (err error) {
	// open the src tar ball file...
	file, err := os.Open(src)
	if err != nil {
		log.Println("Error open tar file:", err)
		return err
	}
	defer file.Close()
	tr := tar.NewReader(file)
	// loop decompress src list files
	for {
		header, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		name := dest + header.Name
		err = os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
		if err != nil {
			log.Println("Error make dir all:", err)
			return err
		}
		f, err := os.Create(name)
		if err != nil {
			log.Println("Error create name:", err)
			return err
		}
		_, err = io.Copy(f, tr)
		if err != nil {
			log.Println("Error write decompress date:", err)
			return err
		}
	}
	return err
}

// DeCompressTarGz function
// input src file list, output dest file path and algorithm which used in unpack, return error info
// this function will base on algorithm to call correspond function
// src file support both absolute and relative paths, like 'C:\\file.tar.gz' or '../test/data/file.tar.gz'
// dest file also support both absolute and relative paths, like 'C:\\' or '../test/data/'
// return err indicate the success or failure function execute
func DeCompressTarGz(src string, dest string) (err error) {
	// open the src tar ball file...
	file, err := os.Open(src)
	if err != nil {
		log.Println("Error open tar gz file:", err)
		return err
	}
	defer file.Close()
	gr, err := gzip.NewReader(file)
	if err != nil {
		log.Println("Error new gzip reader:", err)
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	// loop decompress src list files
	for {
		header, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		name := dest + header.Name
		err = os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
		if err != nil {
			log.Println("Error make dir all:", err)
			return err
		}
		f, err := os.Create(name)
		if err != nil {
			log.Println("Error create name:", err)
			return err
		}
		_, err = io.Copy(f, tr)
		if err != nil {
			log.Println("Error write decompress date:", err)
			return err
		}
	}
	return err
}

// DeCompressTarBz2 function
// input src file list, output dest file path and algorithm which used in unpack, return error info
// this function will base on algorithm to call correspond function
// src file support both absolute and relative paths, like 'C:\\file.tar.bz2' or '../test/data/file.tar.bz2'
// dest file also support both absolute and relative paths, like 'C:\\' or '../test/data/'
// return err indicate the success or failure function execute
func DeCompressTarBz2(src string, dest string) (err error) {
	// open the src tar ball file...
	file, err := os.Open(src)
	if err != nil {
		log.Println("Error open tar bz2 file:", err)
		return err
	}
	defer file.Close()
	br := bzip2.NewReader(file)
	tr := tar.NewReader(br)
	// loop decompress src list files
	for {
		header, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		name := dest + header.Name
		err = os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
		if err != nil {
			log.Println("Error make dir all:", err)
			return err
		}
		f, err := os.Create(name)
		if err != nil {
			log.Println("Error create name:", err)
			return err
		}
		_, err = io.Copy(f, tr)
		if err != nil {
			log.Println("Error write decompress date:", err)
			return err
		}
	}
	return err
}
