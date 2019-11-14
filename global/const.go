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
	CmdFtp        = "ftp"
	CmdQRCode     = "qrcode"
)

const (
	LogPath   = "/log"
	LogNumber = 5
)

const (
	HttpURLRoot           = "/"
	HttpURLSatellite      = HttpURLRoot + "satellite"
	HttpURLPack           = HttpURLSatellite + "/pack"
	HttpURLUnpack         = HttpURLSatellite + "/unpack"
	HttpURLPackProcess    = HttpURLPack + "/p"
	HttpURLUnpackVerbose  = HttpURLUnpack + "/v"
	HttpURLUnpackProcess  = HttpURLUnpack + "/p"
	HttpURLUnpackToFile   = HttpURLUnpack + "/f"
	HttpURLUnpackToMemory = HttpURLUnpack + "/m"
	HttpURLComp           = HttpURLSatellite + "/comp"
	HttpURLDecomp         = HttpURLSatellite + "/decomp"
	HttpURLImages         = HttpURLSatellite + "/images"
	HttpURLImagesQRCode   = HttpURLImages + "/qrcode"
)

const (
	AESBufferSize    = 128 // AES buffer size should be 128, 256, ...
	DESBufferSize    = 128 // DES buffer size should be 128, 256, ...
	RSAPacketSize    = 64  // RSA buffer size should less than 128(Packet)
	RSAUnpackSize    = 128 // RSA buffer size(Unpack)
	Base64BufferSize = 128 // Base64 buffer size
)

const (
	TCPBufferSize    = 4096  // TCP buffer size(Receive)
	UDPBufferSize    = 4096  // UDP buffer size(Receive)
	HTTPWriteTimeout = 10000 // HTTP Write Timeout Time(Millisecond)
	HTTPReadTimeout  = 10000 // HTTP Read Timeout Time(Millisecond)
)
