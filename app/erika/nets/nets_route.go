package nets

import "github.com/gorilla/mux"

func createHttpRouter() (r *mux.Router) {
	r = mux.NewRouter()
	r.HandleFunc(HTTPURLHello, handleHello).Methods("GET")
	return r
}
