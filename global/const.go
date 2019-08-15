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
)

const (
	LogPath   = "/log"
	LogNumber = 5
)

const (
	ConstAESBufferSize    = 128 // AES buffer size should be 128, 256, ...
	ConstDESBufferSize    = 128 // DES buffer size should be 128, 256, ...
	ConstRSAPacketSize    = 64  // RSA buffer size should less than 128(Packet)
	ConstRSAUnpackSize    = 128 // RSA buffer size(Unpack)
	ConstBase64BufferSize = 128 // Base64 buffer size
)

const (
	ConstTCPBufferSize = 4096 // TCP buffer size(Receive)
)
