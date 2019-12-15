package nets

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"satellite/comp"
	"satellite/decomp"
	"satellite/images"
	"satellite/pack"
	"satellite/parses"
	"satellite/unpack"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/skip2/go-qrcode"
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

func handleIndex(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		log.Printf("GET %s", r.RequestURI)
		err = handleGetIndex(w, r)
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

func handleNetsPackProcess(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		log.Printf("GET %s", r.RequestURI)
		err = handleGetNetsPackProcess(w, r)
	case "POST":
		log.Printf("POST %s", r.RequestURI)
		err = handlePostNetsPackProcess(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("%d Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleNetsUnpackVerbose(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		log.Printf("GET %s", r.RequestURI)
		err = handleGetNetsUnpackVerbose(w, r)
	case "POST":
		log.Printf("POST %s", r.RequestURI)
		err = handlePostNetsUnpackVerbose(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("%d Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleNetsUnpackProcess(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		log.Printf("GET %s", r.RequestURI)
		err = handleGetNetsUnpackProcess(w, r)
	case "POST":
		log.Printf("POST %s", r.RequestURI)
		err = handlePostNetsUnpackProcess(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("%d Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleNetsUnpackConfine(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "POST":
		log.Printf("POST %s", r.RequestURI)
		err = handlePostNetsUnpackConfine(w, r)
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
	case "POST":
		log.Printf("POST %s", r.RequestURI)
		err = handlePostNetsUnpackToFile(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("%d Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleNetsUnpackToFileConfine(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "POST":
		log.Printf("POST %s", r.RequestURI)
		err = handlePostNetsUnpackToFileConfine(w, r)
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
	case "POST":
		log.Printf("POST %s", r.RequestURI)
		err = handlePostNetsUnpackToMemory(w, r)
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

func handleNetsImagesQRCodeToFile(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "POST":
		log.Printf("POST %s", r.RequestURI)
		err = handlePostNetsImagesQRCodeToFile(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("%d Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleNetsImagesQRCodeToMemory(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "POST":
		log.Printf("POST %s", r.RequestURI)
		err = handlePostNetsImagesQRCodeToMemory(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("%d Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleNetsParsesIniValue(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		log.Printf("GET %s", r.RequestURI)
		err = handleGetNetsParsesIniValue(w, r)
	case "PUT":
		log.Printf("PUT %s", r.RequestURI)
		err = handlePutNetsParsesIniValue(w, r)
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

func handleGetIndex(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Satellite Project"))
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

func handleGetNetsPackProcess(w http.ResponseWriter, r *http.Request) (err error) {
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
	var t TNetsPackProcessReq
	err = json.Unmarshal(body, &t)
	if err != nil {
		http.Error(w, "Incorrect request body!", http.StatusBadRequest)
		log.Println("Error unmarshal json body:", err)
		log.Printf("%d Bad Request", http.StatusBadRequest)
		return nil
	}
	// check request parameters
	b, err := checkNetsPackProcessParameters(t)
	if err != nil {
		log.Println("Error check pack process parameters:", err)
		return err
	}
	if !b {
		http.Error(w, "Illegal parameters!", http.StatusUnprocessableEntity)
		log.Println("Illegal parameters")
		log.Printf("%d Unprocessable Entity", http.StatusUnprocessableEntity)
		return nil
	}
	// pack file process information
	var resp TNetsPackProcessResp
	ch := make(chan bool)
	count := 0
	finish := false
	go func(resp *TNetsPackProcessResp) {
		var work int64
		// work value
		err = pack.WorkCalculate(t.Src, t.Type, &work)
		if err != nil || work <= 0 {
			log.Println("Error calculate pack work")
			ch <- false
			return
		}
		// done value
		done := atomic.LoadInt64(&pack.Done)
		// assignment
		(*resp).Done = done
		(*resp).Work = work
		ch <- true
	}(&resp)
	for {
		select {
		case r := <-ch:
			if r == false {
				log.Println("Pack process failure:", err)
				return err
			}
			log.Println("Pack process success.")
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

func handlePostNetsPackProcess(w http.ResponseWriter, r *http.Request) (err error) {
	return handleGetNetsPackProcess(w, r)
}

func handleGetNetsUnpackVerbose(w http.ResponseWriter, r *http.Request) (err error) {
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

func handlePostNetsUnpackVerbose(w http.ResponseWriter, r *http.Request) (err error) {
	return handleGetNetsUnpackVerbose(w, r)
}

func handleGetNetsUnpackProcess(w http.ResponseWriter, r *http.Request) (err error) {
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
	var t TNetsUnpackProcessReq
	err = json.Unmarshal(body, &t)
	if err != nil {
		http.Error(w, "Incorrect request body!", http.StatusBadRequest)
		log.Println("Error unmarshal json body:", err)
		log.Printf("%d Bad Request", http.StatusBadRequest)
		return nil
	}
	// check request parameters
	b, err := checkNetsUnpackProcessParameters(t)
	if err != nil {
		log.Println("Error check unpack process parameters:", err)
		return err
	}
	if !b {
		http.Error(w, "Illegal parameters!", http.StatusUnprocessableEntity)
		log.Println("Illegal parameters")
		log.Printf("%d Unprocessable Entity", http.StatusUnprocessableEntity)
		return nil
	}
	// unpack file process information
	var resp TNetsUnpackProcessResp
	ch := make(chan bool)
	count := 0
	finish := false
	go func(resp *TNetsUnpackProcessResp) {
		var algorithm string
		var work int64
		// work value
		err = unpack.WorkCalculate(t.Src, &algorithm, &work)
		if err != nil || work <= 0 {
			log.Println("Error calculate unpack work")
			ch <- false
			return
		}
		// done value
		done := atomic.LoadInt64(&unpack.Done)
		// assignment
		(*resp).Done = done
		(*resp).Work = work
		ch <- true
	}(&resp)
	for {
		select {
		case r := <-ch:
			if r == false {
				log.Println("Unpack process failure:", err)
				return err
			}
			log.Println("Unpack process success.")
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

func handlePostNetsUnpackProcess(w http.ResponseWriter, r *http.Request) (err error) {
	return handleGetNetsUnpackProcess(w, r)
}

func handlePostNetsUnpackConfine(w http.ResponseWriter, r *http.Request) (err error) {
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
		err = unpack.UnpackConfine(t.Src, t.Dest)
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
				log.Println("Unpack confine failure:", err)
				return err
			}
			log.Println("Unpack confine success.")
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

func handlePostNetsUnpackToFile(w http.ResponseWriter, r *http.Request) (err error) {
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

func handlePostNetsUnpackToFileConfine(w http.ResponseWriter, r *http.Request) (err error) {
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
		err = unpack.UnpackToFileConfine(t.Src, t.Target, t.Dest)
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

func handlePostNetsUnpackToMemory(w http.ResponseWriter, r *http.Request) (err error) {
	return handleGetNetsUnpackToMemory(w, r)
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

func handlePostNetsImagesQRCodeToFile(w http.ResponseWriter, r *http.Request) (err error) {
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
	var t TNetsImagesQRCodeToFile
	err = json.Unmarshal(body, &t)
	if err != nil {
		http.Error(w, "Incorrect request body!", http.StatusBadRequest)
		log.Println("Error unmarshal json body:", err)
		log.Printf("%d Bad Request", http.StatusBadRequest)
		return nil
	}
	// check request parameters
	b, err := checkNetsImagesQRCodeToFileParameters(t)
	if err != nil {
		log.Println("Error check images qrcode parameters:", err)
		return err
	}
	if !b {
		http.Error(w, "Illegal parameters!", http.StatusUnprocessableEntity)
		log.Println("Illegal parameters")
		log.Printf("%d Unprocessable Entity", http.StatusUnprocessableEntity)
		return nil
	}
	// generate qrcode
	err = images.QRCodeGenerateToFile(t.Content, qrcode.Highest, t.Size, t.Dest)
	if err != nil {
		log.Println("Images QRCode failure:", err)
		return err
	}
	log.Println("Images QRCode success.")
	// response
	w.Header().Set("Content-Type", "text/plain")
	log.Printf("%d Ok", http.StatusOK)
	return err
}

func handlePostNetsImagesQRCodeToMemory(w http.ResponseWriter, r *http.Request) (err error) {
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
	var t TNetsImagesQRCodeToMemory
	err = json.Unmarshal(body, &t)
	if err != nil {
		http.Error(w, "Incorrect request body!", http.StatusBadRequest)
		log.Println("Error unmarshal json body:", err)
		log.Printf("%d Bad Request", http.StatusBadRequest)
		return nil
	}
	// check request parameters
	b, err := checkNetsImagesQRCodeParameters(t)
	if err != nil {
		log.Println("Error check images qrcode parameters:", err)
		return err
	}
	if !b {
		http.Error(w, "Illegal parameters!", http.StatusUnprocessableEntity)
		log.Println("Illegal parameters")
		log.Printf("%d Unprocessable Entity", http.StatusUnprocessableEntity)
		return nil
	}
	// generate qrcode
	qr, err := images.QRCodeGenerateToMemory(t.Content, qrcode.Highest, t.Size)
	if err != nil {
		log.Println("Images QRCode failure:", err)
		return err
	}
	log.Println("Images QRCode success.")
	// response
	w.Header().Set("Content-Type", "application/json")
	w.Write(qr)
	log.Printf("%d Ok", http.StatusOK)
	return err
}

func handleGetNetsParsesIniValue(w http.ResponseWriter, r *http.Request) (err error) {
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
	var t TNetsParsesIni
	err = json.Unmarshal(body, &t)
	if err != nil {
		http.Error(w, "Incorrect request body!", http.StatusBadRequest)
		log.Println("Error unmarshal json body:", err)
		log.Printf("%d Bad Request", http.StatusBadRequest)
		return nil
	}
	// check request parameters
	b, err := checkNetsParsesIniValueParameters(t)
	if err != nil {
		log.Println("Error check parses ini value parameters:", err)
		return err
	}
	if !b {
		http.Error(w, "Illegal parameters!", http.StatusUnprocessableEntity)
		log.Println("Illegal parameters")
		log.Printf("%d Unprocessable Entity", http.StatusUnprocessableEntity)
		return nil
	}
	// parses ini get value
	switch t.Mode {
	case "get":
		switch t.Type {
		case "string":
			var value string
			err = parses.GetValueFrom(t.Src, t.Section, t.Name, &value)
			if err != nil {
				return err
			}
			t.Value = value
		case "int":
			var value int
			err = parses.GetValueFrom(t.Src, t.Section, t.Name, &value)
			if err != nil {
				return err
			}
			t.Value = strconv.Itoa(value)
		case "float64":
			var value float64
			err = parses.GetValueFrom(t.Src, t.Section, t.Name, &value)
			if err != nil {
				return err
			}
			t.Value = strconv.FormatFloat(value, 'f', -1, 64)
		case "bool":
			var value bool
			err = parses.GetValueFrom(t.Src, t.Section, t.Name, &value)
			if err != nil {
				return err
			}
			t.Value = strconv.FormatBool(value)
		default:
			err = errors.New("unrecognized type name")
			return err
		}
	default:
		err = errors.New("unrecognized type name")
		return err
	}
	// marshal json
	js, err := json.MarshalIndent(&t, "", "\t\t")
	if err != nil {
		log.Println("Error marshal to json:", err)
		return err
	}
	// response
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	log.Printf("%d Ok", http.StatusOK)
	return err
}

func handlePutNetsParsesIniValue(w http.ResponseWriter, r *http.Request) (err error) {
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
	var t TNetsParsesIni
	err = json.Unmarshal(body, &t)
	if err != nil {
		http.Error(w, "Incorrect request body!", http.StatusBadRequest)
		log.Println("Error unmarshal json body:", err)
		log.Printf("%d Bad Request", http.StatusBadRequest)
		return nil
	}
	// check request parameters
	b, err := checkNetsParsesIniValueParameters(t)
	if err != nil {
		log.Println("Error check parses ini value parameters:", err)
		return err
	}
	if !b {
		http.Error(w, "Illegal parameters!", http.StatusUnprocessableEntity)
		log.Println("Illegal parameters")
		log.Printf("%d Unprocessable Entity", http.StatusUnprocessableEntity)
		return nil
	}
	// parses ini get value
	switch t.Mode {
	case "set":
		switch t.Type {
		case "string":
			var value string
			value = t.Value
			err = parses.SetValueTo(t.Src, t.Section, t.Name, value)
			if err != nil {
				return err
			}
		case "int":
			var value int
			value, err = strconv.Atoi(t.Value)
			if err != nil {
				return err
			}
			err = parses.SetValueTo(t.Src, t.Section, t.Name, value)
			if err != nil {
				return err
			}
		case "float64":
			var value float64
			value, err = strconv.ParseFloat(t.Value, 64)
			if err != nil {
				return err
			}
			err = parses.SetValueTo(t.Src, t.Section, t.Name, value)
			if err != nil {
				return err
			}
		case "bool":
			var value bool
			value, err = strconv.ParseBool(t.Value)
			if err != nil {
				return err
			}
			err = parses.SetValueTo(t.Src, t.Section, t.Name, value)
			if err != nil {
				return err
			}
		default:
			err = errors.New("unrecognized type name")
		}
	default:
		err = errors.New("unrecognized type name")
		return err
	}
	// response
	w.Header().Set("Content-Type", "application/json")
	log.Printf("%d Ok", http.StatusOK)
	return err
}
