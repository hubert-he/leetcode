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
	/*
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
	fmt.Println("Serialization: ", Tree.Serialization(root))
	fmt.Println("isSubTree-->")
	 */

	num = []interface{}{1,2,3,nil,nil,4,5}
	num = []interface{}{3,5,1,6,2,0,8,nil,nil,7,4}
	//num = []interface{}{5,1}
	//fmt.Println("Serialization: ", Tree.Serialization(Tree.GenerateBiTree(num)))
	tree := Tree.GenerateBiTree(num)
	ret := Tree.DistanceK4(tree, Tree.Find(tree, 5), 2)
	fmt.Println(ret)
	num = []interface{}{0,nil,1,nil,2,nil,3,4}
	tree = Tree.GenerateBiTree(num)
	fmt.Println(Tree.DistanceK4(tree, Tree.Find(tree, 2), 2))
}
