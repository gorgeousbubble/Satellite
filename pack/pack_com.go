package pack

// pack aes
type TPackAESOne struct {
	Name		[]byte		// [32]byte/256bit
	Key			[]byte		// [16]byte/128bit
	OriginSize	[]byte		// [8]byte/64bit
	CryptSize	[]byte		// [8]byte/64bit
}