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
		WriteTimeout: ConstHTTPWriteTimeout * time.Millisecond,
		ReadTimeout:  ConstHTTPReadTimeout * time.Millisecond,
		Handler:      createHttpRouter(),
	}
	fmt.Println("Start Listen And Server on ", ip+":"+port)
	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil {
		log.Println("Error Listen And Server:", err)
	}
}
