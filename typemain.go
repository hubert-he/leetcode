package main

import (
	"./Reflect"
	"fmt"
	"reflect"
	"time"
)
// reflect 2个重要类型reflect.Type和reflect.Value
// reflect.Type是一个有很多方法的接口，表示Go语言的一个类型，携带的方法用来识别类型以及透视类型的组成部分（包含函数的各个参数）。
// reflect.Type接口只有一个实现，即类型描述符， 接口值中的动态类型也是类型描述符
// reflect.Value可以包含一个任意类型的值，reflect.ValueOf函数接收任意的interface{}并将接口的动态值以reflect.Value的形式返回。
// reflect.ValueOf的返回值也都是具体值, 不过reflect.Value也可以包含一个接口值。
func main() {
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
/*
	var testType reflect.Type
	var testValue reflect.Value
	Reflect.Print(testType)
	Reflect.Print(testValue)
 */
}










