package nets

type TNetsRpcPackMD5EncodeReq struct {
	Src string `json:"src"`
}

type TNetsRpcPackMD5EncodeResp struct {
	Dest string `json:"dest"`
}

type TNetsRpcPackMD5EqualReq struct {
	Src  string `json:"src"`
	Dest string `json:"dest"`
}

type TNetsRpcPackMD5EqualResp struct {
	Equal bool `json:"equal"`
}
