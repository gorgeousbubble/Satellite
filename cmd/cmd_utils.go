package cmd

import (
	"strings"
)

type StrSlice []string

// NewStrSlice function
// this function is mainly used to allow go package flag create a string slice
// usage: packCmd.Var(NewStrSlice([]string{}, &packSrc)...)
// return *StrSlice
func NewStrSlice(s []string, p *[]string) *StrSlice {
	*p = s
	return (*StrSlice)(p)
}

// StrSlice Set function
func (s *StrSlice) Set(str string) error {
	str = strings.ReplaceAll(str, " ", "")
	*s = StrSlice(strings.Split(str, ","))
	return nil
}

// StrSlice String function
func (s *StrSlice) String() string {
	*s = StrSlice{}
	return ""
}
