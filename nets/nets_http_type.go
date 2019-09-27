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
