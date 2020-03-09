package parse

func Unmarshal(path string, out *TAutoExec) (err error) {
	return unmarshalAutoExec(path, out, "json")
}
