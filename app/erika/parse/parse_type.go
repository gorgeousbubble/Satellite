package parse

// Auto execute task template structure...
type TAutoExec struct {
	Name        string             `json:"name" yaml:"name"`
	Id          string             `json:"id" yaml:"id"`
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

// Auto monitor task template structure
type TAutoMonitor struct {
	Name        string                `json:"name" yaml:"name"`
	Id          string                `json:"id" yaml:"id"`
	Description string                `json:"description" yaml:"description"`
	Schedule    string                `json:"schedule" yaml:"schedule"`
	Attribute   TAutoMonitorAttribute `json:"attribute" yaml:"attribute"`
}

type TAutoMonitorAttribute struct {
	Location string `json:"location" yaml:"location"`
	Period   string `json:"period" yaml:"period"`
	Format   string `json:"format" yaml:"format"`
	Width    string `json:"width" yaml:"width"`
	Height   string `json:"height" yaml:"height"`
}
