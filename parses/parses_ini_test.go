package parses

import (
	"fmt"
	"testing"
)

func TestGetValue(t *testing.T) {
	in := []byte(`
[LIVECOREMODE]
LiveCore_Mode=0

[LIVECORESHOW]
LiveCore_Show_Graphics=0
LiveCore_Show_GraphicsFont=24

[LIVECOREWALLPAPERMODE]
LiveCore_Wallpaper_Mode=0
LiveCore_Wallpaper_Audio=1

[LIVECORELOGMODE]
LiveCore_Log_Process=0

[LIVECOREPLAYMODE]
LiveCore_Play_Mode=2

[LIVECOREWINDOW]
LiveCore_Window_Handle=9506642

[LIVECOREVIDEOADDRESS]
LiveCore_Video_Mode=1
LiveCore_Video_Name=Wallpaper1.pak
LiveCore_Video_Address=C:\Users\10295\Videos\29156949_22_0.flv
`)
	r := getValue(in, "LIVECORESHOW", "LiveCore_Show_Graphics")
	if string(r) != "0" {
		t.Fatal("Error parses ini...")
	}
}

func BenchmarkGetValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in := []byte(`
[LIVECOREMODE]
LiveCore_Mode=0

[LIVECORESHOW]
LiveCore_Show_Graphics=0
LiveCore_Show_GraphicsFont=24

[LIVECOREWALLPAPERMODE]
LiveCore_Wallpaper_Mode=0
LiveCore_Wallpaper_Audio=1

[LIVECORELOGMODE]
LiveCore_Log_Process=0

[LIVECOREPLAYMODE]
LiveCore_Play_Mode=2

[LIVECOREWINDOW]
LiveCore_Window_Handle=9506642

[LIVECOREVIDEOADDRESS]
LiveCore_Video_Mode=1
LiveCore_Video_Name=Wallpaper1.pak
LiveCore_Video_Address=C:\Users\10295\Videos\29156949_22_0.flv
`)
		r := getValue(in, "LIVECORESHOW", "LiveCore_Show_Graphics")
		if string(r) != "0" {
			b.Fatal("Error parses ini...")
		}
	}
}

func TestSetValue(t *testing.T) {
	in := []byte(`
[LIVECOREMODE]
LiveCore_Mode=0

[LIVECORESHOW]
LiveCore_Show_Graphics=0
LiveCore_Show_GraphicsFont=24

[LIVECOREWALLPAPERMODE]
LiveCore_Wallpaper_Mode=0
LiveCore_Wallpaper_Audio=1

[LIVECORELOGMODE]
LiveCore_Log_Process=0

[LIVECOREPLAYMODE]
LiveCore_Play_Mode=2

[LIVECOREWINDOW]
LiveCore_Window_Handle=9506642

[LIVECOREVIDEOADDRESS]
LiveCore_Video_Mode=1
LiveCore_Video_Name=Wallpaper1.pak
LiveCore_Video_Address=C:\Users\10295\Videos\29156949_22_0.flv
`)
	r, err := setValue(in, "LIVECORESHOW", "LiveCore_Show_GraphicsFont", "20")
	if err != nil {
		t.Fatal("Error parses ini:", err)
	}
	fmt.Println(string(r))
}

func BenchmarkSetValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in := []byte(`
[LIVECOREMODE]
LiveCore_Mode=0

[LIVECORESHOW]
LiveCore_Show_Graphics=0
LiveCore_Show_GraphicsFont=24

[LIVECOREWALLPAPERMODE]
LiveCore_Wallpaper_Mode=0
LiveCore_Wallpaper_Audio=1

[LIVECORELOGMODE]
LiveCore_Log_Process=0

[LIVECOREPLAYMODE]
LiveCore_Play_Mode=2

[LIVECOREWINDOW]
LiveCore_Window_Handle=9506642

[LIVECOREVIDEOADDRESS]
LiveCore_Video_Mode=1
LiveCore_Video_Name=Wallpaper1.pak
LiveCore_Video_Address=C:\Users\10295\Videos\29156949_22_0.flv
`)
		_, err := setValue(in, "LIVECORESHOW", "LiveCore_Show_GraphicsFont", "20")
		if err != nil {
			b.Fatal("Error parses ini:", err)
		}
	}
}

