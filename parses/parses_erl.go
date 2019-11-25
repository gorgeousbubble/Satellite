package parses

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// MParsesErl global mapping struct search by type 'interface{}'
// put all struct which will be push into 'interface{}' before call function 'Marshal' or 'Unmarshal'(recommend initialized in function 'init()')
// ...
// example:
// MParsesErl = make(map[string]interface{})
// MParsesErl["sub_1"] = subtest1{}
// MParsesErl["sub_2"] = subtest2{}
// MParsesErl["sub_3"] = subtest3{}
// ...
var MParsesErl map[string]interface{}

// Unmarshal erlang stream
func Unmarshal(in []byte, out interface{}) (err error) {
	return unmarshal(in, out)
}

// Marshal erlang stream
func Marshal(in interface{}) (out []byte, err error) {
	return marshal(in)
}

// decode
func extractOneElement(s []byte) (r []byte, sub []byte, err error) {
	// check bytes whether empty?
	if bytes.Equal(s, []byte("")) {
		err = errors.New("bytes is empty now")
		return r, sub, err
	}
	// erlang bytes contains ',' for element
	if !bytes.Contains(s, []byte(",")) {
		err = errors.New("bytes don't have element")
		return r, sub, err
	}
	// find element by ','
	index := bytes.Index(s, []byte(","))
	for i := 0; i < index; i++ {
		r = append(r, s[i])
	}
	// check element without '{}' or '[]'
	if bytes.ContainsAny(r, "{&}&[&]") {
		err = errors.New("bytes may extract list or tuple first")
		return r, sub, err
	}
	sub = bytes.TrimPrefix(s, r)
	sub = bytes.TrimPrefix(sub, []byte(","))
	return r, sub, err
}

func extractOneString(s []byte) (r string, sub []byte, err error) {
	// first, call function extract one element
	sr, sub, err := extractOneElement(s)
	if err != nil {
		return r, sub, err
	}
	// convert []byte -> string
	r = string(sr)
	return r, sub, err
}

func extractOneInt(s []byte) (r int, sub []byte, err error) {
	// first, call function extract one element
	sr, sub, err := extractOneElement(s)
	if err != nil {
		return r, sub, err
	}
	// convert []byte -> int
	r, err = strconv.Atoi(string(sr))
	if err != nil {
		err = errors.New("parameter may not integer")
		return r, sub, err
	}
	return r, sub, err
}

func extractOneFloat64(s []byte) (r float64, sub []byte, err error) {
	// first, call function extract one element
	sr, sub, err := extractOneElement(s)
	if err != nil {
		return r, sub, err
	}
	// convert []byte -> float64
	r, err = strconv.ParseFloat(string(sr), 64)
	if err != nil {
		err = errors.New("parameter may not float64")
		return r, sub, err
	}
	return r, sub, err
}

func extractOneBool(s []byte) (r bool, sub []byte, err error) {
	// first, call function extract one element
	sr, sub, err := extractOneElement(s)
	if err != nil {
		return r, sub, err
	}
	// convert []byte -> bool
	if string(sr) == "false" {
		r = false
	} else if string(sr) == "true" {
		r = true
	} else {
		err = errors.New("parameter may not bool")
		return r, sub, err
	}
	return r, sub, err
}

func extractOneList(s []byte) (r []byte, sub []byte, err error) {
	// erlang bytes contains '[]' for list
	if !bytes.Contains(s, []byte("[")) || !bytes.Contains(s, []byte("]")) {
		err = errors.New("bytes don't have list")
		return r, sub, err
	}
	// check '[' & ']' numbers should equal
	if bytes.Count(s, []byte("[")) != bytes.Count(s, []byte("]")) {
		err = errors.New("bytes list illegal, list may not intergrity")
		return r, sub, err
	}
	// check '[' & ']' sequence
	if bytes.Index(s, []byte("[")) > bytes.Index(s, []byte("]")) {
		err = errors.New("bytes list illegal, sequence of list may not right")
		return r, sub, err
	}
	// extract erlang list
	count := 0
	start := bytes.Index(s, []byte("["))
	for i := start; i < len(s); i++ {
		if s[i] == '[' {
			count++
		} else if s[i] == ']' {
			count--
		}
		r = append(r, s[i])
		if count == 0 {
			break
		}
	}
	sub = bytes.TrimPrefix(s, r)
	sub = bytes.TrimPrefix(sub, []byte(","))
	return r, sub, err
}

