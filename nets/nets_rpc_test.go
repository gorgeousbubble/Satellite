package nets

import (
	"fmt"
	"testing"
)

func TestHttpCallRpcApiMD5Encode(t *testing.T) {
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

func TestHttpCallRpcApiMD5Equal(t *testing.T) {
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

func TestTcpCallRpcApiMD5Encode(t *testing.T) {
	ip := "127.0.0.1"
	port := "12000"
	request := TNetsRpcPackMD5EncodeReq{Src: "Satellite"}
	response := TNetsRpcPackMD5EncodeResp{}
	t.Skip("Skip test rpc service in tcp connection mode...")
	// start rpc tcp server...
	go StartRpcTcpServer(ip, port)
	// start rpc tcp client call function
	err := RpcTcpClientCall(ip, port, "GoApi.MD5Encode", request, &response)
	if err != nil {
		t.Fatal("Error rpc tcp call function:", err)
	}
	fmt.Println("Rpc request:", request)
	fmt.Println("Rpc response:", response)
}

func TestTcpCallRpcApiMD5Equal(t *testing.T) {
	ip := "127.0.0.1"
	port := "12000"
	request := TNetsRpcPackMD5EqualReq{Src: "Satellite", Dest: "c2b5e73361a4bf9d26a73413d0abee5e"}
	response := TNetsRpcPackMD5EqualResp{}
	t.Skip("Skip test rpc service in tcp connection mode...")
	// start rpc tcp server...
	go StartRpcTcpServer(ip, port)
	// start rpc tcp client call function
	err := RpcTcpClientCall(ip, port, "GoApi.MD5Equal", request, &response)
	if err != nil {
		t.Fatal("Error rpc tcp call function:", err)
	}
	fmt.Println("Rpc request:", request)
	fmt.Println("Rpc response:", response)
}
