package nets

import (
	"net/http"
	"net/http/httptest"
	. "satellite/global"
	"strings"
	"testing"
)

func TestHandlePostNetsPack(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLPack, handleNetsPack)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"src": ["../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"], "dest": "../test/data/pack/file_aes.txt", "type": "aes"}`)
	request, _ := http.NewRequest("POST", HttpURLPack, body)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func BenchmarkHandlePostNetsPack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mux := http.NewServeMux()
		mux.HandleFunc(HttpURLPack, handleNetsPack)

		writer := httptest.NewRecorder()
		body := strings.NewReader(`{"src": ["../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"], "dest": "../test/data/pack/file_aes.txt", "type": "aes"}`)
		request, _ := http.NewRequest("POST", HttpURLPack, body)
		mux.ServeHTTP(writer, request)

		if writer.Code != http.StatusOK {
			b.Errorf("Response code is %v", writer.Code)
		}
	}
}

func TestHandlePostNetsUnpack(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLUnpack, handleNetsUnpack)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"src": "../test/data/unpack/file_aes.txt", "dest": "../test/data/unpack/"}`)
	request, _ := http.NewRequest("POST", HttpURLUnpack, body)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func BenchmarkHandlePostNetsUnpack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mux := http.NewServeMux()
		mux.HandleFunc(HttpURLUnpack, handleNetsUnpack)

		writer := httptest.NewRecorder()
		body := strings.NewReader(`{"src": "../test/data/unpack/file_aes.txt", "dest": "../test/data/unpack/"}`)
		request, _ := http.NewRequest("POST", HttpURLUnpack, body)
		mux.ServeHTTP(writer, request)

		if writer.Code != http.StatusOK {
			b.Errorf("Response code is %v", writer.Code)
		}
	}
}

func TestHandleGetNetsUnpackVerbose(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLUnpackVerbose, handleNetsUnpackVerbose)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"src": "../test/data/unpack/file_aes.txt"}`)
	request, _ := http.NewRequest("GET", HttpURLUnpackVerbose, body)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func BenchmarkHandleGetNetsUnpackVerbose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mux := http.NewServeMux()
		mux.HandleFunc(HttpURLUnpackVerbose, handleNetsUnpackVerbose)

		writer := httptest.NewRecorder()
		body := strings.NewReader(`{"src": "../test/data/unpack/file_aes.txt"}`)
		request, _ := http.NewRequest("GET", HttpURLUnpackVerbose, body)
		mux.ServeHTTP(writer, request)

		if writer.Code != http.StatusOK {
			b.Errorf("Response code is %v", writer.Code)
		}
	}
}

func TestHandleGetNetsUnpackProcess(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLUnpackProcess, handleNetsUnpackProcess)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"src": "../test/data/unpack/file_aes.txt"}`)
	request, _ := http.NewRequest("GET", HttpURLUnpackProcess, body)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func BenchmarkHandleGetNetsUnpackProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mux := http.NewServeMux()
		mux.HandleFunc(HttpURLUnpackProcess, handleNetsUnpackProcess)

		writer := httptest.NewRecorder()
		body := strings.NewReader(`{"src": "../test/data/unpack/file_aes.txt"}`)
		request, _ := http.NewRequest("GET", HttpURLUnpackProcess, body)
		mux.ServeHTTP(writer, request)

		if writer.Code != http.StatusOK {
			b.Errorf("Response code is %v", writer.Code)
		}
	}
}

func TestHandleGetNetsUnpackToMemory(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLUnpackToMemory, handleNetsUnpackToMemory)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"src": "../test/data/unpack/file_aes.txt", "target": "file_1.txt"}`)
	request, _ := http.NewRequest("GET", HttpURLUnpackToMemory, body)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func BenchmarkHandleGetNetsUnpackToMemory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mux := http.NewServeMux()
		mux.HandleFunc(HttpURLUnpackToMemory, handleNetsUnpackToMemory)

		writer := httptest.NewRecorder()
		body := strings.NewReader(`{"src": "../test/data/unpack/file_aes.txt", "target": "file_1.txt"}`)
		request, _ := http.NewRequest("GET", HttpURLUnpackToMemory, body)
		mux.ServeHTTP(writer, request)

		if writer.Code != http.StatusOK {
			b.Errorf("Response code is %v", writer.Code)
		}
	}
}

func TestHandlePostNetsUnpackToFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLUnpackToFile, handleNetsUnpackToFile)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"src": "../test/data/unpack/file_aes.txt", "target": "file_1.txt", "dest": "../test/data/unpack/"}`)
	request, _ := http.NewRequest("POST", HttpURLUnpackToFile, body)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func BenchmarkHandlePostNetsUnpackToFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mux := http.NewServeMux()
		mux.HandleFunc(HttpURLUnpackToFile, handleNetsUnpackToFile)

		writer := httptest.NewRecorder()
		body := strings.NewReader(`{"src": "../test/data/unpack/file_aes.txt", "target": "file_1.txt", "dest": "../test/data/unpack/"}`)
		request, _ := http.NewRequest("POST", HttpURLUnpackToFile, body)
		mux.ServeHTTP(writer, request)

		if writer.Code != http.StatusOK {
			b.Errorf("Response code is %v", writer.Code)
		}
	}
}

func TestHandlePostNetsComp(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLComp, handleNetsComp)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"src": ["../test/data/comp/file_1.txt", "../test/data/comp/file_2.txt", "../test/data/comp/file_3.txt", "../test/data/comp/file_4.txt", "../test/data/pack/file_5.txt"], "dest": "../test/data/comp/file.tar.gz", "type": "tar"}`)
	request, _ := http.NewRequest("POST", HttpURLComp, body)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func BenchmarkHandlePostNetsComp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mux := http.NewServeMux()
		mux.HandleFunc(HttpURLComp, handleNetsComp)

		writer := httptest.NewRecorder()
		body := strings.NewReader(`{"src": ["../test/data/comp/file_1.txt", "../test/data/comp/file_2.txt", "../test/data/comp/file_3.txt", "../test/data/comp/file_4.txt", "../test/data/pack/file_5.txt"], "dest": "../test/data/comp/file.tar.gz", "type": "tar"}`)
		request, _ := http.NewRequest("POST", HttpURLComp, body)
		mux.ServeHTTP(writer, request)

		if writer.Code != http.StatusOK {
			b.Errorf("Response code is %v", writer.Code)
		}
	}
}

func TestHandlePostNetsDecomp(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLDecomp, handleNetsDecomp)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"src": "../test/data/decomp/file.tar.gz", "dest": "../test/data/decomp/", "type": "tar"}`)
	request, _ := http.NewRequest("POST", HttpURLDecomp, body)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func BenchmarkHandlePostNetsDecomp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mux := http.NewServeMux()
		mux.HandleFunc(HttpURLDecomp, handleNetsDecomp)

		writer := httptest.NewRecorder()
		body := strings.NewReader(`{"src": "../test/data/decomp/file.tar.gz", "dest": "../test/data/decomp/", "type": "tar"}`)
		request, _ := http.NewRequest("POST", HttpURLDecomp, body)
		mux.ServeHTTP(writer, request)

		if writer.Code != http.StatusOK {
			b.Errorf("Response code is %v", writer.Code)
		}
	}
}
