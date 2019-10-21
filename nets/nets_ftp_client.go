package nets

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
)

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

func StopFtpClient(c *ftp.ServerConn) (err error) {
	err = c.Quit()
	if err != nil {
		fmt.Println("Error quit ftp server:", err)
		log.Println("Error quit ftp server:", err)
		return err
	}
	return err
}

func UploadFile(c *ftp.ServerConn, path string, r io.Reader) (err error) {
	err = c.Stor(path, r)
	if err != nil {
		fmt.Println("Error upload file:", err)
		log.Println("Error upload file:", err)
		panic(err)
	}
	return err
}

func DownloadFile(c *ftp.ServerConn, path string) (r io.Reader, err error) {
	r, err = c.Retr(path)
	if err != nil {
		fmt.Println("Error download file:", err)
		log.Println("Error download file:", err)
		panic(err)
	}
	return r, err
}
