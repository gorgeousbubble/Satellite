package parse

// Auto execute task template structure...
type TAutoExec struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Schedule    string             `json:"schedule"`
	Attribute   TAutoExecAttribute `json:"attribute"`
}

type TAutoExecAttribute struct {
	Location string `json:"location"`
}

type TAutoExecTask struct {
	Exec []TAutoExecTaskCommand `json:"exec"`
}

type TAutoExecTaskCommand struct {
	Command string `json:"command"`
}
