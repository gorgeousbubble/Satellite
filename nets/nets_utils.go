package nets

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	. "satellite/global"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func connTcpSendHandler(c net.Conn) {
	// create stdin reader
	rd := bufio.NewReader(os.Stdin)
	// loop waiting for send message...
	for {
		// read stdin data stream
		in, err := rd.ReadString('\n')
		if err != nil {
			fmt.Println("Error read stdin data stream:", err)
			log.Println("Error read stdin data stream:", err)
			break
		}
		// handle data stream
		in = strings.TrimSpace(in)
		// send data stream
		_, err = c.Write([]byte(in))
		if err != nil {
			fmt.Println("Error write data stream:", err)
			log.Println("Error write data stream:", err)
			break
		}
		t := time.Now()
		fmt.Println("["+c.LocalAddr().String()+"] ", t.Format("15:04:05"))
		fmt.Println("Remote<-Local:", in)
	}
}

func connTcpRecvHandler(c net.Conn) {
	// invalid socket
	if c == nil {
		fmt.Println("Invalid socket connect.")
		log.Println("Invalid socket connect.")
		return
	}
	// create slice to receive data
	buf := make([]byte, TCPBufferSize)
	// loop receive data stream
	for {
		// start read data stream from net
		n, err := c.Read(buf)
		if n == 0 || err != nil {
			if err != nil {
				fmt.Println("Remote host forcibly closed connect:", c.RemoteAddr().String())
				log.Println("Error read data stream:", err)
			}
			err = c.Close()
			if err != nil {
				fmt.Println("Error close socket:", err)
				log.Println("Error close socket:", err)
			}
			break
		}
		// handle data stream
		str := strings.TrimSpace(string(buf[0:n]))
		t := time.Now()
		fmt.Println("["+c.RemoteAddr().String()+"] ", t.Format("15:04:05"))
		fmt.Println("Remote->Local:", str)
	}
}

func connUdpSendHandler(c *net.UDPConn) {
	// create stdin reader
	rd := bufio.NewReader(os.Stdin)
	// loop waiting for send message...
	for {
		// read stdin data stream
		in, err := rd.ReadString('\n')
		if err != nil {
			fmt.Println("Error read stdin data stream:", err)
			log.Println("Error read stdin data stream:", err)
			break
		}
		// handle data stream
		in = strings.TrimSpace(in)
		// send data stream
		_, err = c.Write([]byte(in))
		if err != nil {
			fmt.Println("Error write data stream:", err)
			log.Println("Error write data stream:", err)
			break
		}
		t := time.Now()
		fmt.Println("["+c.LocalAddr().String()+"] ", t.Format("15:04:05"))
		fmt.Println("Remote<-Local:", in)
	}
}

func connUdpRecvHandler(c *net.UDPConn) {
	// invalid socket
	if c == nil {
		fmt.Println("Invalid socket connect.")
		log.Println("Invalid socket connect.")
		return
	}
	// create slice to receive data
	buf := make([]byte, UDPBufferSize)
	// loop receive data stream
	for {
		// start read data stream from net
		n, remoteAddr, err := c.ReadFromUDP(buf)
		if n == 0 || err != nil {
			if err != nil {
				fmt.Println("Error read data stream:", err)
				log.Println("Error read data stream:", err)
			}
			err = c.Close()
			if err != nil {
				fmt.Println("Error close socket:", err)
				log.Println("Error close socket:", err)
			}
			break
		}
		// handle data stream
		str := strings.TrimSpace(string(buf[0:n]))
		t := time.Now()
		fmt.Println("["+remoteAddr.String()+"] ", t.Format("15:04:05"))
		fmt.Println("Remote->Local:", str)
	}
}

func createHttpRouter() (r *mux.Router) {
	r = mux.NewRouter()
	r.HandleFunc(HttpURLRoot, handleRoot).Methods("GET")
	return r
}

func GenerateCA(ip string) (err error) {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, max)
	if err != nil {
		return err
	}
	subject := pkix.Name{
		Organization: []string{"Team Gorgeous Bubble"},
		CommonName:   "Https Service",
	}
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP(ip)},
	}
	pk, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	if err != nil {
		return err
	}
	certOut, err := os.Create("cert.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if err != nil {
		return err
	}
	certOut.Close()
	keyOut, err := os.Create("key.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	if err != nil {
		return err
	}
	keyOut.Close()
	return err
}
