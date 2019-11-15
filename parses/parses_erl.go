package parses

import (
	"bytes"
	"errors"
)

func extractOneElement(s []byte) (r []byte, sub []byte, err error) {
	// erlang bytes contains ',' for element
	if bytes.Contains(s, []byte(",")) {
		err = errors.New("bytes don't have element")
		return r, sub, err
	}
	// find element by ','
	index := bytes.Index(s, []byte(","))
	for i := 0; i < index; i++ {
		r = append(r, s[i])
	}
	if bytes.ContainsAny(r, "{&}&[&]") {
		err = errors.New("bytes may extract list or tuple first")
		return r, sub, err
	}
	sub = bytes.TrimPrefix(s, r)
	sub = bytes.TrimPrefix(sub, []byte(","))
	return r, sub, err
}
