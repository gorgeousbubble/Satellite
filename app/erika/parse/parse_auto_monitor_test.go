package parse

import "testing"

func TestUnmarshalAutoMonitorJson(t *testing.T) {
	path := "../test/data/parse/auto_monitor.json"
	style := "json"
	out := TAutoMonitor{}
	err := unmarshalAutoMonitor(path, &out, style)
	if err != nil {
		t.Fatal("Error unmarshal monitor exec:", err)
	}
}

func BenchmarkUnmarshalAutoMonitorJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		path := "../test/data/parse/auto_monitor.json"
		style := "json"
		out := TAutoMonitor{}
		err := unmarshalAutoMonitor(path, &out, style)
		if err != nil {
			b.Fatal("Error unmarshal monitor exec:", err)
		}
	}
}
