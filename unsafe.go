package main

import (
	"fmt"
	"unsafe"
)

var x struct {
	a bool
	b int16
	c []int
}

func main() {
	fmt.Println(unsafe.Sizeof(x)) // unsafe.Sizeof函数返回操作数在内存中的字节大小，参数可以是任意类型的表达式，但是它并不会对表达式进行求值。一个Sizeof函数调用是一个对应uintptr类型的常量表达式，因此返回的结果可以用作数组类型的长度大小，或者用作计算其他的常量
	fmt.Println(unsafe.Alignof(x)) // unsafe.Alignof 函数返回对应参数的类型需要对齐的倍数,通常情况下布尔和数字类型需要对齐到它们本身的大小(最多8个字节), 其它的类型对齐到机器字大小

	fmt.Println(unsafe.Sizeof(x.a))
	fmt.Println(unsafe.Alignof(x.a))
	fmt.Println(unsafe.Offsetof(x.a))// unsafe.Offsetof 函数的参数必须是一个字段 x.f, 然后返回 f 字段相对于 x 起始地址的偏移量, 包括可能的空洞
}