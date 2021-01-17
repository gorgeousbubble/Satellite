package logs

import (
	"os"
	"path"
)

type fileLogWriter struct {
	fileWriter *os.File
}

func (w *fileLogWriter) Write(b []byte) (n int, err error) {
	return w.fileWriter.Write(b)
}

func (w *fileLogWriter) Close() error {
	return w.fileWriter.Close()
}

func NewFileWriter(name string, flag int) (fw *fileLogWriter, err error) {
	dir := path.Dir(name)
	// check file or path exist
	_, err = os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(dir, 0777)
			if err != nil {
				return nil, err
			}
		}
	}
	// create or open file
	file, err := os.OpenFile(name, flag, 0)
	if err != nil {
		return nil, err
	}
	fw = new(fileLogWriter)
	fw.fileWriter = file
	return fw, err
}