func extractOneTuple(s []byte) (r []byte, sub []byte, err error) {
	// erlang bytes contains '{}' for tuple
	if !bytes.Contains(s, []byte("{")) || !bytes.Contains(s, []byte("}")) {
		err = errors.New("bytes don't have tuple")
		return r, sub, err
	}
	// check '{' & '}' numbers should equal
	if bytes.Count(s, []byte("{")) != bytes.Count(s, []byte("}")) {
		err = errors.New("bytes tuple illegal, tuple may not intergrity")
		return r, sub, err
	}
	// check '{' & '}' sequence
	if bytes.Index(s, []byte("{")) > bytes.Index(s, []byte("}")) {
		err = errors.New("bytes tuple illegal, sequence of tuple may not right")
		return r, sub, err
	}
	// extract erlang tuple
	count := 0
	start := bytes.Index(s, []byte("{"))
	for i := start; i < len(s); i++ {
		if s[i] == '{' {
			count++
		} else if s[i] == '}' {
			count--
		}
		r = append(r, s[i])
		if count == 0 {
			break
		}
	}
	sub = bytes.TrimPrefix(s, r)
	sub = bytes.TrimPrefix(sub, []byte(","))
	return r, sub, err
}

func trimList(s []byte) (r []byte, err error) {
	// check list whether start '[' and end ']'
	if len(s) == 0 || s[0] != '[' || s[len(s)-1] != ']' {
		err = errors.New("bytes is not normal list")
		return r, err
	}
	// trim '[]' and return
	r = bytes.TrimPrefix(s, []byte("["))
	r = bytes.TrimSuffix(r, []byte("]"))
	return r, err
}

func trimTuple(s []byte) (r []byte, err error) {
	// check tuple whether start '{' and end '}'
	if len(s) == 0 || s[0] != '{' || s[len(s)-1] != '}' {
		err = errors.New("bytes is not normal tuple")
		return r, err
	}
	// trim '{}' and return
	r = bytes.TrimPrefix(s, []byte("{"))
	r = bytes.TrimSuffix(r, []byte("}"))
	return r, err
}

func repairTrim(s []byte) (r []byte) {
	// repair bytes after call trim function
	r = append(s, ',')
	return r
}

