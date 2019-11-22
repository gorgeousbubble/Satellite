package parses

import (
	"bytes"
	"fmt"
	"testing"
)

func TestExtractOneElement(t *testing.T) {
	s := []byte("hello,1,[1,2,3],{simple,12}")
	r, sub, err := extractOneElement(s)
	if err != nil {
		t.Fatal("Error extract one element:", err)
	}
	fmt.Println("extract return:", string(r))
	fmt.Println("extract substring:", string(sub))
	if !bytes.Equal(r, []byte("hello")) {
		t.Fatal("Error extract return")
	}
	if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
		t.Fatal("Error extract substring")
	}
}

func BenchmarkExtractOneElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("hello,1,[1,2,3],{simple,12}")
		r, sub, err := extractOneElement(s)
		if err != nil {
			b.Fatal("Error extract one element:", err)
		}
		if !bytes.Equal(r, []byte("hello")) {
			b.Fatal("Error extract return")
		}
		if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
			b.Fatal("Error extract substring")
		}
	}
}

func TestExtractOneList(t *testing.T) {
	s := []byte("[apple,orange],1,[1,2,3],{simple,12}")
	r, sub, err := extractOneList(s)
	if err != nil {
		t.Fatal("Error extract one list:", err)
	}
	fmt.Println("extract return:", string(r))
	fmt.Println("extract substring:", string(sub))
	if !bytes.Equal(r, []byte("[apple,orange]")) {
		t.Fatal("Error extract return")
	}
	if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
		t.Fatal("Error extract substring")
	}
}

func BenchmarkExtractOneList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("[apple,orange],1,[1,2,3],{simple,12}")
		r, sub, err := extractOneList(s)
		if err != nil {
			b.Fatal("Error extract one list:", err)
		}
		if !bytes.Equal(r, []byte("[apple,orange]")) {
			b.Fatal("Error extract return")
		}
		if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
			b.Fatal("Error extract substring")
		}
	}
}

func TestExtractOneTuple(t *testing.T) {
	s := []byte("{stars,514},1,[1,2,3],{simple,12}")
	r, sub, err := extractOneTuple(s)
	if err != nil {
		t.Fatal("Error extract one tuple:", err)
	}
	fmt.Println("extract return:", string(r))
	fmt.Println("extract substring:", string(sub))
	if !bytes.Equal(r, []byte("{stars,514}")) {
		t.Fatal("Error extract return")
	}
	if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
		t.Fatal("Error extract substring")
	}
}

func BenchmarkExtractOneTuple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("{stars,514},1,[1,2,3],{simple,12}")
		r, sub, err := extractOneTuple(s)
		if err != nil {
			b.Fatal("Error extract one tuple:", err)
		}
		if !bytes.Equal(r, []byte("{stars,514}")) {
			b.Fatal("Error extract return")
		}
		if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
			b.Fatal("Error extract substring")
		}
	}
}

func TestTrimList(t *testing.T) {
	s := []byte("[apple,orange]")
	r, err := trimList(s)
	if err != nil {
		t.Fatal("Error trim list:", err)
	}
	fmt.Println("trim return:", string(r))
	if !bytes.Equal(r, []byte("apple,orange")) {
		t.Fatal("Error trim result")
	}
}

func BenchmarkTrimList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("[apple,orange]")
		r, err := trimList(s)
		if err != nil {
			b.Fatal("Error trim list:", err)
		}
		if !bytes.Equal(r, []byte("apple,orange")) {
			b.Fatal("Error trim result")
		}
	}
}

func TestTrimTuple(t *testing.T) {
	s := []byte("{stars,514}")
	r, err := trimTuple(s)
	if err != nil {
		t.Fatal("Error trim tuple:", err)
	}
	fmt.Println("trim return:", string(r))
	if !bytes.Equal(r, []byte("stars,514")) {
		t.Fatal("Error trim result")
	}
}

func BenchmarkTrimTuple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("{stars,514}")
		r, err := trimTuple(s)
		if err != nil {
			b.Fatal("Error trim tuple:", err)
		}
		if !bytes.Equal(r, []byte("stars,514")) {
			b.Fatal("Error trim result")
		}
	}
}

