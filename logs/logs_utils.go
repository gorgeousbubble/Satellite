package logs

type Handler interface {
	Write(b []byte) (n int, err error)
	Close() error
}