func decodeOneParameter(in []byte, out interface{}) (err error) {
	// get pointer's value...
	var rType = reflect.TypeOf(out)
	var rValue = reflect.ValueOf(out)
	fmt.Println("type of out interface:", rType)
	fmt.Println("value of out interface:", rValue)
	fmt.Println(rType.Kind())
	// check the out type kind
	if rType.Kind() != reflect.Ptr {
		err = errors.New("out interface should be struct pointer")
		fmt.Println(err)
		return err
	}
	// get real variable value...
	rType = rType.Elem()
	rValue = rValue.Elem()
	fmt.Println("type of point:", rType)
	fmt.Println("value of point:", rValue)
	fmt.Println(rType.Kind())
	// switch the kind of type...
	switch rType.Kind() {
	case reflect.Struct:
		var rem = in
		// traverse struct fields
		for i := 0; i < rType.NumField(); i++ {
			// get struct field value...
			t := rType.Field(i)
			f := rValue.Field(i)
			fmt.Printf("struct field[%v]\n", i)
			fmt.Printf("type of field[%v]:%v\n", i, t)
			fmt.Printf("value of field[%v]:%v\n", i, f)
			// parse tag
			tag := t.Tag.Get("erl")
			fields := strings.Split(tag, ",")
			if len(fields) > 1 {
				tag = fields[0]
			}
			// swich the tag & parse element...
			switch tag {
			case "string":
				r, sub, err := extractOneString(rem)
				if err != nil {
					return err
				}
				rem = sub
				f.Set(reflect.ValueOf(r))
			case "int":
				r, sub, err := extractOneInt(rem)
				if err != nil {
					return err
				}
				rem = sub
				f.Set(reflect.ValueOf(r))
			case "float64":
				r, sub, err := extractOneFloat64(rem)
				if err != nil {
					return err
				}
				rem = sub
				f.Set(reflect.ValueOf(r))
			case "bool":
				r, sub, err := extractOneBool(rem)
				if err != nil {
					return err
				}
				rem = sub
				f.Set(reflect.ValueOf(r))
			case "list":
				r, sub, err := extractOneList(rem)
				if err != nil {
					return err
				}
				rem = sub
				r, err = trimList(r)
				if err != nil {
					return err
				}
				r = repairTrim(r)
				err = decodeOneParameter(r, f.Addr().Interface())
				if err != nil {
					return err
				}
			case "tuple":
				r, sub, err := extractOneTuple(rem)
				if err != nil {
					return err
				}
				rem = sub
				r, err = trimTuple(r)
				if err != nil {
					return err
				}
				r = repairTrim(r)
				err = decodeOneParameter(r, f.Addr().Interface())
				if err != nil {
					return err
				}
			default:
				err = errors.New("unrecognized struct field type")
				return err
			}
		}
	case reflect.Slice:
		var rem = in
		// switch the kind of sub type...
		fmt.Println(rValue.Type().Elem().Kind())
		switch rValue.Type().Elem().Kind() {
		case reflect.String:
			// traverse slice elements(string)
			var e error
			for {
				r, sub, e := extractOneString(rem)
				if e != nil {
					break
				}
				rValue = reflect.Append(rValue, reflect.ValueOf(r))
				fmt.Println("rValue:", rValue)
				rem = sub
				fmt.Println("rem:", string(rem))
			}
			// parse error leave
			if len(rem) != 0 {
				err = e
				fmt.Println(err)
				return err
			}
			reflect.ValueOf(out).Elem().Set(rValue)
		case reflect.Int:
			// traverse slice elements(int)
			var e error
			for {
				r, sub, e := extractOneInt(rem)
				if e != nil {
					break
				}
				rValue = reflect.Append(rValue, reflect.ValueOf(r))
				fmt.Println("rValue:", rValue)
				rem = sub
				fmt.Println("rem:", string(rem))
			}
			// parse error leave
			if len(rem) != 0 {
				err = e
				fmt.Println(err)
				return err
			}
			reflect.ValueOf(out).Elem().Set(rValue)
		case reflect.Float64:
			// traverse slice elements(float64)
			var e error
			for {
				r, sub, e := extractOneFloat64(rem)
				if e != nil {
					break
				}
				rValue = reflect.Append(rValue, reflect.ValueOf(r))
				fmt.Println("rValue:", rValue)
				rem = sub
				fmt.Println("rem:", string(rem))
			}
			// parse error leave
			if len(rem) != 0 {
				err = e
				fmt.Println(err)
				return err
			}
			reflect.ValueOf(out).Elem().Set(rValue)
		case reflect.Bool:
			// traverse slice elements(bool)
			var e error
			for {
				r, sub, e := extractOneBool(rem)
				if e != nil {
					break
				}
				rValue = reflect.Append(rValue, reflect.ValueOf(r))
				fmt.Println("rValue:", rValue)
				rem = sub
				fmt.Println("rem:", string(rem))
			}
			// parse error leave
			if len(rem) != 0 {
				err = e
				fmt.Println(err)
				return err
			}
			reflect.ValueOf(out).Elem().Set(rValue)
		case reflect.Struct:
			// traverse slice elements(struct)
			var e error
			for {
				r, sub, e := extractOneTuple(rem)
				if e != nil {
					break
				}
				rem = sub
				fmt.Println("rem:", string(rem))
				r, err = trimTuple(r)
				if err != nil {
					return err
				}
				r = repairTrim(r)
				fmt.Println(string(r))
				// new struct value
				o := reflect.New(rValue.Type().Elem())
				fmt.Printf("type:%v,value:%v\n", reflect.TypeOf(o), reflect.ValueOf(o))
				fmt.Println(o.CanAddr())
				err = decodeOneParameter(r, o.Interface())
				if err != nil {
					return err
				}
				fmt.Printf("o:%v,elem:%v\n", o, o.Elem())
				rValue = reflect.Append(rValue, reflect.ValueOf(o.Elem().Interface()))
				fmt.Println("rValue:", rValue)
			}
			// parse error leave
			if len(rem) != 0 {
				err = e
				fmt.Println(err)
				return err
			}
			reflect.ValueOf(out).Elem().Set(rValue)
		case reflect.Interface:
			// traverse slice elements(interface)
			var e error
			for {
				r, sub, e := extractOneTuple(rem)
				if e != nil {
					break
				}
				rem = sub
				fmt.Println("r:", string(r))
				fmt.Println("rem:", string(rem))
				r, err = trimTuple(r)
				if err != nil {
					return err
				}
				r = repairTrim(r)
				fmt.Println(string(r))
				// check type by name of interface(extract type name)
				name, _, err := extractOneString(r)
				if err != nil {
					return err
				}
				fmt.Println("name of type:", string(name))
				// traverse map elements
				for k, v := range MParsesErl {
					if k == string(name) {
						// new struct value
						o := reflect.New(reflect.TypeOf(v))
						fmt.Printf("type:%v,value:%v\n", reflect.TypeOf(o), reflect.ValueOf(o))
						fmt.Println(o.CanAddr())
						err = decodeOneParameter(r, o.Interface())
						if err != nil {
							return err
						}
						fmt.Printf("o:%v,elem:%v\n", o, o.Elem())
						rValue = reflect.Append(rValue, reflect.ValueOf(o.Elem().Interface()))
						fmt.Println("rValue:", rValue)
						break
					}
				}
			}
			// parse error leave
			if len(rem) != 0 {
				err = e
				fmt.Println(err)
				return err
			}
			reflect.ValueOf(out).Elem().Set(rValue)
		default:
			err = errors.New("unrecognized list element type")
			return err
		}
	/*case reflect.Interface:
	fmt.Println("hello reflect interface...")
	var rem = in
	// first, extract one parameter as string...
	r, sub, err := extractOneString(rem)
	if err != nil {
		return err
	}
	rem = sub
	fmt.Println("r:", string(r))
	fmt.Println("rem:", string(rem))
	fmt.Println("breakout...")*/
	default:
		err = errors.New("unrecognized reflect type")
		return err
	}
	return err
}

