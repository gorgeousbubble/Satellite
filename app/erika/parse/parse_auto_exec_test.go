package parse

import (
	"testing"
)

func TestUnmarshalAutoExecJson(t *testing.T) {
	path := "../test/data/parse/auto_exec.json"
	style := "json"
	out := TAutoExec{}
	err := unmarshalAutoExec(path, &out, style)
	if err != nil {
		t.Fatal("Error unmarshal auto exec:", err)
	}
}

func BenchmarkUnmarshalAutoExecJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		path := "../test/data/parse/auto_exec.json"
		style := "json"
		out := TAutoExec{}
		err := unmarshalAutoExec(path, &out, style)
		if err != nil {
			b.Fatal("Error unmarshal auto exec:", err)
		}
	}
}

func TestMarshalAutoExecJson(t *testing.T) {
	path := "../test/data/parse/auto_exec.json"
	style := "json"
	in := TAutoExec{
		Name:        "auto_exec_sync_git",
		Description: "Automatic synchronization Git with remote repository.",
		Schedule:    "@hourly",
		Attribute: TAutoExecAttribute{
			Location: "/home/pi/repos/satellite",
		},
		Task: TAutoExecTask{
			Exec: []TAutoExecTaskCommand{
				{
					Command: "git pull gogs master",
				},
				{
					Command: "git add .",
				},
				{
					Command: "git commit -m 'Automatic Commit by Erika.'",
				},
				{
					Command: "git push gogs master",
				},
			},
		},
	}
	err := marshalAutoExec(path, &in, style)
	if err != nil {
		t.Fatal("Error marshal auto exec:", err)
	}
}

func BenchmarkMarshalAutoExecJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		path := "../test/data/parse/auto_exec.json"
		style := "json"
		in := TAutoExec{
			Name:        "auto_exec_sync_git",
			Description: "Automatic synchronization Git with remote repository.",
			Schedule:    "@hourly",
			Attribute: TAutoExecAttribute{
				Location: "/home/pi/repos/satellite",
			},
			Task: TAutoExecTask{
				Exec: []TAutoExecTaskCommand{
					{
						Command: "git pull gogs master",
					},
					{
						Command: "git add .",
					},
					{
						Command: "git commit -m 'Automatic Commit by Erika.'",
					},
					{
						Command: "git push gogs master",
					},
				},
			},
		}
		err := marshalAutoExec(path, &in, style)
		if err != nil {
			b.Fatal("Error marshal auto exec:", err)
		}
	}
}
