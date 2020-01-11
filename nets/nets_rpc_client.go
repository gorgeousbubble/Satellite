package nets

import (
	"errors"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"reflect"
)

func RpcHttpClientCall(ip string, port string, method string, request interface{}, response interface{}) (err error) {
	// check request format...
	var rType = reflect.TypeOf(request)
	// check the request type kind
	if rType.Kind() != reflect.Struct {
		err = errors.New("request interface should be struct")
		log.Println("Error request format:", err)
		return err
	}
	// check response format...
	rType = reflect.TypeOf(response)
	// check the response type kind
	if rType.Kind() != reflect.Ptr {
		err = errors.New("response interface should be struct pointer")
		log.Println("Error response format:", err)
		return err
	}
	// dial rpc http server...
	conn, err := rpc.DialHTTP("tcp", ip+":"+port)
	if err != nil {
		log.Println("Error dial rpc http server:", err)
		return err
	}
	defer conn.Close()
	// call rpc remote api...
	err = conn.Call(method, request, response)
	if err != nil {
		log.Println("Error call rpc remote api:", err)
		return err
	}
	return err
}

func RpcTcpClientCall(ip string, port string, method string, request interface{}, response interface{}) (err error) {
	// check request format...
	var rType = reflect.TypeOf(request)
	// check the request type kind
	if rType.Kind() != reflect.Struct {
		err = errors.New("request interface should be struct")
		log.Println("Error request format:", err)
		return err
	}
	// check response format...
	rType = reflect.TypeOf(response)
	// check the response type kind
	if rType.Kind() != reflect.Ptr {
		err = errors.New("response interface should be struct pointer")
		log.Println("Error response format:", err)
		return err
	}
	// dial rpc tcp server...
	conn, err := jsonrpc.Dial("tcp", ip + ":" + port)
	if err != nil {
		log.Println("Error dial rpc tcp server:", err)
		return err
	}
	defer conn.Close()
	// call rpc remote api...
	err = conn.Call(method, request, response)
	if err != nil {
		log.Println("Error call rpc remote api:", err)
		return err
	}
	return err
}
