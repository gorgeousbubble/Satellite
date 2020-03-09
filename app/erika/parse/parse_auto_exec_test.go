package parse

import (
	"testing"
)

func TestUnmarshalAutoExecJson(t *testing.T) {
	path := "../template/auto_exec.json"
	style := "json"
	out := TAutoExec{}
	err := unmarshalAutoExec(path, &out, style)
	if err != nil {
		t.Fatal("Error unmarshal auto exec:", err)
	}
}

func BenchmarkUnmarshalAutoExecJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		path := "../template/auto_exec.json"
		style := "json"
		out := TAutoExec{}
		err := unmarshalAutoExec(path, &out, style)
		if err != nil {
			b.Fatal("Error unmarshal auto exec:", err)
		}
	}
}
