package decomp

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"strings"
)

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
