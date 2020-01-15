package nets

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

func StartRpcHttpServer(ip string, port string) {
	// rpc register interface...
	err := rpc.Register(new(GoApi))
	if err != nil {
		fmt.Println("Error RPC register interface:", err)
		log.Println("Error RPC register interface:", err)
		os.Exit(1)
	}
	// rpc handle http
	rpc.HandleHTTP()
	// rpc start http service...
	l, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Error listen tcp:", err)
		log.Println("Error listen tcp:", err)
		os.Exit(1)
	}
	fmt.Println("Start Listen And Server on ", ip+":"+port)
	// start http service...
	err = http.Serve(l, nil)
	if err != nil {
		fmt.Println("Error serve http:", err)
		log.Println("Error serve http:", err)
		os.Exit(1)
	}
}

func StartRpcTcpServer(ip string, port string) {
	// rpc register interface...
	err := rpc.Register(new(GoApi))
	if err != nil {
		fmt.Println("Error RPC register interface:", err)
		log.Println("Error RPC register interface:", err)
		os.Exit(1)
	}
	// rpc start tcp service...
	l, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Error listen tcp:", err)
		log.Println("Error listen tcp:", err)
		os.Exit(1)
	}
	fmt.Println("Start Listen And Server on ", ip+":"+port)
	for {
		// client connect to server...
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
}
