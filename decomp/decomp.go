package decomp

import (
	"errors"
	"fmt"
)

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