func decode(in []byte, out interface{}) (err error) {
	// get pointer's value...
	var rType = reflect.TypeOf(out)
	var rValue = reflect.ValueOf(out)
	fmt.Println("type of out interface:", rType)
	fmt.Println("value of out interface:", rValue)
	fmt.Println(rType.Kind())
	// check the out type kind
	if rType.Kind() != reflect.Ptr {
		err = errors.New("out interface should be struct pointer")
		fmt.Println(err)
		return err
	}
	// get real variable value...
	rType = rType.Elem()
	rValue = rValue.Elem()
	fmt.Println("type of point:", rType)
	fmt.Println("value of point:", rValue)
	fmt.Println(rType.Kind())
	// type should be struct
	if rType.Kind() != reflect.Struct {
		err = errors.New("real variable should be struct")
		fmt.Println(err)
		return err
	}
	// extract type name
	var r = in
	name, _, err := extractOneString(r)
	if err != nil {
		return err
	}
	fmt.Println("name of type:", string(name))
	fmt.Println("global mapping:", MParsesErl)
	// traverse struct fields
	for i := 0; i < rType.NumField(); i++ {
		// get struct field value...
		t := rType.Field(i)
		f := rValue.Field(i)
		fmt.Printf("struct field[%v]\n", i)
		fmt.Printf("type of field[%v]:%v\n", i, t)
		fmt.Printf("value of field[%v]:%v\n", i, f)
		// check list elements type
		for k, v := range MParsesErl {
			if k == string(name) && reflect.TypeOf(v).Name() == f.Type().Elem().Name() {
				// new struct value
				o := reflect.New(reflect.TypeOf(v))
				fmt.Printf("type:%v,value:%v\n", reflect.TypeOf(o), reflect.ValueOf(o))
				fmt.Println(o.CanAddr())
				err = decodeOneParameter(r, o.Interface())
				if err != nil {
					return err
				}
				fmt.Printf("o:%v,elem:%v\n", o, o.Elem())
				f = reflect.Append(f, reflect.ValueOf(o.Elem().Interface()))
				fmt.Println("rValue:", f)
				break
			}
		}
		reflect.ValueOf(out).Elem().Field(i).Set(f)
	}
	return err
}

