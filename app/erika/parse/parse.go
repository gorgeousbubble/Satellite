package parse

func Unmarshal(path string, out interface{}) (err error) {
	return err
}

func UnmarshalAutoExec(path string, out *TAutoExec) (err error) {
	return unmarshalAutoExec(path, out, "json")
}

func UnmarshalAutoMonitor(path string, out *TAutoMonitor) (err error) {
	return unmarshalAutoMonitor(path, out, "json")
}
