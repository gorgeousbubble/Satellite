package comp

import (
	"errors"
	"fmt"
)

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
