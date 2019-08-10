package cmd

import (
	"strings"
)

type StrSlice []string

func NewStrSlice(s []string, p *[]string) *StrSlice {
	*p = s
	return (*StrSlice)(p)
}

func (s *StrSlice) Set(str string) error {
	str = strings.ReplaceAll(str, " ", "")
	*s = StrSlice(strings.Split(str, ","))
	return nil
}

func (s *StrSlice) String() string {
	*s = StrSlice{}
	return ""
}