func unmarshal(in []byte, out interface{}) (err error) {
	var s [][]byte
	// split the stream by symbol '\n' in order to delete comments
	s1 := bytes.Split(in, []byte("\n"))
	for _, v := range s1 {
		// delete comments
		if bytes.Contains(v, []byte("%")) {
			index := bytes.Index(v, []byte("%"))
			v = append(v[:index], v[len(v):]...)
		}
		// delete C/C++ comments?
		if bytes.Contains(v, []byte("//")) {
			index := bytes.Index(v, []byte("//"))
			v = append(v[:index], v[len(v):]...)
		}
		// delete all space
		v = bytes.ReplaceAll(v, []byte(" "), []byte(""))
		// delete all return
		v = bytes.ReplaceAll(v, []byte("\r"), []byte(""))
		// delete all blank lines, packet to s slice
		if !bytes.Equal(v, []byte("")) {
			s = append(s, v)
		}
		fmt.Println(string(v))
	}
	data := bytes.Join(s, []byte(""))
	fmt.Println(string(data))
	// split the data by symbol '{' and '}.', according to syntax
	s = bytes.Split(data, []byte("}."))
	for _, v := range s {
		// valid line
		if !bytes.Contains(v, []byte("{")) {
			continue
		}
		// delete '{'
		index := bytes.Index(v, []byte("{"))
		v = append(v[:index], v[index+1:]...)
		fmt.Println(string(v))
		// repair with ','
		v = repairTrim(v)
		// decode
		err = decode(v, out)
		if err != nil {
			return err
		}
	}
	return err
}

// encode
func wrapOneElement(in interface{}) (r []byte, err error) {
	var rType = reflect.TypeOf(in)
	var rValue = reflect.ValueOf(in)
	fmt.Println("type of in interface:", rType)
	fmt.Println("value of in interface:", rValue)
	fmt.Println(rType.Kind())
	// switch the s type kind
	switch rType.Kind() {
	case reflect.String:
		r = []byte(rValue.Interface().(string))
		r = append(r, ',')
	case reflect.Bool:
		f := rValue.Interface().(bool)
		if f == true {
			r = []byte("true")
		} else {
			r = []byte("false")
		}
		r = append(r, ',')
	case reflect.Int:
		r = []byte(strconv.Itoa(rValue.Interface().(int)))
		r = append(r, ',')
	case reflect.Float64:
		r = []byte(strconv.FormatFloat(rValue.Interface().(float64), 'f', -1, 64))
		r = append(r, ',')
	case reflect.Struct:
		var s [][]byte
		// traverse struct fields
		for i := 0; i < rType.NumField(); i++ {
			// get struct field value...
			t := rType.Field(i)
			f := rValue.Field(i)
			fmt.Printf("struct field[%v]\n", i)
			fmt.Printf("type of field[%v]:%v\n", i, t)
			fmt.Printf("value of field[%v]:%v\n", i, f)
			// parse tag
			tag := t.Tag.Get("erl")
			fields := strings.Split(tag, ",")
			if len(fields) > 1 {
				tag = fields[0]
			}
			// swich the tag & parse element...
			switch tag {
			case "string":
				fallthrough
			case "int":
				fallthrough
			case "float64":
				fallthrough
			case "bool":
				fallthrough
			case "tuple":
				fallthrough
			case "list":
				fallthrough
			case "element":
				rs, err := wrapOneElement(f.Interface())
				if err != nil {
					return r, err
				}
				if i == rType.NumField()-1 {
					rs = trimElement(rs)
				}
				s = append(s, rs)
			default:
				err = errors.New("unrecognized struct field type")
				return r, err
			}
		}
		// repair tuple...
		r = bytes.Join(s, []byte(""))
		r = repairTuple(r)
		r = repairTrim(r)
		fmt.Println(string(r))
	case reflect.Slice:
		// switch the kind of sub type...
	default:
		err = errors.New("unrecognized element type")
		return r, err
	}
	return r, err
}

func trimElement(s []byte) (r []byte) {
	r = bytes.TrimSuffix(s, []byte(","))
	return r
}

func repairList(s []byte) (r []byte) {
	ss := make([][]byte, 3)
	ss[0] = []byte("[")
	ss[1] = s
	ss[2] = []byte("]")
	r = bytes.Join(ss, []byte(""))
	return r
}

func repairTuple(s []byte) (r []byte) {
	ss := make([][]byte, 3)
	ss[0] = []byte("{")
	ss[1] = s
	ss[2] = []byte("}")
	r = bytes.Join(ss, []byte(""))
	return r
}

func marshal(in interface{}) (out []byte, err error) {
	return out, err
}
