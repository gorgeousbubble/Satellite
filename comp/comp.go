package comp

import (
	"github.com/pkg/errors"
)

func Compress(src []string, dest string, algorithm string) (err error) {
	switch algorithm {
	case "tar":
		err = CompressTarGz(src, dest)
	case "zip":
		err = CompressZip(src, dest)
	default:
		err = errors.New("Undefined compress algorithm.")
	}
	return err
}
