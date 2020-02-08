package global

const (
	AppName    = "satellite"
	AppVersion = "v1.00a"
	AppAuthor  = "alopex"
)

const (
	CmdHelp       = "help"
	CmdPacket     = "pack"
	CmdUnpack     = "unpack"
	CmdCompress   = "comp"
	CmdDecompress = "decomp"
	CmdTcp        = "tcp"
	CmdUdp        = "udp"
	CmdHttp       = "http"
	CmdHttps      = "https"
	CmdFtp        = "ftp"
	CmdRpc        = "rpc"
	CmdQRCode     = "qrcode"
	CmdShell      = "shell"
	CmdParses     = "parses"
)

const (
	LogPath   = "/log"
	LogNumber = 5
)

const (
	HttpURLRoot                 = "/"
	HttpURLSatellite            = HttpURLRoot + "satellite"
	HttpURLPack                 = HttpURLSatellite + "/pack"
	HttpURLUnpack               = HttpURLSatellite + "/unpack"
	HttpURLPackProcess          = HttpURLPack + "/p"
	HttpURLUnpackVerbose        = HttpURLUnpack + "/v"
	HttpURLUnpackProcess        = HttpURLUnpack + "/p"
	HttpURLUnpackConfine        = HttpURLUnpack + "/c"
	HttpURLUnpackToFile         = HttpURLUnpack + "/f"
	HttpURLUnpackToFileConfine  = HttpURLUnpack + "/cf"
	HttpURLUnpackToMemory       = HttpURLUnpack + "/m"
	HttpURLComp                 = HttpURLSatellite + "/comp"
	HttpURLDecomp               = HttpURLSatellite + "/decomp"
	HttpURLImages               = HttpURLSatellite + "/images"
	HttpURLImagesQRCode         = HttpURLImages + "/qrcode"
	HttpURLImagesQRCodeToFile   = HttpURLImagesQRCode + "/f"
	HttpURLImagesQRCodeToMemory = HttpURLImagesQRCode + "/m"
	HttpURLParses               = HttpURLSatellite + "/parses"
	HttpURLParsesIni            = HttpURLParses + "/ini"
)

const (
	AESBufferSize    = 128  // AES buffer size should be 128, 256, ...
	DESBufferSize    = 128  // DES buffer size should be 128, 256, ...
	RSAPacketSize    = 64   // RSA buffer size should less than 128(Packet)
	RSAUnpackSize    = 128  // RSA buffer size(Unpack)
	Base64BufferSize = 128  // Base64 buffer size
	ConfineFiles     = 5    // Confine go-routine concurrent files
	ConfineBuffers   = 8192 // Confine go-routine concurrent buffers
)

const (
	TCPBufferSize    = 4096  // TCP buffer size(Receive)
	UDPBufferSize    = 4096  // UDP buffer size(Receive)
	HTTPWriteTimeout = 10000 // HTTP Write Timeout Time(Millisecond)
	HTTPReadTimeout  = 10000 // HTTP Read Timeout Time(Millisecond)
)

const (
	NetHttpTimeout = 100 // Net HTTP timeout(100ms)
)
