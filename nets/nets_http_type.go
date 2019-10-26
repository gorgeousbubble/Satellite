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

type TNetsPackProcessReq struct {
	Src  []string `json:"src"`
	Type string   `json:"type"`
}

type TNetsPackProcessResp struct {
	Done int64 `json:"done"`
	Work int64 `json:"work"`
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

type TNetsUnpackProcessReq struct {
	Src string `json:"src"`
}

type TNetsUnpackProcessResp struct {
	Done int64 `json:"done"`
	Work int64 `json:"work"`
}

type TNetsUnpackToFile struct {
	Src    string `json:"src"`
	Target string `json:"target"`
	Dest   string `json:"dest"`
}

type TNetsUnpackToMemory struct {
	Src    string `json:"src"`
	Target string `json:"target"`
}

type TNetsComp struct {
	Src  []string `json:"src"`
	Dest string   `json:"dest"`
	Type string   `json:"type"`
}

type TNetsDecomp struct {
	Src  string `json:"src"`
	Dest string `json:"dest"`
	Type string `json:"type"`
}
