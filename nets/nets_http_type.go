package nets

type TNetsPack struct {
	Src  []string `json:"src"`
	Dest string   `json:"dest"`
	Type string   `json:"type"`
}

type TNetsUnpack struct {
	Src  string `json:"src"`
	Dest string `json:"dest"`
}

type TNetsUnpackVerboseReq struct {
	Src string `json:"src"`
}

type TNetsUnpackVerboseResp struct {
	files []TNetUnpackFileInfo `json:"files"`
}

type TNetUnpackFileInfo struct {
	name      string `json:"name"`
	size      string `json:"size"`
	algorithm string `json:"type"`
}
