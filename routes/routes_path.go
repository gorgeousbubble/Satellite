package routes

import "strings"

// trim the extend name of file name
// example: file.txt -> file
func TrimSuffixPoint(name string) string {
	// try to find suffix point
	n := strings.Index(name, ".")
	if n == -1 {
		return name
	}
	short := name[0:n]
	return short
}

// get the extend name of file name
// example: file.txt -> txt
func GetSuffixPoint(name string) string {
	// try to find suffix point
	n := strings.Index(name, ".")
	if n == -1 {
		return ""
	}
	short := name[n+1:]
	return short
}

// trim the last slash of path name
// example: ../test/data/ -> ../test/data
func TrimSuffixSlash(path string) string {
	// try to find suffix slash
	if path[len(path)-1] == '/' {
		return path[0 : len(path)-1]
	} else if path[len(path)-1] == '\\' {
		return path[0 : len(path)-2]
	}
	return path
}
