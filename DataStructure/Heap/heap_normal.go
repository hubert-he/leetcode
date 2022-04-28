package Heap

import "fmt"

/* 这一部分是 heap 的原始实现
** heap 分为 大根堆 和 小根堆
** 下面给出heap的关键点，辅助快速回忆起构建一个heap关键节点
** 1. 堆是一个完全二叉树， 因此可以使用 连续数组作为 其存储结构
** 2. 堆树 与 数组的映射关系：
	parent[i] 		= floor{ (i-1)/2 }
	leftChild[i]	= 2 * i + 1
	rightChild[i]	= 2 * i + 2
** 3. 最后一个非叶子节点的下标映射为:
		n / 2 - 1 其中 n 为 数组大小
		还有个思路，知道 最后一个位置 为 n-1
		他的父节点就是 最后一个非叶子节点的下标： (n-1-1)/2 = n/2 - 1
** 4. 什么是 heapify, 需要仔细 学习 heapify  与 堆构建的关系
** 5. 构建一个堆
 */

func heapify(arr []int, i, n int){
	if i >= n{
		return
	}
	// 自上而下调整
	left  := 2 * i + 1
	right := 2 * i + 2
	t := i
	if left < n && arr[t] < arr[left]{
		t = left
	}
	if right < n && arr[t] < arr[right]{
		t = right
	}
	if t != i{
		arr[t], arr[i] = arr[i], arr[t]
		heapify(arr, t, n)
	}
}

func BuildHeap(arr []int)[]int{
	// 从最后一个非叶子节点开始 自下而上，进行heapify， 逐步保证构成堆
	// 因此构建的复杂度在 O(nlogn)
	n := len(arr)
	/* 对最后一个非叶子节点的理解
	last_node := n - 1
	last_node_parent := (last_node - 1) / 2
	for i := last_node_parent; i >= 0; i--{
		heapify(arr, i, n)
	}
	 */
	// 下面是代码优化版本的
	for i := n/2-1; i >= 0; i--{
		heapify(arr, i, n)
	}
	return arr
}

func Pop(hp *[]int, i int)int{
	 heap := *hp
	 fmt.Println("Pop: ", i)
	 heap[0], heap[i] = heap[i], heap[0]
	 heapify(heap, 0, i-1)
	 return heap[i]
}

func heap_sort(arr []int)[]int{
	n := len(arr)
	tree := BuildHeap(arr) // 构建大根堆
	// 连续pop后，堆消失，数组被排序
	for i := n-1; i > 0; i--{
		Pop(&tree, i)
	}
	return tree
}

func heap_sort_2(arr []int) []int{
	n := len(arr)
	tree := BuildHeap(arr)
	for i := n-1; i > 0; i--{
		tree[i], tree[0] = tree[0], tree[i]
		heapify(tree, 0, i)
	}
}