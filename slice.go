package main

import "fmt"

func main() {
	 baseUsage()
	//appendUsage()

}

func appendUsage(){
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
	var nums []int  // nil slice
	printSlice(nums) // len=0 cap=0 slice=[]int(nil)

	nums = append(nums, 0)
	printSlice(nums)

	var nums2 = []int{}
	nums2 = append(nums2, 0,1)
	printSlice(nums2)

	var nums3 = make([]int, 0)
	nums = append(nums3, 1,2,3)
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

func printSlice(value interface{}){
	switch value.(type) {
	case []string:
		fmt.Printf("len=%d cap=%d slice=%#v\n", len(value.([]string)), cap(value.([]string)), value.([]string))
	case []int:
		x := value.([]int)
		fmt.Printf("len=%d cap=%d slice=%#v\n", len(value.([]int)), cap(x), x)
	}

}
