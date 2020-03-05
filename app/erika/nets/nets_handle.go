package nets

import (
	"net/http"
	"satellite/app/erika/logs"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		logs.Event("GET ", r.RequestURI)
		err = handleGetHello(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logs.Error("Internal Server Error ", http.StatusInternalServerError)
		return
	}
}

func handleGetHello(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("I'm Erika, nice to meet you~"))
	return err
}
