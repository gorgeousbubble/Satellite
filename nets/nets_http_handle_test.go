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
