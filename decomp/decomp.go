package decomp

import (
	"errors"
	"fmt"
)

// DeCompress function
// input src file list, output dest file path and algorithm which used in unpack, return error info
// this function will base on algorithm to call correspond function
// src file support both absolute and relative paths, like 'C:\\file.tar.gz' or '../test/data/file.tar.gz'
// dest file also support both absolute and relative paths, like 'C:\\' or '../test/data/'
// algorithm now support 'tar', 'tar.gz', 'zip', you can send both up case and low case
// return err indicate the success or failure function execute
func DeCompress(src string, dest string, algorithm string) (err error) {
	switch algorithm {
	case "tar":
		err = DeCompressTar(src, dest)
	case "tar.gz":
		err = DeCompressTarGz(src, dest)
	case "zip":
		err = DeCompressZip(src, dest)
	default:
		s := fmt.Sprint("Undefined decompress algorithm.")
		err = errors.New(s)
	}
	return err
}
