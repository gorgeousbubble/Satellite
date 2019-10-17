package nets

import (
	"Satellite/comp"
	"Satellite/decomp"
	"encoding/json"
	"errors"
	"fmt"
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
	case "GET":
		log.Printf("GET %s", r.RequestURI)
		err = handleGetNetsUnpack(w, r)
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

func handleNetsUnpackToFile(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		log.Printf("GET %s", r.RequestURI)
		err = handleGetNetsUnpackToFile(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("%d Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleNetsUnpackToMemory(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		log.Printf("GET %s", r.RequestURI)
		err = handleGetNetsUnpackToMemory(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("%d Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleNetsComp(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "POST":
		log.Printf("POST %s", r.RequestURI)
		err = handlePostNetsComp(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("%d Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleNetsDecomp(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "POST":
		log.Printf("POST %s", r.RequestURI)
		err = handlePostNetsDecomp(w, r)
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
	// start unpack files
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

func handleGetNetsUnpack(w http.ResponseWriter, r *http.Request) (err error) {
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
	var t TNetsUnpackVerboseReq
	err = json.Unmarshal(body, &t)
	if err != nil {
		http.Error(w, "Incorrect request body!", http.StatusBadRequest)
		log.Println("Error unmarshal json body:", err)
		log.Printf("%d Bad Request", http.StatusBadRequest)
		return nil
	}
	// check request parameters
	b, err := checkNetsUnpackVerboseParameters(t)
	if err != nil {
		log.Println("Error check unpack verbose parameters:", err)
		return err
	}
	if !b {
		http.Error(w, "Illegal parameters!", http.StatusUnprocessableEntity)
		log.Println("Illegal parameters")
		log.Printf("%d Unprocessable Entity", http.StatusUnprocessableEntity)
		return nil
	}
	// unpack file verbose information
	var resp TNetsUnpackVerboseResp
	ch := make(chan bool)
	count := 0
	finish := false
	go func(resp *TNetsUnpackVerboseResp) {
		var files []string
		var sizes []int
		var algorithm string
		err = unpack.ExtractInfo(t.Src, &files, &sizes, &algorithm)
		if err != nil {
			log.Println("Error extract unpack information")
			ch <- false
			return
		}
		for k, v := range files {
			var t TNetsUnpackFileInfo
			t.Name = v
			t.Size = fmt.Sprintf("%dkb", sizes[k]/1024)
			t.Type = algorithm
			(*resp).Files = append((*resp).Files, t)
		}
		ch <- true
		return
	}(&resp)
	for {
		select {
		case r := <-ch:
			if r == false {
				log.Println("Unpack verbose failure:", err)
				return err
			}
			log.Println("Unpack verbose success.")
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
	// marshal json
	js, err := json.MarshalIndent(&resp, "", "\t\t")
	if err != nil {
		log.Println("Error marshal to json:", err)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	log.Printf("%d Ok", http.StatusOK)
	return err
}

func handleGetNetsUnpackToFile(w http.ResponseWriter, r *http.Request) (err error) {
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
	var t TNetsUnpackToFile
	err = json.Unmarshal(body, &t)
	if err != nil {
		http.Error(w, "Incorrect request body!", http.StatusBadRequest)
		log.Println("Error unmarshal json body:", err)
		log.Printf("%d Bad Request", http.StatusBadRequest)
		return nil
	}
	// check request parameters
	b, err := checkNetsUnpackToFileParameters(t)
	if err != nil {
		log.Println("Error check unpack to file parameters:", err)
		return err
	}
	if !b {
		http.Error(w, "Illegal parameters!", http.StatusUnprocessableEntity)
		log.Println("Illegal parameters")
		log.Printf("%d Unprocessable Entity", http.StatusUnprocessableEntity)
		return nil
	}
	// unpack file to file
	ch := make(chan bool)
	count := 0
	finish := false
	go func() {
		err = unpack.UnpackToFile(t.Src, t.Target, t.Dest)
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
				log.Println("Unpack to file failure:", err)
				return err
			}
			log.Println("Unpack to file success.")
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
	w.Header().Set("Content-Type", "application/json")
	log.Printf("%d Ok", http.StatusOK)
	return err
}

func handleGetNetsUnpackToMemory(w http.ResponseWriter, r *http.Request) (err error) {
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
	var t TNetsUnpackToMemory
	err = json.Unmarshal(body, &t)
	if err != nil {
		http.Error(w, "Incorrect request body!", http.StatusBadRequest)
		log.Println("Error unmarshal json body:", err)
		log.Printf("%d Bad Request", http.StatusBadRequest)
		return nil
	}
	// check request parameters
	b, err := checkNetsUnpackToMemoryParameters(t)
	if err != nil {
		log.Println("Error check unpack to memory parameters:", err)
		return err
	}
	if !b {
		http.Error(w, "Illegal parameters!", http.StatusUnprocessableEntity)
		log.Println("Illegal parameters")
		log.Printf("%d Unprocessable Entity", http.StatusUnprocessableEntity)
		return nil
	}
	// unpack file to memory
	var dest []byte
	ch := make(chan bool)
	count := 0
	finish := false
	go func(resp *[]byte) {
		err = unpack.UnpackToMemory(t.Src, t.Target, resp)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
		return
	}(&dest)
	for {
		select {
		case r := <-ch:
			if r == false {
				log.Println("Unpack to memory failure:", err)
				return err
			}
			log.Println("Unpack to memory success.")
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
	w.Header().Set("Content-Type", "application/json")
	w.Write(dest)
	log.Printf("%d Ok", http.StatusOK)
	return err
}

func handlePostNetsComp(w http.ResponseWriter, r *http.Request) (err error) {
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
	var t TNetsComp
	err = json.Unmarshal(body, &t)
	if err != nil {
		http.Error(w, "Incorrect request body!", http.StatusBadRequest)
		log.Println("Error unmarshal json body:", err)
		log.Printf("%d Bad Request", http.StatusBadRequest)
		return nil
	}
	// check request parameters
	b, err := checkNetsCompParameters(t)
	if err != nil {
		log.Println("Error check comp parameters:", err)
		return err
	}
	if !b {
		http.Error(w, "Illegal parameters!", http.StatusUnprocessableEntity)
		log.Println("Illegal parameters")
		log.Printf("%d Unprocessable Entity", http.StatusUnprocessableEntity)
		return nil
	}
	// refactor source files
	t.Src, err = refactorNetsCompSource(t.Src)
	if err != nil {
		log.Println("Error refactor source files:", err)
		return err
	}
	// start compress files
	ch := make(chan bool)
	count := 0
	finish := false
	go func() {
		err = comp.Compress(t.Src, t.Dest, t.Type)
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
				log.Println("Compress failure:", err)
				return err
			}
			log.Println("Compress success.")
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

func handlePostNetsDecomp(w http.ResponseWriter, r *http.Request) (err error) {
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
	var t TNetsDecomp
	err = json.Unmarshal(body, &t)
	if err != nil {
		http.Error(w, "Incorrect request body!", http.StatusBadRequest)
		log.Println("Error unmarshal json body:", err)
		log.Printf("%d Bad Request", http.StatusBadRequest)
		return nil
	}
	// check request parameters
	b, err := checkNetsDecompParameters(t)
	if err != nil {
		log.Println("Error check decomp parameters:", err)
		return err
	}
	if !b {
		http.Error(w, "Illegal parameters!", http.StatusUnprocessableEntity)
		log.Println("Illegal parameters")
		log.Printf("%d Unprocessable Entity", http.StatusUnprocessableEntity)
		return nil
	}
	// start decompress files
	ch := make(chan bool)
	count := 0
	finish := false
	go func() {
		err = decomp.DeCompress(t.Src, t.Dest, t.Type)
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
				log.Println("Decompress failure:", err)
				return err
			}
			log.Println("Decompress success.")
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
