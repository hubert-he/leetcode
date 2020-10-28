package Tree

import (
	"fmt"
	"time"
)

type BiTreeNode struct {
	Val interface{}
	Left *BiTreeNode
	Right *BiTreeNode
}

const (
	PreOrder = iota
	PreOrderIter
	MidOrder
	PostOrder
	PostOrderIter
	LayerOrder
)

func GenerateBiTree(values []interface{}) (*BiTreeNode){
	var cursor, root *BiTreeNode
	if len(values) == 0{
		return nil
	}
	var Queue = []*BiTreeNode{}
	root = &BiTreeNode{values[0], nil, nil}
	Queue = append(Queue, root)
	index := 1
	for len(Queue) != 0 {
		var left, right interface{}
		if index < len(values) {
			left = values[index]
			index += 1
		}
		if index < len(values) {
			right = values[index]
			index += 1
		}
		cursor = Queue[0]
		Queue = Queue[1:]
		if left != nil {
			cursor.Left = &BiTreeNode{left, nil, nil}
			Queue = append(Queue, cursor.Left)
		}
		if right != nil {
			cursor.Right = &BiTreeNode{right, nil, nil}
			Queue = append(Queue, cursor.Right)
		}

	}
	return root
}

func PrintBiTree(root *BiTreeNode, t int) []interface{} {
	switch t {
	case PreOrder:
		return preOrder(root)
	case PreOrderIter:
		return preOrderIter(root)
	case PostOrder:
		return postOrder(root)
	case PostOrderIter:
		return postOrderIter(root)
	case MidOrder:
		return midOrder(root)
	case LayerOrder:
		return layerOrderDFS(root)
	default:
		return nil
	}
}

func preOrder(root *BiTreeNode) []interface{} {
	var serial = []interface{}{}
	if root == nil {
		return serial
	}
	serial = append(serial, root.Val)
	serial = append(serial, preOrder(root.Left)...)
	serial = append(serial, preOrder(root.Right)...)
	return serial
}

func preOrderIter(root *BiTreeNode) []interface{} {
	var serial = []interface{}{}
	var stack = []*BiTreeNode{}
	if root == nil {
		return serial
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		// pop out
		item := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		serial = append(serial, item.Val)
		// push in
		if item.Right != nil {
			stack = append(stack, item.Right)
		}
		if item.Left != nil {
			stack = append(stack, item.Left)
		}
	}
	return serial
}

func midOrder(root *BiTreeNode) []interface{} {
	var serial = []interface{}{}
	if root == nil {
		return serial
	}
	serial = append(serial, midOrder(root.Left)...)
	serial = append(serial, root.Val)
	serial = append(serial, midOrder(root.Right)...)
	return serial
}
func postOrder(root *BiTreeNode) []interface{} {
	var serial = []interface{}{}
	if root == nil {
		return serial
	}
	serial = append(serial, postOrder(root.Left)...)
	serial = append(serial, postOrder(root.Right)...)
	serial = append(serial, root.Val)
	return serial
}
func postOrderIter(root *BiTreeNode) []interface{} {
	var prev *BiTreeNode
	serial := []interface{}{}
	stack := []*BiTreeNode{root}
	for len(stack) > 0 {
		curr := stack[len(stack) - 1]
		if curr.Right == nil && curr.Left == nil {
			serial = append(serial, curr.Val)
			stack = stack[:len(stack)-1]
			prev = curr
		} else if prev != nil && (prev == curr.Left || prev == curr.Right) {//向上回溯
			serial = append(serial, curr.Val)
			stack = stack[:len(stack)-1]
			prev = curr
		} else {
			// 必须先右后左, 匹配栈的先进先出
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
			if curr.Left != nil {
				stack = append(stack, curr.Left)
			}
		}
	}
	return serial
}

func postOrderIterII(root *BiTreeNode) []interface{} {
	var serial = []interface{}{}
	if root == nil {
		return serial
	}
	stack := []*BiTreeNode{root}
	for len(stack) > 0 {
		// 获取栈顶元素
		top := stack[len(stack) - 1]
		// 遍历到最左边
		for top.Left != nil {
			stack = append(stack, top.Left)
			top = top.Left
		}
		for top.Right != nil {
			stack = append(stack, top.Right)
			top = top.Right
		}
		//拿出元素
		serial = append(serial, top.Val)
		stack = stack[:len(stack) - 1]

	}
	return serial
}

func layerOrder(root *BiTreeNode) []interface{} {
	t1 := time.Now() // get current time
	var serial = []interface{}{} // []interface{}类型 是一个切片，切片元素的类型恰好是interface{}
	// var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
	if root == nil {
		return serial
	}
	var Queue = []*BiTreeNode{}
	Queue = append(Queue, root)
	for len(Queue) != 0 {
		if Queue[0] != nil {
			serial = append(serial, Queue[0].Val)
			Queue = append(Queue, Queue[0].Left)
			Queue = append(Queue, Queue[0].Right)
		} else {
			serial = append(serial, nil)
		}
		Queue = Queue[1:]
	}
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
	return serial
}

func layerOrder2(root *BiTreeNode) []interface{} {
	t1 := time.Now() // get current time
	var serial = []interface{}{} // []interface{}类型 是一个切片，切片元素的类型恰好是interface{}
	// var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
	if root == nil {
		return serial
	}
	var Queue = []*BiTreeNode{}
	//var Queue = make([]*BiTreeNode, 4,16)
	Queue = append(Queue, root)
	for len(Queue) != 0 {
		size := len(Queue)
		for i := 0; i < size; i++{
			if Queue[i] != nil {
				serial = append(serial, Queue[i].Val)
				Queue = append(Queue, Queue[i].Left)
				Queue = append(Queue, Queue[i].Right)
			} else {
				serial = append(serial, nil)
			}
		}
		/*
		for _, item := range(Queue){
			if item != nil {
				serial = append(serial, item.Val)
				Queue = append(Queue, item.Left)
				Queue = append(Queue, item.Right)
			} else {
				serial = append(serial, nil)
			}
		}*/
		Queue = Queue[size:]
	}
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
	return serial
}

func layerOrderDFS(root *BiTreeNode) []interface{} {
	list := [][]interface{}{}

	layerDFS(&list, root, 0)
	serial := []interface{}{}
	// 转换二维slice为一维
	for _, item := range(list) {
		serial = append(serial, item...) // 通过初始化转换
	}
	return serial
}

func layerDFS(list *[][]interface{}, root *BiTreeNode, height int){
	if root == nil {
		return
	}
	if height >= len(*list) {
		// new slice if nil
		*list = append(*list, []interface{}{root.Val})
	} else {
		(*list)[height] = append((*list)[height], root.Val)
	}
	layerDFS(list, root.Left, height+1)
	layerDFS(list, root.Right, height+1)
}
