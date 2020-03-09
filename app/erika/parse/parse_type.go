package parse

// Auto execute task template structure...
type TAutoExec struct {
	Name        string             `json:"name" yaml:"name"`
	Description string             `json:"description" yaml:"description"`
	Schedule    string             `json:"schedule" yaml:"schedule"`
	Attribute   TAutoExecAttribute `json:"attribute" yaml:"attribute"`
	Task        TAutoExecTask      `json:"task" yaml"task"`
}

type TAutoExecAttribute struct {
	Location string `json:"location" yaml:"location"`
}

type TAutoExecTask struct {
	Exec []TAutoExecTaskCommand `json:"exec" yaml:"exec"`
}

type TAutoExecTaskCommand struct {
	Command string `json:"command" yaml:"command"`
}
