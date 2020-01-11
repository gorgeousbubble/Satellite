package nets

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
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
