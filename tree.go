package main

import (
	"./Tree"
	"fmt"
	"runtime"
)

func main() {
	/*
		var t = []int{2,1,3,nil,4,nil,7}
		var num = make([]interface{}, len(t))
		for c,_ := range(t) {
			num[c] = t[c]
		}
	*/
	runtime.GOMAXPROCS(1)
	var num = []interface{}{2, 1, 3, nil, 4, nil, 7, nil, nil, 5, 6}
	root := Tree.GenerateBiTree(num)
	fmt.Println("PreOrder: ", Tree.PrintBiTree(root, Tree.PreOrder))
	fmt.Println("PreOrder Iter: ", Tree.PrintBiTree(root, Tree.PreOrderIter))
	fmt.Println("MidOrder: ", Tree.PrintBiTree(root, Tree.MidOrder))
	fmt.Println("MidOder iter: ", Tree.PrintBiTree(root, Tree.MidOrderIter))
	fmt.Println("PostOredr: ", Tree.PrintBiTree(root, Tree.PostOrder))
	fmt.Println("PostOrder Iter: ", Tree.PrintBiTree(root, Tree.PostOrderIter))
	fmt.Println("PostOrder Iter reverse: ", Tree.PrintBiTree(root, Tree.PostOrderIterII))
	fmt.Println("PostOrder Iter IV: ", Tree.PrintBiTree(root, Tree.PostOrderIterIII))
	fmt.Println("LayerOrder: ", Tree.PrintBiTree(root, Tree.LayerOrder))

	fmt.Println("isSubTree-->")
}
