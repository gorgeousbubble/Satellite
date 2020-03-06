package nets

import (
	"fmt"
	"net/http"
	"satellite/app/erika/logs"
	"time"
)

func StartHttpsServer(ip string, port string) (err error) {
	server := http.Server{
		Addr:         ip + ":" + port,
		WriteTimeout: HTTPWriteTimeout * time.Millisecond,
		ReadTimeout:  HTTPReadTimeout * time.Millisecond,
		Handler:      createHttpRouter(),
	}
	fmt.Println("Start Listen And Server on ", ip+":"+port)
	err = GenerateCA(ip)
	if err != nil {
		logs.Error("Error Create CA:", err)
		return err
	}
	err = server.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil {
		logs.Error("Error Listen And Server:", err)
		return err
	}
	return err
}