func TestRepairTrim(t *testing.T) {
	s := []byte("apple,1")
	r := repairTrim(s)
	if !bytes.Equal(r, []byte("apple,1,")) {
		t.Fatal("Error repair trim list or tuple")
	}
}

func BenchmarkRepairTrim(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("apple,1")
		r := repairTrim(s)
		if !bytes.Equal(r, []byte("apple,1,")) {
			b.Fatal("Error repair trim list or tuple")
		}
	}
}

func TestDecodeOneParameter(t *testing.T) {
	type subtest struct {
		Name   string `erl:"string"`
		Number int    `erl:"int"`
	}
	type test struct {
		Name    string  `erl:"string"`
		Content string  `erl:"string"`
		Number  int     `erl:"int"`
		List    []int   `erl:"list"`
		Tuple   subtest `erl:"tuple"`
	}
	in := []byte("test_name,hello,1,[1,2,3],{simple,12}")
	out := test{}
	err := decodeOneParameter(in, &out)
	if err != nil {
		t.Fatal("Error decode on parameter:", err)
	}
	fmt.Println(out)
}

func BenchmarkDecodeOneParameter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		type subtest struct {
			Name   string `erl:"string"`
			Number int    `erl:"int"`
		}
		type test struct {
			Name    string  `erl:"string"`
			Content string  `erl:"string"`
			Number  int     `erl:"int"`
			List    []int   `erl:"list"`
			Tuple   subtest `erl:"tuple"`
		}
		in := []byte("test_name,hello,1,[1,2,3],{simple,12}")
		out := test{}
		err := decodeOneParameter(in, &out)
		if err != nil {
			b.Fatal("Error decode on parameter:", err)
		}
	}
}

func TestDecodeOneParameter2(t *testing.T) {
	type subtest struct {
		Name   string `erl:"string"`
		Number int    `erl:"int"`
	}
	type test struct {
		Name    string    `erl:"string"`
		Content string    `erl:"string"`
		Number  int       `erl:"int"`
		List    []subtest `erl:"list"`
	}
	in := []byte("asteroid,earth,9,[{apple,1},{orange,2},{lemon,5}]")
	out := test{}
	err := decodeOneParameter(in, &out)
	if err != nil {
		t.Fatal("Error decode on parameter:", err)
	}
	fmt.Println(out)
}

func TestDecodeOneParameter3(t *testing.T) {
	type subtest struct {
		Name   string `erl:"string"`
		Number int    `erl:"int"`
		List   []int  `erl:"list"`
	}
	type test struct {
		List []subtest `erl:"list"`
	}
	in := []byte("[{apple,1,[2,3]},{orange,2,[1,5,4]},{lemon,5,[]}]")
	out := test{}
	err := decodeOneParameter(in, &out)
	if err != nil {
		t.Fatal("Error decode on parameter:", err)
	}
	fmt.Println(out)
}

func TestDecodeOneParameter4(t *testing.T) {
	type subsub struct {
		Name    string `erl:"string"`
		Content string `erl:"string"`
		Value   int    `erl:"int"`
	}
	type subtest1 struct {
		Name  string   `erl:"string"`
		Index int      `erl:"int"`
		Sub   []subsub `erl:"list"`
	}
	type subtest2 struct {
		Name    string `erl:"string"`
		SubName string `erl:"string"`
	}
	type subtest3 struct {
		Name string `erl:"string"`
		Up   int    `erl:"int"`
		Down int    `erl:"int"`
		Sub  subsub `erl:"tuple"`
	}
	type test struct {
		Name    string        `erl:"string"`
		Index   int           `erl:"int"`
		Options []interface{} `erl:"list"`
	}
	in := []byte("test,1,[{sub_1,2,[{subs,apple,3},{subs,lemon,5},{subs,banana,1}]},{sub_3,2,1,{subs,peach,7}},{sub_2,asteroid}]")
	out := test{}
	fmt.Println(out)
	ParsesErl = make(map[string]interface{})
	ParsesErl["sub_1"] = subtest1{}
	ParsesErl["sub_2"] = subtest2{}
	ParsesErl["sub_3"] = subtest3{}
	err := decodeOneParameter(in, &out)
	if err != nil {
		t.Fatal("Error decode on parameter:", err)
	}
	fmt.Println(out)
}
