package pack

// pack aes
type TPackAES struct {
	Name		[]byte		// [32]byte/256bit
	Author		[]byte		// [16]byte/128bit
	Number		[]byte		// [4]byte/32bit
}

type TPackAESOne struct {
	Name		[]byte		// [32]byte/256bit
	Key			[]byte		// [16]byte/128bit
	OriginSize	[]byte		// [4]byte/32bit
	CryptSize	[]byte		// [4]byte/32bit
}

// pack rsa
type TPackRSA struct {
	Name		[]byte		// [32]byte/256bit
	Author		[]byte		// [16]byte/128bit
	Number		[]byte		// [4]byte/32bit
}

type TPackRSAOne struct {
	Name		[]byte		// [32]byte/256bit
	Key			[]byte		// [1024]byte/128bit
	OriginSize	[]byte		// [4]byte/32bit
	CryptSize	[]byte		// [4]byte/32bit
}

// pack base64
type TPackBase64 struct {
	Name		[]byte		// [32]byte/256bit
	Author		[]byte		// [16]byte/128bit
	Number		[]byte		// [4]byte/32bit
}

type TPackBase64One struct {
	Name		[]byte		// [32]byte/256bit
	Size		[]byte		// [4]byte/32bit
}