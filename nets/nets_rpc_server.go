package nets

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func StartRpcHttpServer(ip string, port string) (err error) {
	// rpc register interface...
	err = rpc.Register(new(GoApi))
	if err != nil {
		log.Println("Error RPC register interface:", err)
		return err
	}
	// rpc handle http
	rpc.HandleHTTP()
	// rpc start http service...
	l, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		log.Println("Error listen tcp:", err)
		return err
	}
	fmt.Println("Start Listen And Server on ", ip+":"+port)
	// start http service...
	err = http.Serve(l, nil)
	if err != nil {
		log.Println("Error serve http:", err)
		return err
	}
	return err
}

func StartRpcTcpServer(ip string, port string) (err error) {
	// rpc register interface...
	err = rpc.Register(new(GoApi))
	if err != nil {
		log.Println("Error RPC register interface:", err)
		return err
	}
	// rpc start tcp service...
	l, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		log.Println("Error listen tcp:", err)
		return err
	}
	fmt.Println("Start Listen And Server on ", ip+":"+port)
	// rpc loop wait for connect...
	go func(l net.Listener) {
		for {
			// client connect to server
			conn, err := l.Accept()
			if err != nil {
				continue
			}
			// handle json transfer
			go func(conn net.Conn) {
				jsonrpc.ServeConn(conn)
				fmt.Println("RPC client connect to server:", conn.RemoteAddr())
			}(conn)
		}
	}(l)
	return err
}