func TestGetValueStringFrom(t *testing.T) {
	src := "../test/data/parses/test.ini"
	r, err := getValueStringFrom(src, "LIVECOREVIDEOADDRESS", "LiveCore_Video_Name")
	if err != nil {
		t.Fatal("Error parses ini:", err)
	}
	if r != "Wallpaper1.pak" {
		t.Fatal("Error parses ini content")
	}
}

func BenchmarkGetValueStringFrom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/parses/test.ini"
		r, err := getValueStringFrom(src, "LIVECOREVIDEOADDRESS", "LiveCore_Video_Name")
		if err != nil {
			b.Fatal("Error parses ini:", err)
		}
		if r != "Wallpaper1.pak" {
			b.Fatal("Error parses ini content")
		}
	}
}

func TestGetValueIntFrom(t *testing.T) {
	src := "../test/data/parses/test.ini"
	r, err := getValueIntFrom(src, "LIVECOREVIDEOADDRESS", "LiveCore_Video_Mode")
	if err != nil {
		t.Fatal("Error parses ini:", err)
	}
	if r != 1 {
		t.Fatal("Error parses ini content")
	}
}

func BenchmarkGetValueIntFrom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/parses/test.ini"
		r, err := getValueIntFrom(src, "LIVECOREVIDEOADDRESS", "LiveCore_Video_Mode")
		if err != nil {
			b.Fatal("Error parses ini:", err)
		}
		if r != 1 {
			b.Fatal("Error parses ini content")
		}
	}
}

func TestGetValueBoolFrom(t *testing.T) {
	src := "../test/data/parses/test_simple.ini"
	r, err := getValueBoolFrom(src, "BOOL", "Switch_On")
	if err != nil {
		t.Fatal("Error parses ini:", err)
	}
	if r != true {
		t.Fatal("Error parses ini content")
	}
}

func TestGetValueFloat64From(t *testing.T) {
	src := "../test/data/parses/test_simple.ini"
	r, err := getValueFloat64From(src, "FLOAT", "Pi")
	if err != nil {
		t.Fatal("Error parses ini:", err)
	}
	if r != 3.1415926 {
		t.Fatal("Error parses ini content")
	}
}

func TestGetValueFrom(t *testing.T) {
	var value string
	src := "../test/data/parses/test.ini"
	err := GetValueFrom(src, "LIVECOREVIDEOADDRESS", "LiveCore_Video_Name", &value)
	if err != nil {
		t.Fatal("Error parses ini:", err)
	}
	if value != "Wallpaper1.pak" {
		t.Fatal("Error parses ini content")
	}
}

func BenchmarkGetValueFrom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var value string
		src := "../test/data/parses/test.ini"
		err := GetValueFrom(src, "LIVECOREVIDEOADDRESS", "LiveCore_Video_Name", &value)
		if err != nil {
			b.Fatal("Error parses ini:", err)
		}
		if value != "Wallpaper1.pak" {
			b.Fatal("Error parses ini content")
		}
	}
}

func TestGetValueFrom2(t *testing.T) {
	var value int
	src := "../test/data/parses/test.ini"
	err := GetValueFrom(src, "LIVECOREVIDEOADDRESS", "LiveCore_Video_Mode", &value)
	if err != nil {
		t.Fatal("Error parses ini:", err)
	}
	if value != 1 {
		t.Fatal("Error parses ini content")
	}
}

func BenchmarkGetValueFrom2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var value int
		src := "../test/data/parses/test.ini"
		err := GetValueFrom(src, "LIVECOREVIDEOADDRESS", "LiveCore_Video_Mode", &value)
		if err != nil {
			b.Fatal("Error parses ini:", err)
		}
		if value != 1 {
			b.Fatal("Error parses ini content")
		}
	}
}

