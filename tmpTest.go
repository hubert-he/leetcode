package main

import (
	"./Tree"
	"fmt"
)

type Int int
func (this Int) Less(target interface{}) bool{
	if this < target.(Int) {
		return true
	}
	return false
}

func (this Int) Equal(i interface{}) bool{
	if int(i.(Int)) == int(this){
		return true
	}
	return false
}

func main() {
	t := []int{-10,-3,0,5,9}
	s := make([]interface{}, len(t))
	for i, v := range t {
		s[i] = v
	}
	b := Tree.NewBSTFromSortedList(s)
	Tree.PrintBiTree(b.GetRoot(), Tree.PreOrder)
	var a *Tree.AVLTree
	nums := []int{81, 887,847, 59, 81, 318, 425, 540, 456, 300}
	for _, value := range nums {
		key := (Int)(value)
		if a == nil {
			a = Tree.NewAVLTree(key)
		} else {
			a.Insert(key, 0)
		}
	}
	fmt.Println(a.IsBalanced())
	a.Remove(Int(59))
	fmt.Println(a.IsBalanced())
	a.Remove(Int(300))
	fmt.Println(a.IsBalanced())
	//a.Remove(Int(nums[5]))
	//a.Remove(Int(847))
	fmt.Println("")
	a.Search(Int(81))
}
