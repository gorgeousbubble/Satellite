package parses

import (
	"bytes"
	"fmt"
)

func getValue(in []byte, section string, key string) (r []byte) {
	// split data by '\n'
	var s [][]byte
	s1 := bytes.Split(in, []byte("\n"))
	for _, v := range s1 {
		fmt.Println(string(v))
		// delete comments
		if bytes.Contains(v, []byte(";")) {
			index := bytes.Index(v, []byte(";"))
			v = append(v[:index], v[len(v):]...)
		}
		// delete all space
		v = bytes.ReplaceAll(v, []byte(" "), []byte(""))
		// delete all return
		v = bytes.ReplaceAll(v, []byte("\r"), []byte(""))
		// delete all blank lines, packet to s slice
		if !bytes.Equal(v, []byte("")) {
			s = append(s, v)
		}
	}
	// seek for section
	var name []byte
	for _, v := range s {
		fmt.Println(string(v))
		// section start with '[' and end with ']'
		if v[0] == '[' && v[len(v)-1] == ']' {
			name = v[1 : len(v)-1]
		} else if string(name) == section {
			// seek for key
			pairs := bytes.Split(v, []byte("="))
			if len(pairs) == 2 {
				// key value before '='
				k := bytes.TrimSpace(pairs[0])
				if string(k) == key {
					r = bytes.TrimSpace(pairs[1])
					return r
				}
			}
		}
	}
	r = []byte("")
	return r
}
