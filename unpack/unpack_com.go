package unpack

// unpack aes
type TUnpackAES struct {
	Name		[]byte		// [32]byte/256bit
	Author		[]byte		// [16]byte/128bit
	Type		[]byte		// [8]byte/64bit
	Number		[]byte		// [4]byte/32bit
}

type TUnpackAESOne struct {
	Name		[]byte		// [32]byte/256bit
	Key			[]byte		// [16]byte/128bit
	OriginSize	[]byte		// [4]byte/32bit
	CryptSize	[]byte		// [4]byte/32bit
}

// unpack des
type TUnpack3DES struct {
	Name		[]byte		// [32]byte/256bit
	Author		[]byte		// [16]byte/128bit
	Type		[]byte		// [8]byte/64bit
	Number		[]byte		// [4]byte/32bit
}

type TUnpack3DESOne struct {
	Name		[]byte		// [32]byte/256bit
	Key			[]byte		// [24]byte/192bit
	OriginSize	[]byte		// [4]byte/32bit
	CryptSize	[]byte		// [4]byte/32bit
}

type TUnpackDES struct {
	Name		[]byte		// [32]byte/256bit
	Author		[]byte		// [16]byte/128bit
	Type		[]byte		// [8]byte/64bit
	Number		[]byte		// [4]byte/32bit
}

type TUnpackDESOne struct {
	Name		[]byte		// [32]byte/256bit
	Key			[]byte		// [8]byte/64bit
	OriginSize	[]byte		// [4]byte/32bit
	CryptSize	[]byte		// [4]byte/32bit
}

// unpack rsa
type TUnpackRSA struct {
	Name		[]byte		// [32]byte/256bit
	Author		[]byte		// [16]byte/128bit
	Type		[]byte		// [8]byte/64bit
	Number		[]byte		// [4]byte/32bit
}

type TUnpackRSAOne struct {
	Name		[]byte		// [32]byte/256bit
	Key			[]byte		// [1024]byte/128bit
	OriginSize	[]byte		// [4]byte/32bit
	CryptSize	[]byte		// [4]byte/32bit
}

// unpack base64
type TUnpackBase64 struct {
	Name		[]byte		// [32]byte/256bit
	Author		[]byte		// [16]byte/128bit
	Type		[]byte		// [8]byte/64bit
	Number		[]byte		// [4]byte/32bit
}

type TUnpackBase64One struct {
	Name		[]byte		// [32]byte/256bit
	Size		[]byte		// [4]byte/32bit
}
