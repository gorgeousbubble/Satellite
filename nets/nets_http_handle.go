package nets

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"satellite/pack"
	"satellite/unpack"
	"time"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		log.Printf("GET %s", r.RequestURI)
		err = handleGetRoot(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("%d Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleNetsPack(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "POST":
		log.Printf("POST %s", r.RequestURI)
		err = handlePostNetsPack(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("%d Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleNetsUnpack(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "POST":
		log.Printf("POST %s", r.RequestURI)
		err = handlePostNetsUnpack(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("%d Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleGetRoot(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello,World!"))
	return err
}

func handlePostNetsPack(w http.ResponseWriter, r *http.Request) (err error) {
	defer r.Body.Close()
	// read request body
	len := r.ContentLength
	body := make([]byte, len)
	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error read request body:", err)
		return err
	}
	// unmarshal json body
	var t TNetsPack
	err = json.Unmarshal(body, &t)
	if err != nil {
		http.Error(w, "Incorrect request body!", http.StatusBadRequest)
		log.Println("Error unmarshal json body:", err)
		log.Printf("%d Bad Request", http.StatusBadRequest)
		return nil
	}
	// check request parameters
	b, err := checkNetsPackParameters(t)
	if err != nil {
		log.Println("Error check pack parameters:", err)
		return err
	}
	if !b {
		http.Error(w, "Illegal parameters!", http.StatusUnprocessableEntity)
		log.Println("Illegal parameters")
		log.Printf("%d Unprocessable Entity", http.StatusUnprocessableEntity)
		return nil
	}
	// refactor source files
	t.Src, err = refactorNetsPackSource(t.Src)
	if err != nil {
		log.Println("Error refactor source files:", err)
		return err
	}
	// start pack files
	ch := make(chan bool)
	count := 0
	finish := false
	go func() {
		err = pack.Pack(t.Src, t.Dest, t.Type)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
		return
	}()
	for {
		select {
		case r := <-ch:
			if r == false {
				log.Println("Pack failure:", err)
				return err
			}
			log.Println("Pack success.")
			finish = true
			break
		default:
			count++
			time.Sleep(100 * time.Millisecond)
		}
		if finish {
			break
		}
		if count >= 100 {
			err = errors.New("timeout")
			return err
		}
	}
	w.Header().Set("Content-Type", "text/plain")
	log.Printf("%d Ok", http.StatusOK)
	return err
}

func handlePostNetsUnpack(w http.ResponseWriter, r *http.Request) (err error) {
	defer r.Body.Close()
	// read request body
	len := r.ContentLength
	body := make([]byte, len)
	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error read request body:", err)
		return err
	}
	// unmarshal json body
	var t TNetsUnpack
	err = json.Unmarshal(body, &t)
	if err != nil {
		http.Error(w, "Incorrect request body!", http.StatusBadRequest)
		log.Println("Error unmarshal json body:", err)
		log.Printf("%d Bad Request", http.StatusBadRequest)
		return nil
	}
	// check request parameters
	b, err := checkNetsUnpackParameters(t)
	if err != nil {
		log.Println("Error check unpack parameters:", err)
		return err
	}
	if !b {
		http.Error(w, "Illegal parameters!", http.StatusUnprocessableEntity)
		log.Println("Illegal parameters")
		log.Printf("%d Unprocessable Entity", http.StatusUnprocessableEntity)
		return nil
	}
	// start pack files
	ch := make(chan bool)
	count := 0
	finish := false
	go func() {
		err = unpack.Unpack(t.Src, t.Dest)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
		return
	}()
	for {
		select {
		case r := <-ch:
			if r == false {
				log.Println("Unpack failure:", err)
				return err
			}
			log.Println("Unpack success.")
			finish = true
			break
		default:
			count++
			time.Sleep(100 * time.Millisecond)
		}
		if finish {
			break
		}
		if count >= 100 {
			err = errors.New("timeout")
			return err
		}
	}
	w.Header().Set("Content-Type", "text/plain")
	log.Printf("%d Ok", http.StatusOK)
	return err
}
