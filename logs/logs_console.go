package logs

import "io"

type consoleLogWriter struct {
	consoleWriter io.Writer
}

func (w *consoleLogWriter) Write(b []byte) (n int, err error) {
	return w.consoleWriter.Write(b)
}

func (w *consoleLogWriter) Close() error {
	return nil
}

func NewConsoleWriter(w io.Writer) (cw *consoleLogWriter, err error) {
	cw = new(consoleLogWriter)
	cw.consoleWriter = w
	return cw, err
}
