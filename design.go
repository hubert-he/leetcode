package main

import (
	l295 "./design/Find-Median-from-Data-Stream-295"
	"fmt"
)
func main() {
	obj := l295.Constructor();
	obj.AddNum(1)
	obj.AddNum(2)
	fmt.Println(obj.FindMedian())
	obj.AddNum(3)
	fmt.Println(obj.FindMedian())

	obj2 := l295.Constructor();
	fmt.Println(obj2.FindMedian())
	obj2.AddNum(1513)
	fmt.Println(obj2.FindMedian())
	obj2.AddNum(5083)
	fmt.Println(obj2.FindMedian())
	obj2.AddNum(4386)
	fmt.Println(obj2.FindMedian())
	obj2.AddNum(2296)
	fmt.Println(obj2.FindMedian())
	obj2.AddNum(1370)
	fmt.Println(obj2.FindMedian())
}
