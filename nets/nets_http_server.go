package nets

import (
	"net/http"
	. "satellite/global"
	"time"
)

func StartHttpServer(ip string, port string) {
	server := http.Server{
		Addr:         ip + ":" + port,
		WriteTimeout: ConstHTTPWriteTimeout * time.Millisecond,
		ReadTimeout:  ConstHTTPReadTimeout * time.Millisecond,
		Handler:      createHttpRouter(),
	}
	server.ListenAndServe()
}
