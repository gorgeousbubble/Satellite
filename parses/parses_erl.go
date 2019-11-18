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
	// check element without '{}' or '[]'
	if bytes.ContainsAny(r, "{&}&[&]") {
		err = errors.New("bytes may extract list or tuple first")
		return r, sub, err
	}
	sub = bytes.TrimPrefix(s, r)
	sub = bytes.TrimPrefix(sub, []byte(","))
	return r, sub, err
}

func extractOneList(s []byte) (r []byte, sub []byte, err error) {
	// erlang bytes contains '[]' for list
	if !bytes.Contains(s, []byte("[")) || !bytes.Contains(s, []byte("]")) {
		err = errors.New("bytes don't have list")
		return r, sub, err
	}
	// check '[' & ']' numbers should equal
	if bytes.Count(s, []byte("[")) != bytes.Count(s, []byte("]")) {
		err = errors.New("bytes list illegal, list may not intergrity")
		return r, sub, err
	}
	// check '[' & ']' sequence
	if bytes.Index(s, []byte("[")) > bytes.Index(s, []byte("]")) {
		err = errors.New("bytes list illegal, sequence of list may not right")
		return r, sub, err
	}
	// extract erlang list
	count := 0
	start := bytes.Index(s, []byte("["))
	for i := start; i < len(s); i++ {
		if s[i] == '[' {
			count++
		} else if s[i] == ']' {
			count--
		}
		r = append(r, s[i])
		if count == 0 {
			break
		}
	}
	sub = bytes.TrimPrefix(s, r)
	sub = bytes.TrimPrefix(sub, []byte(","))
	return r, sub, err
}

func extractOneTuple(s []byte) (r []byte, sub []byte, err error) {
	// erlang bytes contains '{}' for tuple
	if !bytes.Contains(s, []byte("{")) || !bytes.Contains(s, []byte("}")) {
		err = errors.New("bytes don't have tuple")
		return r, sub, err
	}
	// check '{' & '}' numbers should equal
	if bytes.Count(s, []byte("{")) != bytes.Count(s, []byte("}")) {
		err = errors.New("bytes tuple illegal, tuple may not intergrity")
		return r, sub, err
	}
	// check '{' & '}' sequence
	if bytes.Index(s, []byte("{")) > bytes.Index(s, []byte("}")) {
		err = errors.New("bytes tuple illegal, sequence of tuple may not right")
		return r, sub, err
	}
	// extract erlang tuple
	count := 0
	start := bytes.Index(s, []byte("{"))
	for i := start; i < len(s); i++ {
		if s[i] == '{' {
			count++
		} else if s[i] == '}' {
			count--
		}
		r = append(r, s[i])
		if count == 0 {
			break
		}
	}
	sub = bytes.TrimPrefix(s, r)
	sub = bytes.TrimPrefix(sub, []byte(","))
	return r, sub, err
}
