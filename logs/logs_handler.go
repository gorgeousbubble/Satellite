package logs

import "io"

type StreamHandler struct {
	w io.Writer
}

type NullHandler struct {
}

func NewStreamHandler(w io.Writer) (*StreamHandler, error) {
	h := new(StreamHandler)
	h.w = w
	return h, nil
}

func NewNullHandler() (*NullHandler, error) {
	return new(NullHandler), nil
}

func (h *StreamHandler) Write(b []byte) (n int, err error) {
	return h.w.Write(b)
}

func (h *StreamHandler) Close() error {
	return nil
}

func (h *NullHandler) Write(b []byte) (n int, err error) {
	return len(b), nil
}

func (h *NullHandler) close() error {
	return nil
}
