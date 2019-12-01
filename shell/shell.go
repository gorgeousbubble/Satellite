package shell

import (
	"errors"
	"fmt"
)

func Shell(src string, dest string, algorithm string) (err error) {
	switch algorithm {
	case "upx":
		err = shellUpx(src, dest)
	default:
		s := fmt.Sprint("Undefined shell algorithm.")
		err = errors.New(s)
	}
	return err
}
