package nets

import (
	"fmt"
	"log"
	"net/http"
	. "satellite/global"
	"time"
)

func StartHttpsServer(ip string, port string) {
	server := http.Server{
		Addr:         ip + ":" + port,
		WriteTimeout: HTTPWriteTimeout * time.Millisecond,
		ReadTimeout:  HTTPReadTimeout * time.Millisecond,
		Handler:      createHttpRouter(),
	}
	fmt.Println("Start Listen And Server on ", ip+":"+port)
	err := GenerateCA(ip)
	if err != nil {
		log.Println("Error Create CA:", err)
	}
	err = server.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil {
		log.Println("Error Listen And Server:", err)
	}
}
