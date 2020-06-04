# Nets package interfaces
The Nets package function interfaces description.

## Introduction
Nets package is a functional package. It can offer different kind of network connection drive interfaces. You can use this package to realize network connection, e-main send and receive, HTTP/HTTPS service and client, and so on. Nets package based on the go 'net' package, it support multiple network protocols such as TCP, UDP HTTP, HTTPS, FTP, RPC, SMTP, POP3, etc.

## Feature of package
The package is mainly used for network connection which has been realized by 'nets' package. You can refer to corresponding go file.

#### Network service and useful tools
  * The nets package offer interfaces let program easy to realize network connection
  * Support TCP/UDP server and client
  * Support HTTP/HTTPS server and client
  * Support FTP service
  * support RPC service
  * Support SMTP/POP3/GMAIL send and receive e-mail
  * Support interfaces to allow external application to interactive with Satellite
  * Useful and effective

#### HTTP/HTTPS service support application API
  * Encrypt/Decrypt files
  * Pack/Unpack files
  * Compress/Decompress files
  * QR-Code Generator
  * Parse documents

## External definitions
There are some constant defined in file 'global/const.go', mainly about HTTP/HTTPS service URI. You can redefine it in 'nets' package it you want to use this package independence.

## Usage of interfaces
There are many interfaces in nets package. You can refer to corresponding golang file in order to find futher more function.

#### Start TCP Service
  * Start TCP Server
  ```batch
  func TestStartTcpServer(t *testing.T) {
	ip := "127.0.0.1"
	port := "11514"
	// skip this test case
	t.Skip("IGNORE: TestStartTcpServer")
	// start tcp server goroutine
	go StartTcpServer(ip, port)
	// start tcp client...
	StartTcpClient(ip, port)
  }
  ```

  * Start TCP Client
  ```batch
  func TestStartTcpClient(t *testing.T) {
	ip := "127.0.0.1"
	port := "11515"
	// skip this test case
	t.Skip("IGNORE: TestStartTcpClient")
	// start tcp server goroutine
	go StartTcpServer(ip, port)
	// start tcp client...
	StartTcpClient(ip, port)
  }
  ```

#### Start UDP Service
  * Start UDP Server
  ```batch
  func TestStartUdpServer(t *testing.T) {
	ip := "127.0.0.1"
	port := "12514"
	// skip this test case
	t.Skip("IGNORE: TestStartUdpServer")
	// start udp server goroutine
	go StartUdpServer(ip, port)
	// start udp client...
	StartUdpClient(ip, port)
  }
  ```

  * Start UDP Client
  ```batch
  func TestStartUdpClient(t *testing.T) {
	ip := "127.0.0.1"
	port := "12515"
	// skip this test case
	t.Skip("IGNORE: TestStartUdpClient")
	// start udp server goroutine
	go StartUdpServer(ip, port)
	// start udp client...
	StartUdpClient(ip, port)
  }
  ```

#### Start HTTP/HTTPS Service
  * Start HTTP Server
  ```batch
  func TestStartHttpServer(t *testing.T) {
	ip := "127.0.0.1"
	port := "12514"
	// skip this test case
	t.Skip("IGNORE: TestStartHttpServer")
	// start http server
	StartHttpServer(ip, port)
  }
  ```

  * Start HTTPS Server
  ```batch
  func TestStartHttpsServer(t *testing.T) {
	ip := "127.0.0.1"
	port := "12514"
	// skip this test case
	t.Skip("IGNORE: TestStartHttpsServer")
	// start https server
	StartHttpsServer(ip, port)
  }
  ```

  * Start HTTP/HTTPS Client  
  for HTTP client, you can use the interfaces in 'http' go package.  
  for HTTPS client, you should define a httpClient first and modify it's configure like this:
  ```batch
  var httpsClient *http.Client

  func init() {
	 tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	httpsClient = &http.Client{
		Transport: tr,
	}
  }
  ```
  
  #### Start FTP Service
  * Start FTP Server
  ```batch
  func TestStartFtpServer(t *testing.T) {
	ip := "127.0.0.1"
	port := "10021"
	// skip this test case
	t.Skip("IGNORE: TestStartFtpServer")
	// start ftp server
	go StartFtpServer(ip, port)
	// start ftp client
	r, err := http.Get("http://" + ip + ":" + port)
	if err != nil {
		t.Errorf("Error GET ftp server url:%v", err)
	}
	if r.StatusCode != http.StatusOK {
		t.Errorf("Error GET ftp server response")
	}
  }
  ```

  * Start FTP Client
  ```batch
  // UploadFile function
  // this function is mainly used to Upload the file to ftp server
  // input ServerConn
  // path indicate the file path which you want to upload
  // return err indicate the success or failure function execute
  func UploadFile(c *ftp.ServerConn, path string, r io.Reader) (err error) {
	err = c.Stor(path, r)
	if err != nil {
		fmt.Println("Error upload file:", err)
		log.Println("Error upload file:", err)
		panic(err)
	}
	return err
  }

  // DownloadFile function
  // this function is mainly used to Download the file from ftp server
  // input ServerConn
  // path indicate the file path which you want to upload
  // return err indicate the success or failure function execute
  func DownloadFile(c *ftp.ServerConn, path string) (r io.Reader, err error) {
	r, err = c.Retr(path)
	if err != nil {
		fmt.Println("Error download file:", err)
		log.Println("Error download file:", err)
		panic(err)
	}
	return r, err
  }
  ```