func TestGetValueFrom3(t *testing.T) {
	var value bool
	src := "../test/data/parses/test_simple.ini"
	err := GetValueFrom(src, "BOOL", "Switch_On", &value)
	if err != nil {
		t.Fatal("Error parses ini:", err)
	}
	if value != true {
		t.Fatal("Error parses ini content")
	}
}

func TestGetValueFrom4(t *testing.T) {
	var value float64
	src := "../test/data/parses/test_simple.ini"
	err := GetValueFrom(src, "FLOAT", "Pi", &value)
	if err != nil {
		t.Fatal("Error parses ini:", err)
	}
	if value != 3.1415926 {
		t.Fatal("Error parses ini content")
	}
}

func TestSetValueStringTo(t *testing.T) {
	src := "../test/data/parses/test.ini"
	err := setValueStringTo(src, "LIVECOREVIDEOADDRESS", "LiveCore_Video_Name", "Wallpaper1.pak")
	if err != nil {
		t.Fatal("Error parses ini:", err)
	}
}

func BenchmarkSetValueStringTo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/parses/test.ini"
		err := setValueStringTo(src, "LIVECOREVIDEOADDRESS", "LiveCore_Video_Name", "Wallpaper1.pak")
		if err != nil {
			b.Fatal("Error parses ini:", err)
		}
	}
}

func TestSetValueIntTo(t *testing.T) {
	src := "../test/data/parses/test.ini"
	err := setValueIntTo(src, "LIVECOREVIDEOADDRESS", "LiveCore_Video_Mode", 1)
	if err != nil {
		t.Fatal("Error parses ini:", err)
	}
}

func BenchmarkSetValueIntTo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/parses/test.ini"
		err := setValueIntTo(src, "LIVECOREVIDEOADDRESS", "LiveCore_Video_Mode", 1)
		if err != nil {
			b.Fatal("Error parses ini:", err)
		}
	}
}

func TestSetValueBoolTo(t *testing.T) {
	src := "../test/data/parses/test_simple.ini"
	err := setValueBoolTo(src, "BOOL", "Switch_On", true)
	if err != nil {
		t.Fatal("Error parses ini:", err)
	}
}

func TestSetValueFloat64To(t *testing.T) {
	src := "../test/data/parses/test_simple.ini"
	err := setValueFloat64To(src, "FLOAT", "Pi", 3.1415926)
	if err != nil {
		t.Fatal("Error parses ini:", err)
	}
}

func TestSetValueTo(t *testing.T) {
	src := "../test/data/parses/test.ini"
	err := SetValueTo(src, "LIVECOREVIDEOADDRESS", "LiveCore_Video_Name", "Wallpaper1.pak")
	if err != nil {
		t.Fatal("Error parses ini:", err)
	}
}

func BenchmarkSetValueTo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/parses/test.ini"
		err := SetValueTo(src, "LIVECOREVIDEOADDRESS", "LiveCore_Video_Name", "Wallpaper1.pak")
		if err != nil {
			b.Fatal("Error parses ini:", err)
		}
	}
}

func TestSetValueTo2(t *testing.T) {
	src := "../test/data/parses/test.ini"
	err := SetValueTo(src, "LIVECOREVIDEOADDRESS", "LiveCore_Video_Mode", 1)
	if err != nil {
		t.Fatal("Error parses ini:", err)
	}
}

func BenchmarkSetValueTo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/parses/test.ini"
		err := SetValueTo(src, "LIVECOREVIDEOADDRESS", "LiveCore_Video_Mode", 1)
		if err != nil {
			b.Fatal("Error parses ini:", err)
		}
	}
}

func TestSetValueTo3(t *testing.T) {
	src := "../test/data/parses/test_simple.ini"
	err := SetValueTo(src, "BOOL", "Switch_On", true)
	if err != nil {
		t.Fatal("Error parses ini:", err)
	}
}

func TestSetValueTo4(t *testing.T) {
	src := "../test/data/parses/test_simple.ini"
	err := SetValueTo(src, "FLOAT", "Pi", 3.1415926)
	if err != nil {
		t.Fatal("Error parses ini:", err)
	}
}
