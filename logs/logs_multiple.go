package logs

import (
	"fmt"
	"os"
	"path"
)

type multipleLogWriter struct {
	multipleWriter *os.File
	fileName       string
	maxBytes       int
	curBytes       int
	backupCount    int
}

func (w *multipleLogWriter) Write(b []byte) (n int, err error) {
	w.doRollOver()
	n, err = w.multipleWriter.Write(b)
	w.curBytes += n
	return n, err
}

func (w *multipleLogWriter) Close() error {
	if w.multipleWriter != nil {
		return w.multipleWriter.Close()
	}
	return nil
}

func NewMultipleFileWriter(fileName string, maxBytes int, backupCount int) (mw *multipleLogWriter, err error) {
	dir := path.Dir(fileName)
	err = os.Mkdir(dir, 0777)
	if err != nil {
		return nil, err
	}
	mw = new(multipleLogWriter)
	if maxBytes <= 0 {
		return nil, fmt.Errorf("invalid max bytes")
	}
	mw.fileName = fileName
	mw.maxBytes = maxBytes
	mw.backupCount = backupCount
	mw.multipleWriter, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return mw, err
	}
	file, err := mw.multipleWriter.Stat()
	if err != nil {
		return mw, err
	}
	mw.curBytes = int(file.Size())
	return mw, err
}

func (w *multipleLogWriter) doRollOver() (err error) {
	if w.curBytes < w.maxBytes {
		return fmt.Errorf("curBytes overflow")
	}
	file, err := w.multipleWriter.Stat()
	if err != nil {
		return err
	}
	if w.maxBytes < 0 {
		return fmt.Errorf("invalid max bytes")
	} else if file.Size() < int64(w.maxBytes) {
		w.curBytes = int(file.Size())
		return err
	}
	if w.backupCount > 0 {
		w.multipleWriter.Close()
		for i := w.backupCount - 1; i > 0; i-- {
			sfn := fmt.Sprintf("%s.%d", w.fileName, i)
			dfn := fmt.Sprintf("%s.%d", w.fileName, i+1)
			os.Rename(sfn, dfn)
		}
		dfn := fmt.Sprintf("%s.1", w.fileName)
		os.Rename(w.fileName, dfn)
		w.multipleWriter, err = os.OpenFile(w.fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		w.curBytes = 0
		file, err := w.multipleWriter.Stat()
		if err != nil {
			return err
		}
		w.curBytes = int(file.Size())
	}
}
