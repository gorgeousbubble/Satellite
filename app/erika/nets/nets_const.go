package nets

const (
	HTTPURLRoot  = "erika/v1"
	HTTPURLHello = HTTPURLRoot + "/hello"
)

const (
	TCPBufferSize    = 4096  // TCP buffer size(Receive)
	UDPBufferSize    = 4096  // UDP buffer size(Receive)
	HTTPWriteTimeout = 10000 // HTTP Write Timeout Time(Millisecond)
	HTTPReadTimeout  = 10000 // HTTP Read Timeout Time(Millisecond)
)

const (
	NetHttpTimeout = 600 // Net HTTP timeout(100ms)
)
