package logs

type LogWriter interface {
	Write(b []byte) (n int, err error)
	Close() error
}
