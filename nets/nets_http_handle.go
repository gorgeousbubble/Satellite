package nets

import (
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGetRoot(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGetRoot(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello,World!"))
	return err
}
