package parse

import "testing"

func TestUnmarshalAutoMonitorJson(t *testing.T) {
	path := "../test/data/parse/auto_monitor.json"
	style := "json"
	out := TAutoMonitor{}
	err := unmarshalAutoMonitor(path, &out, style)
	if err != nil {
		t.Fatal("Error unmarshal auto monitor:", err)
	}
}

func BenchmarkUnmarshalAutoMonitorJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		path := "../test/data/parse/auto_monitor.json"
		style := "json"
		out := TAutoMonitor{}
		err := unmarshalAutoMonitor(path, &out, style)
		if err != nil {
			b.Fatal("Error unmarshal auto monitor:", err)
		}
	}
}

func TestMarshalAutoMonitorJson(t *testing.T) {
	path := "../test/data/parse/auto_exec.json"
	style := "json"
	in := TAutoMonitor{
		Name:        "auto_rasp_monitor",
		Description: "Automatic monitor room by raspberry camera.",
		Schedule:    "@hourly",
		Attribute: TAutoMonitorAttribute{
			Location: "/home/pi/monitor",
			Period: "600000",
			Format: "h264",
			Width: "1920",
			Height: "1080",
		},
	}
	err := marshalAutoMonitor(path, &in, style)
	if err != nil {
		t.Fatal("Error marshal auto monitor:", err)
	}
}
