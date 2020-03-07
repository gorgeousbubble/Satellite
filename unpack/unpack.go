package unpack

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
)

// Unpack function
// input src file which was packed or encrypted, output dest file path, return error info
// this function will base on algorithm to call correspond function
// src file support both absolute and relative paths, like 'C:\\file.pak' or '../test/data/file.pak'
// dest file also support both absolute and relative paths, like 'C:\\' or '../test/data/'
// algorithm now support 'AES', 'DES', '3DES', 'RSA' and 'BASE64', but you don't need to care it~
// return err indicate the success or failure function execute
func Unpack(src string, dest string) (err error) {
	// first, open the file
	file, err := os.Open(src)
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
		err = UnpackAES(src, dest)
	case "DES", "des":
		err = UnpackDES(src, dest)
	case "3DES", "3des":
		err = Unpack3DES(src, dest)
	case "RSA", "rsa":
		err = UnpackRSA(src, dest)
	case "BASE64", "base64":
		err = UnpackBase64(src, dest)
	default:
		s := fmt.Sprint("Undefined unpack algorithm.")
		err = errors.New(s)
	}
	return err
}

// UnpackConfine function
// unpack file with restrict goroutine(if we do not restrict goroutine, memory will soon be occupied)
// you can adjust confine file and confine buffer when you need change
// other function is same as 'Unpack'
func UnpackConfine(src string, dest string) (err error) {
	// first, open the file
	file, err := os.Open(src)
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
		err = UnpackAESConfine(src, dest)
	case "DES", "des":
		err = UnpackDESConfine(src, dest)
	case "3DES", "3des":
		err = Unpack3DESConfine(src, dest)
	case "RSA", "rsa":
		err = UnpackRSAConfine(src, dest)
	case "BASE64", "base64":
		err = UnpackBase64Confine(src, dest)
	default:
		s := fmt.Sprint("Undefined unpack algorithm.")
		err = errors.New(s)
	}
	return err
}

// UnpackToFile function
// unpack file select target file. If there are many files in package, you can use this function to just decrypt one of them.
// src file support both absolute and relative paths, like 'C:\\file.pak' or '../test/data/file.pak'
// dest file also support both absolute and relative paths, like 'C:\\' or '../test/data/'
// target string is the file which you want to decrypt from package. for instance, if the original name of file is 'capture.png',
// you should fill target segment with 'capture.png'
// return err indicate the success or failure function execute
func UnpackToFile(src string, target string, dest string) (err error) {
	// first, open the file
	file, err := os.Open(src)
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
		err = UnpackAESToFile(src, target, dest)
	case "DES", "des":
		err = UnpackDESToFile(src, target, dest)
	case "3DES", "3des":
		err = Unpack3DESToFile(src, target, dest)
	case "RSA", "rsa":
		err = UnpackRSAToFile(src, target, dest)
	case "BASE64", "base64":
		err = UnpackBase64ToFile(src, target, dest)
	default:
		s := fmt.Sprint("Undefined unpack algorithm.")
		err = errors.New(s)
	}
	return err
}

// UnpackToFileConfine function
// unpack file select target one file with restrict goroutine(if we do not restrict goroutine, memory will soon be occupied)
// you can adjust confine file and confine buffer when you need change
// other function is same as 'UnpackToFile'
func UnpackToFileConfine(src string, target string, dest string) (err error) {
	// first, open the file
	file, err := os.Open(src)
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
		err = UnpackAESToFileConfine(src, target, dest)
	case "DES", "des":
		err = UnpackDESToFileConfine(src, target, dest)
	case "3DES", "3des":
		err = Unpack3DESToFileConfine(src, target, dest)
	case "RSA", "rsa":
		err = UnpackRSAToFileConfine(src, target, dest)
	case "BASE64", "base64":
		err = UnpackBase64ToFileConfine(src, target, dest)
	default:
		s := fmt.Sprint("Undefined unpack algorithm.")
		err = errors.New(s)
	}
	return err
}

// UnpackToMemory function
// unpack file to memory instate of file. It usually use in a situation which need protect data security.
// src file support both absolute and relative paths, like 'C:\\file.pak' or '../test/data/file.pak'
// dest is a slice which used to receive decrypt data. You can send '[]byte' slice address here.
// target string is the file which you want to decrypt from package. for instance, if the original name of file is 'capture.png',
// you should fill target segment with 'capture.png'
// return err indicate the success or failure function execute
func UnpackToMemory(src string, target string, dest *[]byte) (err error) {
	// first, open the file
	file, err := os.Open(src)
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
		err = UnpackAESToMemory(src, target, dest)
	case "DES", "des":
		err = UnpackDESToMemory(src, target, dest)
	case "3DES", "3des":
		err = Unpack3DESToMemory(src, target, dest)
	case "RSA", "rsa":
		err = UnpackRSAToMemory(src, target, dest)
	case "BASE64", "base64":
		err = UnpackBase64ToMemory(src, target, dest)
	default:
		s := fmt.Sprint("Undefined unpack algorithm.")
		err = errors.New(s)
	}
	return err
}

// ExtractInfo function
// This function is mainly used for check verbose information of package.
// src file support both absolute and relative paths, like 'C:\\file.pak' or '../test/data/file.pak'
// dest string slice will return the files name in package.
// sz int slice will return the file number in package.
// algorithm will return which algorithm used by encrypt package.
// return err indicate the success or failure function execute
func ExtractInfo(src string, dest *[]string, sz *[]int, algorithm *string) (err error) {
	// first, open the file
	file, err := os.Open(src)
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
		err = UnpackAESExtractInfo(src, dest, sz)
		*algorithm = "aes"
	case "DES", "des":
		err = UnpackDESExtractInfo(src, dest, sz)
		*algorithm = "des"
	case "3DES", "3des":
		err = Unpack3DESExtractInfo(src, dest, sz)
		*algorithm = "3des"
	case "RSA", "rsa":
		err = UnpackRSAExtractInfo(src, dest, sz)
		*algorithm = "rsa"
	case "BASE64", "base64":
		err = UnpackBase64ExtractInfo(src, dest, sz)
		*algorithm = "base64"
	default:
		s := fmt.Sprint("Undefined unpack algorithm.")
		err = errors.New(s)
	}
	return err
}

func WorkCalculate(src string, algorithm *string, work *int64) (err error) {
	// first, open the file
	file, err := os.Open(src)
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
		*work, err = UnpackAESWorkCalculate(src)
		*algorithm = "AES"
	case "DES", "des":
		*work, err = UnpackDESWorkCalculate(src)
		*algorithm = "DES"
	case "3DES", "3des":
		*work, err = Unpack3DESWorkCalculate(src)
		*algorithm = "3DES"
	case "RSA", "rsa":
		*work, err = UnpackRSAWorkCalculate(src)
		*algorithm = "RSA"
	case "BASE64", "base64":
		*work, err = UnpackBase64WorkCalculate(src)
		*algorithm = "BASE64"
	default:
		s := fmt.Sprint("Undefined unpack algorithm.")
		err = errors.New(s)
	}
	return err
}
