package unpack

// pack aes
type TUnpackAES struct {
	Name		[]byte		// [32]byte/256bit
	Author		[]byte		// [16]byte/128bit
	Number		[]byte		// [4]byte/32bit
}

type TUnpackAESOne struct {
	Name		[]byte		// [32]byte/256bit
	Key			[]byte		// [16]byte/128bit
	OriginSize	[]byte		// [4]byte/32bit
	CryptSize	[]byte		// [4]byte/32bit
}