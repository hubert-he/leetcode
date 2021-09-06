package Reflect

import (
	"fmt"
	"reflect"
)

/*判断反射值的空和有效性
func (Value) IsNil
func (v Value) IsNil() bool
IsNil报告v持有的值是否为nil，常用于判断指针是否为空
v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。
注意IsNil并不总是等价于go语言中值与nil的常规比较。
例如：如果v是通过使用某个值为nil的接口调用ValueOf函数创建的，v.IsNil()返回真，但是如果v是Value零值，会panic。
 */
//异常判断
func InterfaceIsNil1(i interface{}) bool {
	ret := i == nil
	if !ret { //需要进一步做判断
		defer func() {
			recover()
		}()
		ret = reflect.ValueOf(i).IsNil() //值类型做异常判断，会panic的
	}
	return ret
}

//类型判断
func InterfaceIsNil2(i interface{}) bool {
	ret := i == nil  // interface 的== 只能判断 动态值与 动态类型均为nil的情况为nil
	if !ret { //需要进一步做判断
		vi := reflect.ValueOf(i)
		kind := reflect.ValueOf(i).Kind()
		fmt.Println(kind)
		if kind == reflect.Slice ||
			kind == reflect.Map ||
			kind == reflect.Chan ||
			kind == reflect.Interface ||
			kind == reflect.Func ||
			kind == reflect.Ptr {
			return vi.IsNil()
		}
	}
	return ret
}

