package Reflect

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func PrintType(x interface{}) string {
	type stringer interface {
		String() string
	}
	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {return "true"}
		return "false"
	default:
		// array chan func map pointer slice struct
		return "need reflect"
	}
}

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
			return strconv.FormatInt(v.Int(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T): \n", name, x)
	display(name, reflect.ValueOf(x))
}

func display(path string, v reflect.Value) {
	switch v.Kind(){
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	}
}

func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()

	t2 := reflect.TypeOf(x)
	fmt.Printf("type %s/%s meths: %d/%d\n", t, t2, v.NumMethod(), t2.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s %s\n", t, t.Method(i).Name, strings.TrimPrefix(methType.String(), "func"))
	}
}

func PrintFromType(x interface{}){
	t := reflect.TypeOf(x)
	fmt.Printf("type %s/%s meths: %d\n", x, t, t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		methType := t.Method(i).Type
		fmt.Printf("func (%s) %s %s\n", t, t.Method(i).Name, strings.TrimPrefix(methType.String(), "func"))
	}
}











