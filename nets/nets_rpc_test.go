package nets

import (
	"fmt"
	"testing"
)

func TestCallRpcApiMD5Encode(t *testing.T) {
	ip := "127.0.0.1"
	port := "12000"
	request := TNetsRpcPackMD5EncodeReq{Src: "Satellite"}
	response := TNetsRpcPackMD5EncodeResp{}
	// start rpc http server...
	go StartRpcHttpServer(ip, port)
	// start rpc http client call function
	err := RpcHttpClientCall(ip, port, "GoApi.MD5Encode", request, &response)
	if err != nil {
		t.Fatal("Error rpc http call function:", err)
	}
	fmt.Println("Rpc request:", request)
	fmt.Println("Rpc response:", response)
}

func TestCallRpcApiMD5Equal(t *testing.T) {
	ip := "127.0.0.1"
	port := "12000"
	request := TNetsRpcPackMD5EqualReq{Src: "Satellite", Dest: "c2b5e73361a4bf9d26a73413d0abee5e"}
	response := TNetsRpcPackMD5EqualResp{}
	// start rpc http server...
	go StartRpcHttpServer(ip, port)
	// start rpc http client call function
	err := RpcHttpClientCall(ip, port, "GoApi.MD5Equal", request, &response)
	if err != nil {
		t.Fatal("Error rpc http call function:", err)
	}
	fmt.Println("Rpc request:", request)
	fmt.Println("Rpc response:", response)
}
