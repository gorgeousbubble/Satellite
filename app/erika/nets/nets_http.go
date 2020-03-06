package nets

import (
	"fmt"
	"net/http"
	"satellite/app/erika/logs"
	"time"
)

func StartHttpServer(ip string, port string) (err error) {
	server := http.Server{
		Addr:         ip + ":" + port,
		WriteTimeout: HTTPWriteTimeout * time.Millisecond,
		ReadTimeout:  HTTPReadTimeout * time.Millisecond,
		Handler:      createHttpRouter(),
	}
	fmt.Println("Start Listen And Server on ", ip+":"+port)
	err = server.ListenAndServe()
	if err != nil {
		logs.Error("Error Listen And Server:", err)
		return err
	}
	return err
}
