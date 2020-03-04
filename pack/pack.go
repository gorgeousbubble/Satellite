package pack

import (
	"errors"
	"fmt"
)

// Pack function
// input src file list, output dest file path and algorithm which used in pack, return error info
// this function will base on algorithm to call correspond function
// src file support both absolute and relative paths, like 'C:\\file.txt' or '../test/data/file.txt'
// dest file also support both absolute and relative paths, like 'C:\\package.pak' or '../test/data/package.pak'
// algorithm now support 'AES', 'DES', '3DES', 'RSA' and 'BASE64', you can send both up case and low case
// return err indicate the success or failure function execute
func Pack(src []string, dest string, algorithm string) (err error) {
	switch algorithm {
	case "AES", "aes":
		err = PackAES(src, dest)
	case "DES", "des":
		err = PackDES(src, dest)
	case "3DES", "3des":
		err = Pack3DES(src, dest)
	case "RSA", "rsa":
		err = PackRSA(src, dest)
	case "BASE64", "base64":
		err = PackBase64(src, dest)
	default:
		s := fmt.Sprint("Undefined pack algorithm.")
		err = errors.New(s)
	}
	return err
}

// WorkCalculate function
// input src file list, algorithm which used in pack and output work value, return error info
// this function will called by calculate work
// algorithm now support 'AES', 'DES', '3DES', 'RSA' and 'BASE64', you can send both up case and low case
// work value is total work force that will be done
// return err indicate the success or failure function execute
func WorkCalculate(src []string, algorithm string, work *int64) (err error) {
	switch algorithm {
	case "AES", "aes":
		*work, err = PackAESWorkCalculate(src)
	case "DES", "des":
		*work, err = PackDESWorkCalculate(src)
	case "3DES", "3des":
		*work, err = PackDESWorkCalculate(src)
	case "RSA", "rsa":
		*work, err = PackRSAWorkCalculate(src)
	case "BASE64", "base64":
		*work, err = PackBase64WorkCalculate(src)
	default:
		s := fmt.Sprint("Undefined pack algorithm.")
		err = errors.New(s)
	}
	return err
}
