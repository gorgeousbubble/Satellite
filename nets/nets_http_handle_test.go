package nets

import (
	"net/http"
	"net/http/httptest"
	. "satellite/global"
	"strings"
	"testing"
)

func TestHandleGetRoot(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLRoot, handleRoot)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", HttpURLRoot, nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func BenchmarkHandleGetRoot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mux := http.NewServeMux()
		mux.HandleFunc(HttpURLRoot, handleRoot)

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", HttpURLRoot, nil)
		mux.ServeHTTP(writer, request)

		if writer.Code != http.StatusOK {
			b.Errorf("Response code is %v", writer.Code)
		}
	}
}

func TestHandleGetIndex(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLSatellite, handleIndex)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", HttpURLSatellite, nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func BenchmarkHandleGetIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mux := http.NewServeMux()
		mux.HandleFunc(HttpURLSatellite, handleIndex)

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", HttpURLSatellite, nil)
		mux.ServeHTTP(writer, request)

		if writer.Code != http.StatusOK {
			b.Errorf("Response code is %v", writer.Code)
		}
	}
}

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

func TestHandleGetNetsPackProcess(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLPackProcess, handleNetsPackProcess)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"src": ["../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"], "type": "aes"}`)
	request, _ := http.NewRequest("GET", HttpURLPackProcess, body)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func BenchmarkHandleGetNetsPackProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mux := http.NewServeMux()
		mux.HandleFunc(HttpURLPackProcess, handleNetsPackProcess)

		writer := httptest.NewRecorder()
		body := strings.NewReader(`{"src": ["../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"], "type": "aes"}`)
		request, _ := http.NewRequest("GET", HttpURLPackProcess, body)
		mux.ServeHTTP(writer, request)

		if writer.Code != http.StatusOK {
			b.Errorf("Response code is %v", writer.Code)
		}
	}
}

func TestHandlePostNetsPackProcess(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLPackProcess, handleNetsPackProcess)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"src": ["../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"], "type": "aes"}`)
	request, _ := http.NewRequest("POST", HttpURLPackProcess, body)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func BenchmarkHandlePostNetsPackProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mux := http.NewServeMux()
		mux.HandleFunc(HttpURLPackProcess, handleNetsPackProcess)

		writer := httptest.NewRecorder()
		body := strings.NewReader(`{"src": ["../test/data/pack/file_1.txt", "../test/data/pack/file_2.txt", "../test/data/pack/file_3.txt", "../test/data/pack/file_4.txt", "../test/data/pack/file_5.txt"], "type": "aes"}`)
		request, _ := http.NewRequest("POST", HttpURLPackProcess, body)
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

func TestHandlePostNetsUnpackVerbose(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLUnpackVerbose, handleNetsUnpackVerbose)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"src": "../test/data/unpack/file_aes.txt"}`)
	request, _ := http.NewRequest("POST", HttpURLUnpackVerbose, body)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func BenchmarkHandlePostNetsUnpackVerbose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mux := http.NewServeMux()
		mux.HandleFunc(HttpURLUnpackVerbose, handleNetsUnpackVerbose)

		writer := httptest.NewRecorder()
		body := strings.NewReader(`{"src": "../test/data/unpack/file_aes.txt"}`)
		request, _ := http.NewRequest("POST", HttpURLUnpackVerbose, body)
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

func TestHandlePostNetsUnpackConfine(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLUnpackConfine, handleNetsUnpackConfine)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"src": "../test/data/unpack/file_aes.txt", "dest": "../test/data/unpack/"}`)
	request, _ := http.NewRequest("POST", HttpURLUnpackConfine, body)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func BenchmarkHandlePostNetsUnpackConfine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mux := http.NewServeMux()
		mux.HandleFunc(HttpURLUnpackConfine, handleNetsUnpackConfine)

		writer := httptest.NewRecorder()
		body := strings.NewReader(`{"src": "../test/data/unpack/file_aes.txt", "dest": "../test/data/unpack/"}`)
		request, _ := http.NewRequest("POST", HttpURLUnpackConfine, body)
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

