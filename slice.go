package main

import "fmt"

const N = 2

func main() {
	//baseUsage()
	//appendUsage()
	removeUsage()
}

func appendUsage() {
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	slice0 := source[2:3] // cap: 5-2=3
	printSlice(slice0)
	printSlice(source)
	slice0 = append(slice0, "watermelon")
	printSlice(slice0)
	printSlice(source)
	slice := source[2:3:3] // 限制cap cap: 3-2=1
	printSlice(slice)
	printSlice(source)
	slice = append(slice, "kiwi")
	printSlice(slice)
	printSlice(source)
}
func baseUsage() {
	var nums []int   // nil slice
	printSlice(nums) // len=0 cap=0 slice=[]int(nil)

	nums = append(nums, 0)
	printSlice(nums)

	var nums2 = []int{}
	nums2 = append(nums2, 0, 1)
	printSlice(nums2)

	var nums3 = make([]int, 0)
	nums = append(nums3, 1, 2, 3)
	printSlice(nums)

	var nums4 = []int{}
	nums4 = make([]int, 1, 5)
	// x := nums4[3:] 索引必须在可用的范围内
	x := nums4[0:]
	printSlice(x)
	sum := copy(nums4, nums) // copy 函数拷贝的多少 取决于dest即第一个参数的可用长度的多少

	fmt.Println(x, sum)
	printSlice(nums4)

	sx := []interface{}{}
	sx = append(sx, nil)
	sx = append(sx, 99)
	fmt.Printf("\n-> %#v\n", sx[0])
}

// http://c.biancheng.net/view/30.html
func removeUsage() {
	// 从slice首位置删除元素
	a := []int{1, 2, 3}
	// 移动指针
	a = a[1:] // 删除第1个元素
	a = a[N:] // 删除第N个元素
	printSlice(a)
	// 使用append
	// 将后面的数据向开头移动，可以用 append 原地完成（所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）
	a = []int{33, 34, 35}
	a = append(a[:0], a[1:]...) // 删除首元素
	fmt.Println(a[:0])
	printSlice(a)
	a = append(a[:0], a[N:]...) // 删除开头N个元素
	// 用 copy() 函数来删除开头的元素
	a = []int{33, 34, 36, 37, 38, 39, 40}
	a = a[:copy(a, a[1:])]
	a = a[:copy(a, a[N:])]

	// 从中间位置删除
	// 对于删除中间的元素，需要对剩余的元素进行一次整体挪动，同样可以用 append 或 copy 原地完成
	a = []int{1, 2, 3, 4, 5, 6}
	i := 3
	a = append(a[:i], a[i+1:]...)  // 删除中间1个元素
	a = append(a[:i], a[i+N:]...)  // 删除中间N个元素
	a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
	a = a[:i+copy(a[i:], a[i+N:])] // 删除中间N个元素
}

func printSlice(value interface{}) {
	switch value.(type) {
	case []string:
		fmt.Printf("len=%d cap=%d slice=%#v\n", len(value.([]string)), cap(value.([]string)), value.([]string))
	case []int:
		x := value.([]int)
		fmt.Printf("len=%d cap=%d slice=%#v\n", len(value.([]int)), cap(x), x)
	}

}
