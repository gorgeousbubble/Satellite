package nets

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
)

// StartFtpClient function
// this function is mainly used to start ftp client
// host is the ip address of host machine
// user is the user name of ftp server
// pass is the password of ftp server
// if connect successful it will return the ftp server connection handle, else it will be nil
// return err indicate the success or failure function execute
func StartFtpClient(host string, user string, pass string) (c *ftp.ServerConn, err error) {
	// connect to ftp server
	c, err = ftp.Dial(host+":21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		fmt.Println("Error connect to ftp server:", err)
		log.Println("Error connect to ftp server:", err)
		return c, err
	}
	// login ftp server
	err = c.Login(user, pass)
	if err != nil {
		fmt.Println("Error login ftp server:", err)
		log.Println("Error login ftp server:", err)
		return c, err
	}
	return c, err
}

// StopFtpClient function
// this function is mainly used to stop the connection with ftp
// return err indicate the success or failure function execute
func StopFtpClient(c *ftp.ServerConn) (err error) {
	err = c.Quit()
	if err != nil {
		fmt.Println("Error quit ftp server:", err)
		log.Println("Error quit ftp server:", err)
		return err
	}
	return err
}

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