func TestHandlePostNetsUnpackToMemory(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLUnpackToMemory, handleNetsUnpackToMemory)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"src": "../test/data/unpack/file_aes.txt", "target": "file_1.txt"}`)
	request, _ := http.NewRequest("POST", HttpURLUnpackToMemory, body)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func BenchmarkHandlePostNetsUnpackToMemory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mux := http.NewServeMux()
		mux.HandleFunc(HttpURLUnpackToMemory, handleNetsUnpackToMemory)

		writer := httptest.NewRecorder()
		body := strings.NewReader(`{"src": "../test/data/unpack/file_aes.txt", "target": "file_1.txt"}`)
		request, _ := http.NewRequest("POST", HttpURLUnpackToMemory, body)
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

func TestHandlePostNetsUnpackToFileConfine(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLUnpackToFileConfine, handleNetsUnpackToFileConfine)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"src": "../test/data/unpack/file_aes.txt", "target": "file_1.txt", "dest": "../test/data/unpack/"}`)
	request, _ := http.NewRequest("POST", HttpURLUnpackToFileConfine, body)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func BenchmarkHandlePostNetsUnpackToFileConfine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mux := http.NewServeMux()
		mux.HandleFunc(HttpURLUnpackToFileConfine, handleNetsUnpackToFileConfine)

		writer := httptest.NewRecorder()
		body := strings.NewReader(`{"src": "../test/data/unpack/file_aes.txt", "target": "file_1.txt", "dest": "../test/data/unpack/"}`)
		request, _ := http.NewRequest("POST", HttpURLUnpackToFileConfine, body)
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

func TestHandlePostNetsImagesQRCodeToFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLImagesQRCodeToFile, handleNetsImagesQRCodeToFile)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"content": "https://github.com/gorgeousbubble", "size": 256, "dest":"../test/data/images/qrcode/qr_test.png"}`)
	request, _ := http.NewRequest("POST", HttpURLImagesQRCodeToFile, body)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func BenchmarkHandlePostNetsImagesQRCodeToFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mux := http.NewServeMux()
		mux.HandleFunc(HttpURLImagesQRCodeToFile, handleNetsImagesQRCodeToFile)

		writer := httptest.NewRecorder()
		body := strings.NewReader(`{"content": "https://github.com/gorgeousbubble", "size": 256, "dest":"../test/data/images/qrcode/qr_test.png"}`)
		request, _ := http.NewRequest("POST", HttpURLImagesQRCodeToFile, body)
		mux.ServeHTTP(writer, request)

		if writer.Code != http.StatusOK {
			b.Errorf("Response code is %v", writer.Code)
		}
	}
}

func TestHandlePostNetsImagesQRCodeToMemory(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLImagesQRCodeToMemory, handleNetsImagesQRCodeToMemory)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"content": "https://localhost:8080/", "size": 256}`)
	request, _ := http.NewRequest("POST", HttpURLImagesQRCodeToMemory, body)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func BenchmarkHandlePostNetsImagesQRCodeToMemory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mux := http.NewServeMux()
		mux.HandleFunc(HttpURLImagesQRCodeToMemory, handleNetsImagesQRCodeToMemory)

		writer := httptest.NewRecorder()
		body := strings.NewReader(`{"content": "https://localhost:8080/", "size": 256}`)
		request, _ := http.NewRequest("POST", HttpURLImagesQRCodeToMemory, body)
		mux.ServeHTTP(writer, request)

		if writer.Code != http.StatusOK {
			b.Errorf("Response code is %v", writer.Code)
		}
	}
}

func TestHandleGetNetsParsesIniValue(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLParsesIni, handleNetsParsesIniValue)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"src": "../test/data/parses/test_simple.ini", "mode": "get", "section": "BOOL", "name": "Switch_On", "type": "bool", "value": ""}`)
	request, _ := http.NewRequest("GET", HttpURLParsesIni, body)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestHandlePutNetsParsesIniValue(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(HttpURLParsesIni, handleNetsParsesIniValue)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"src": "../test/data/parses/test_simple.ini", "mode": "set", "section": "BOOL", "name": "Switch_On", "type": "bool", "value": "true"}`)
	request, _ := http.NewRequest("PUT", HttpURLParsesIni, body)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
}
