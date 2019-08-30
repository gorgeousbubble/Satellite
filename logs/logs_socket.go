package logs

import (
	"encoding/binary"
	"net"
	"time"
)

type socketLogWriter struct {
	socket   net.Conn
	protocol string
	addr     string
}

func (w *socketLogWriter) Write(b []byte) (n int, err error) {
	if err = w.connect(); err != nil {
		return
	}
	buf := make([]byte, len(b)+4)
	binary.BigEndian.PutUint32(buf, uint32(len(b)))
	copy(buf[4:], b)
	n, err = w.socket.Write(buf)
	if err != nil {
		w.socket.Close()
		w.socket = nil
	}
	return
}

func (w *socketLogWriter) Close() error {
	if w.socket != nil {
		w.socket.Close()
	}
	return nil
}

func (w *socketLogWriter) connect() (err error) {
	if w.socket != nil {
		return err
	}
	w.socket, err = net.DialTimeout(w.protocol, w.addr, 5*time.Second)
	if err != nil {
		return err
	}
	return err
}

func NewSocketWriter(protocal string, addr string) (sw *socketLogWriter, err error) {
	sw = new(socketLogWriter)
	sw.protocol = protocal
	sw.addr = addr
	return sw, err
}
