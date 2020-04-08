package comp

import (
	"errors"
	"fmt"
)

// Compress function
// input src file list, output dest file path and algorithm which used in pack, return error info
// this function will base on algorithm to call correspond function
// src file support both absolute and relative paths, like 'C:\\file.txt' or '../test/data/file.txt'
// dest file also support both absolute and relative paths, like 'C:\\package.pak' or '../test/data/package.pak'
// algorithm now support 'tar', 'tar.gz', 'zip', you can send both up case and low case
// return err indicate the success or failure function execute
func Compress(src []string, dest string, algorithm string) (err error) {
	switch algorithm {
	case "tar":
		err = CompressTar(src, dest)
	case "tar.gz":
		err = CompressTarGz(src, dest)
	case "zip":
		err = CompressZip(src, dest)
	default:
		s := fmt.Sprint("Undefined compress algorithm.")
		err = errors.New(s)
	}
	return err
}
