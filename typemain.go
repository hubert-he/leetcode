package main

import (
	"./syntax/Reflect"
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"reflect"
	"time"
)

func f0() (int){
	r := 0
	defer func(){
		r++
	}()
	return r
}
func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}
func f2() (r int){
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
func f3()(r int){
	defer func(r int){
		r += 5
	}(r)
	return 1
}
func f4()(r int){
	defer func(r *int){
		*r += 5
	}(&r)
	return 1
}
type T struct{
	a int
}
func (t T)Get() int{
	return t.a
}
func (t *T) Set(i int) {
	t.a = i
}
func (t *T)Print(){
	fmt.Printf("%p, %v, %d \n", t, t, t.a)
}
func main() {
	var t *T = &T{}
	t.Set(2)
	t.Get()
	(*t).Set(55)

	// 方法值 method value
	f := t.Set
	f(33)
	t.Print()

	fx := (*t).Set
	fx(45)
	t.Print()

	// 方法表达式
	// (T).Get(t)
	(T).Get(*t)
	m := T{a:3}
	// T.Get(&m)  cannot use &m (type *T) as type T in argument to T.Get
	(T).Get(m)
	// method expression 表达式调用 编译器不会进行自动转换
	// T.Set(*t, 44)  invalid method expression T.Set (needs pointer receiver: (*T).Set)
	ff := T.Get; ff(*t)
	ff = (T).Get; ff(*t)
	ffx := (*T).Set; ffx(t, 23)
	(*T).Set(t, 54)

	fmt.Println((*t).Get())
	fmt.Println("f0=", f0())
	fmt.Println("f1=", f1())
	fmt.Println("f2=", f2())
	fmt.Println("f3=", f3())
	fmt.Println("f4=", f4())
	reflectBasic()
	ReflectDyn()
	reflectValue()
	/*
	StructTest()
	StructTest2()
	StructTest3()
	StructTest4()
	 */
	//ReflectTest()
}

func StructTest(){
	u1 := UserInfo{Name: "q1mi", Age: 18}
	b, _ := json.Marshal(&u1)
	var m map[string]interface{}
	/*
	To unmarshal JSON into an interface value, Unmarshal stores one of these in the interface value:
	 bool, for JSON booleans
	 float64, for JSON numbers
	 string, for JSON strings
	 []interface{}, for JSON arrays
	 map[string]interface{}, for JSON objects
	 nil for JSON null
	 */
	_ = json.Unmarshal(b, &m)
	var u2 UserInfo
	// Unmarshal 在原路转struct时，int类型没有改变
	_ = json.Unmarshal(b, &u2)
	fmt.Printf("u2=>%#v  %T\n", u2, u2.Age)
	for k, v := range m{
		// struct 转 map[string]interface{} 过程中 value的类型由int变成了float64
		fmt.Printf("key: %v value(%T): %v\n", k, v, v)
	}
}

func StructTest2(){
	u1 := UserInfo{Name: "q1mi", Age: 18}
	m1, _ := ToMap(&u1, "json") // 结构体定义指定的tag为json
	for k, v := range m1{
		fmt.Printf("key: %v value(%T): %v\n", k, v, v)
	}
}
//使用反射遍历结构体字段的方式生成map
func ToMap(in interface{}, tagName string) (map[string]interface{}, error){
	out := make(map[string]interface{})
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()  // in 是结构体指针
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accept struct/struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历struct字段，指定tagName值为map中key；字段值为map中value
	for i := 0; i < v.NumField(); i++{
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out,nil
}
// 通过第三方库 structs， 注意它使用的自定义结构体tag是 structs
// 因此 UserInfo结构体需要加 structs tag
func StructTest3() {
	// UserInfo 用户信息，加载额外的structs
	type UserInfo struct {
		Name string `json:"name" structs:"name"`
		Age  int    `json:"age" structs:"age"`
	}
	u1 := UserInfo{Name: "q1mi", Age: 18}
	m1 := structs.Map(&u1)
	for k, v := range m1{
		fmt.Printf("key: %v value(%T): %v\n", k, v, v)
	}
}
// 嵌套情况
func StructTest4() {
	// Profile 配置信息
	type Profile struct {
		Hobby string `json:"hobby" structs:"hobby"`
	}
	// UserInfo 用户信息
	type UserInfo struct {
		Name string `json:"name" structs:"name"`
		Age  int    `json:"age" structs:"age"`
		Profile Profile	`json:"profile" structs:"profile"`
	}
	u1 := UserInfo{Name: "q1mi", Age: 18, Profile: Profile{"双色球"}}
	m1 := structs.Map(&u1)
	fmt.Printf("==> %v %T\n", m1, m1)
	for k, v := range m1{
		vv := reflect.ValueOf(v)
		fmt.Println(vv.Kind() )
		if vv.Kind() == reflect.Map{
			for ek,ev := range v.(map[string]interface{})  {
				fmt.Println(ek," ", ev)
			}
		}
		fmt.Printf("key: %v value(%T): %v -- %v -- %v\n", k, v, v, vv.Kind(), reflect.TypeOf(v))
	}
	PrintMapInterface(m1)
}
func PrintMapInterface(target map[string]interface{}){
	for k, v := range target{
		fmt.Printf("{%v:", k)
		switch v.(type) {
		case nil,string,int,int64,int32:
			fmt.Printf("%v}", v)
		case map[string]interface{}:
			PrintMapInterface(v.(map[string]interface{}))
			fmt.Printf("}")
		}
	}
}

// reflect 2个重要类型reflect.Type和reflect.Value
// reflect.Type是一个有很多方法的接口，表示Go语言的一个类型，携带的方法用来识别类型以及透视类型的组成部分（包含函数的各个参数）。
// reflect.Type接口只有一个实现，即类型描述符， 接口值中的动态类型也是类型描述符
// reflect.Value可以包含一个任意类型的值，reflect.ValueOf函数接收任意的interface{}并将接口的动态值以reflect.Value的形式返回。
// reflect.ValueOf的返回值也都是具体值, 不过reflect.Value也可以包含一个接口值。
func ReflectTest(){
	// reflect.TypeOf函数接受任何interface{}参数，并且把接口中的动态类型以reflect.Type形式返回。
	t := reflect.TypeOf(3) // Typeof返回一个接口值对应的动态类型，总返回具体类型（不会是接口类型）
	// 将具体类型值赋给一个接口类型时会发生一个隐式类型转换
	// 转换会生成一个包含两部分内容的接口值： 动态类型部分是操作数的类型(int), 动态值部分是操作数的值(3)
	fmt.Println(t.String())
	fmt.Printf("%T\n", 3) // 内部实现就使用了reflect.TypeOf
	fmt.Println(t)

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Printf("%#v\n", v)
	fmt.Println(v.String()) // reflect.Value也满足fmt.Stringer 除非Value包含的是一个字符串，否则String方法的结果仅仅暴露类型

	t = v.Type()
	fmt.Println(t.String())
	// reflect.ValueOf的逆操作是relect.Value.Interface方法。返回一个interface{}接口值，与reflect.Value包含同一个具体值
	x := v.Interface() // 返回interface{}空接口
	i := x.(int) // 类型断言
	fmt.Printf("%d\n", i)
	var y int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(Reflect.Any(x))
	fmt.Println(Reflect.Any(d))
	fmt.Println(Reflect.Any([]int64{y}))
	fmt.Println(Reflect.Any([]time.Duration{d}))

	ix := 1
	vx := reflect.ValueOf(&ix)
	var vxx reflect.Value = vx.Elem()
	vxx.SetInt(10)
	fmt.Println(vxx)

	Reflect.Print(time.Hour)
	Test()
	/*
		var testType reflect.Type
		var testValue reflect.Value
		Reflect.Print(testType)
		Reflect.Print(testValue)
	*/
}

type Student struct{
	Name	string "学生姓名"
	Age		int	`a:"1111"b:"3333"`
}

func reflectBasic(){
	s := Student{}
	rt := reflect.TypeOf(s)
	fieldName, ok := rt.FieldByName("Name") // StructField 返回值
	if ok {
		fmt.Println(fieldName)
		fmt.Printf("%#v\n", fieldName)
	}
	if fieldAge, ok := rt.FieldByName("Age"); ok {
		fmt.Printf("%#v\n", fieldAge)
		fmt.Println(fieldAge.Tag.Get("a")) // 可以像jsion一样取tag里的数据
	}
	fmt.Println("type name: ", rt.Name())// Student
	fmt.Println("type_NumField: ", rt.NumField()) // 2
	fmt.Println("type pkgPath: ", rt.PkgPath()) //main
	fmt.Println("type String: ", rt.String()) // main.Student
	fmt.Println("type.Kind.String: ", rt.Kind().String()) // struct
	fmt.Printf("type.Kind: %#v\n", rt.Kind()) //  0x19 struct
	/*
	type.Field[0].Name:="Name"
	type.Field[1].Name:="Age"
	 */
	for i := 0; i < rt.NumField(); i++{
		fmt.Printf("type.Field[%d].Name:=%#v \n", i, rt.Field(i).Name)
	}

	sc := make([]int, 10)
	sc = append(sc, 1, 2, 3, 4)
	sct := reflect.TypeOf(sc)

	// 获取slice元素的Type
	scet := sct.Elem()
	// slice element type.Kind()=int | 2 | int
	fmt.Printf("slice element type.Kind()=%v | %d | %v \n", scet.Kind(), scet.Kind(), scet.String())
	// slice element type.Name()=int - type.NumMethod()=0
	fmt.Printf("slice element type.Name()=%v - type.NumMethod()=%v\n", scet.Name(), scet.NumMethod())
	// slice element type.PkgPath()= |
	fmt.Printf("slice element type.PkgPath()=%v | %v\n", scet.PkgPath(), sct.PkgPath())
}

type INT int
type A struct {
	a int
}
type B struct {
	b string
}
type Ita interface {
	String() string
}
func (b B) String() string {
	return b.b
}
func ReflectDyn(){
	var a INT = 12
	var b int = 14
	// 实参是具体类型，reflect.TypeOf返回时其静态类型, 具体的类型信息
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)
	// ta=INT tb=int
	fmt.Printf("ta=%v tb=%v\n", ta.Name(), tb.Name())
	// 获得底层基础类型; ta.Kind = int  tb.Kind = int
	fmt.Printf("ta.Kind = %v  tb.Kind = %v \n", ta.Kind().String(), tb.Kind().String())
	// INT 与 int 是2个类型，两者不等
	if ta == tb {
		fmt.Println("ta==tb")
	}else {
		fmt.Println("ta != tb")
	}

	s1, s2 := A{1}, B{"tata"}
	// A  struct 实参是具体类型，reflect.TypeOf 返回的是其静态类型 Type的Kind方法返回的是其底层基础类型
	fmt.Println(reflect.TypeOf(s1).Name(), reflect.TypeOf(s1).Kind().String())
	// B  struct 实参是具体类型，reflect.TypeOf 返回的是其静态类型
	fmt.Println(reflect.TypeOf(s2).Name(), reflect.TypeOf(s1).Kind().String())

	ita := new(Ita) // Ita是interface, 但是ita是*Ita指针类型
	var itb Ita = s2
	var itc Ita
	// 实参是未绑定具体变量的接口类型，reflect.TypeOf返回的是接口类型本身，也就是接口的静态类型
	fmt.Println(reflect.TypeOf(ita).Elem().Name()) 				// Ita
	fmt.Println(reflect.TypeOf(ita).Elem().Kind().String())		// interface
	// TypeOf 参数 必须是个实例，不能是nil
	fmt.Printf("-->%T\n", itc) // return nil
	fmt.Println(reflect.TypeOf(itc)) // <nil>
	//fmt.Println(reflect.TypeOf(itc).Name()) // 错误 nil 空指针
	//fmt.Println(reflect.TypeOf(itc).Kind().String())
	// 实参是绑定了具体变量的接口类型，reflect.TypeOf 返回的是绑定的具体类型，即接口的动态类型
	fmt.Println(reflect.TypeOf(itb).Name())				// B
	fmt.Println(reflect.TypeOf(itb).Kind().String())	// struct

}

type User struct {
	Id		int
	Name	string
	Age		int
}
func (this User) String() {
	fmt.Println("User:", this.Id, this.Name, this.Age)
}
func Info(o interface{}){
	var to reflect.Type = reflect.TypeOf(o)
	// 获取Value
	var v reflect.Value = reflect.ValueOf(o)
	// 通过Value 获取Type
	var t reflect.Type = v.Type()
	// Type  User
	fmt.Println("Type ", t.Name())
	// 通过reflect.TypeOf(o) 和 reflect.ValueOf(o).Type() 等价的
	if to == t {
		fmt.Println(to, t.Name(), to.String())
	}
	// 访问接口字段名、字段类型和字段值信息
	for i := 0; i < to.NumField(); i++{
		var field reflect.StructField = to.Field(i) // 字段类型
		var Value reflect.Value = v.Field(i)
		var value interface{} = Value.Interface()
		// 类型查询
		switch value := value.(type) {
		case int:
			fmt.Printf(" %6s: %v = %d[%T]\n", field.Name, field.Type, value, value)
		case string:
			fmt.Printf(" %6s: %v = %s[%T]\n", field.Name, field.Type, value, value)
		default:
			fmt.Printf(" %6s: %v = %d[%T]\n", field.Name, field.Type, value,value)
		}
	}
}
func reflectValue(){
	u := User{1, "Tom", 30}
	Info(u)

	va := reflect.ValueOf(u)
	vb := reflect.ValueOf(&u)
	// 传入的是值，副本，因此原值是不可修改的
	fmt.Println(va.CanSet(), va.FieldByName("Name").CanSet()) //false false
	// 传入的指针 指针不可修改，是副本，但指针指向的对象是可以修改的
	fmt.Println(vb.CanSet(), vb.Elem().FieldByName("Name").CanSet()) // false, true

	fmt.Printf("%v \n", vb)
	name := "shlly"
	vc := reflect.ValueOf(name)
	// 通过set函数修改变量的值
	vb.Elem().FieldByName("Name").Set(vc)
	fmt.Println(vb)
}




// UserInfo 用户信息
type UserInfo struct {
	Name string `json:"name" structs:"name"`
	Age  int    `json:"age" structs:"age"`
	Profile `json:"profile" structs:"profile"`
}

// Profile 配置信息
type Profile struct {
	Hobby string `json:"hobby" structs:"hobby"`
}

func Test(){
	u1 := UserInfo{Name: "q1mi", Age: 18, Profile: Profile{"双色球"}}
	m3 := structs.Map(&u1)

	for k, v := range m3 {
		if o, ok := v.(map[string]interface{}); ok {
			fmt.Println(ok," ",o)
		}
		switch value := v.(type){
		case map[string]interface{}:
			for ek, ev := range value {
				fmt.Println(ek,ev)
			}
		}
		fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
	}
}










