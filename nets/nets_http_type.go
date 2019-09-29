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
	Files []TNetsUnpackFileInfo `json:"files"`
}

type TNetsUnpackFileInfo struct {
	Name string `json:"name"`
	Size string `json:"size"`
	Type string `json:"type"`
}
