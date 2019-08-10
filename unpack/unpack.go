package unpack

import (
	"bufio"
	"bytes"
	"github.com/pkg/errors"
	"log"
	"os"
)

func Unpack(srcfile string, destpath string) (err error) {
	// first, open the file
	file, err := os.Open(srcfile)
	if err != nil {
		log.Println("Error open file:", err)
		return err
	}
	// second, read file data
	buf := make([]byte, 60)
	rd := bufio.NewReader(file)
	_, err = rd.Read(buf)
	if err != nil {
		log.Println("Error read file:", err)
		return err
	}
	// third, close the file
	err = file.Close()
	if err != nil {
		log.Println("Error close file:", err)
		return err
	}
	// fourth, find the algorithm
	buf = buf[48:56]
	index := bytes.IndexByte(buf, 0)
	tp := string(buf[0:index])
	switch tp {
	case "AES", "aes":
		err = UnpackAES(srcfile, destpath)
	case "DES", "des":
		err = UnpackDES(srcfile, destpath)
	case "3DES", "3des":
		err = Unpack3DES(srcfile, destpath)
	case "RSA", "rsa":
		err = UnpackRSA(srcfile, destpath)
	case "BASE64", "base64":
		err = UnpackBase64(srcfile, destpath)
	default:
		err = errors.New("Undefined unpack algorithm.")
		log.Println(err)
	}
	return err
}
