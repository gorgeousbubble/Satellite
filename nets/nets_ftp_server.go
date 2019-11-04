package nets

import (
	"fmt"
	"log"
	"net/http"
)

func StartFtpServer(ip string, port string) {
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Println("Start FTP Listen And Server on ", ip+":"+port)
	err := http.ListenAndServe(ip+":"+port, nil)
	if err != nil {
		log.Println("Error Listen And Server:", err)
	}
}
