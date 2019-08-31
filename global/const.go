package global

const (
	AppName    = "satellite"
	AppVersion = "v1.00a"
	AppAuthor  = "alopex"
)

const (
	CmdPacket     = "pack"
	CmdUnpack     = "unpack"
	CmdCompress   = "comp"
	CmdDecompress = "decomp"
	CmdTcp        = "tcp"
	CmdUdp        = "udp"
	CmdHttp       = "http"
	CmdHttps      = "https"
)

const (
	LogPath   = "/log"
	LogNumber = 5
)

const (
	HttpURLRoot = "/"
)

const (
	AESBufferSize    = 128 // AES buffer size should be 128, 256, ...
	DESBufferSize    = 128 // DES buffer size should be 128, 256, ...
	RSAPacketSize    = 64  // RSA buffer size should less than 128(Packet)
	RSAUnpackSize    = 128 // RSA buffer size(Unpack)
	Base64BufferSize = 128 // Base64 buffer size
)

const (
	ConstTCPBufferSize    = 4096 // TCP buffer size(Receive)
	ConstUDPBufferSize    = 4096 // UDP buffer size(Receive)
	ConstHTTPWriteTimeout = 1000 // HTTP Write Timeout Time(Millisecond)
	ConstHTTPReadTimeout  = 1000 // HTTP Read Timeout Time(Millisecond)
)
