package Tree

import (
	"testing"
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
/*
func (this *Int) Swap(target interface{}) {
	*this, target = target.(Int), *this
}
*/

func generateBiTree(values []interface{}) *avlTreeNode {
	var cursor, root *avlTreeNode
	if len(values) == 0 {
		return nil
	}
	var Queue = []*avlTreeNode{}
	root = newNode(values[0])
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
			cursor.Left = newNode(left)
			Queue = append(Queue, cursor.Left)
		}
		if right != nil {
			cursor.Right = newNode(right)
			Queue = append(Queue, cursor.Right)
		}

	}
	return root
}

func TestAVLTree_IsBalanced(t *testing.T) {
	for _, test := range []struct{
		num		[]interface{}
		want 	bool
	}{
		{[]interface{}{}, true},
		{[]interface{}{Int(1)}, true},
		{[]interface{}{Int(1), Int(2)}, true},
		{[]interface{}{Int(81), nil, Int(887), Int(847), nil, nil, nil}, false},
		{[]interface{}{Int(1), Int(2), Int(3)}, true},
		{[]interface{}{Int(2), Int(1), Int(3), nil, Int(4), nil, Int(7), nil, nil, Int(5), Int(6)}, false},
		{[]interface{}{Int(318), Int(81), Int(887), Int(59), Int(300), Int(456), Int(887), nil, nil, nil, nil, Int(425), Int(540), nil, nil}, true},
	}{
		root := generateBiTree(test.num)
		_, is_balanced := isBalanced(root)
		if is_balanced != test.want{
			t.Errorf("%#v => balance = %t", test.num, !test.want)
		}
	}
}

func TestAVLTree_Search(t *testing.T) {
	var tree *AVLTree
	for _, test := range []struct{
		num		[]int
		key		int
		height	int
		want	interface{}
	}{
		{[]int{}, 45, 0, nil},
		{[]int{1}, 1, 1, 1},
		{[]int{1, 2, 3}, 1, 2, 1},
		{[]int{1, 2, 3}, 3, 2, 3},
		{[]int{1, 2, 3}, 11, 2, nil},
		{[]int{81, 887,847, 59, 81, 318, 425, 540, 456, 300}, 318, 4, 318},
	}{
		for _, value := range test.num {
			key := (Int)(value)
			if tree == nil {
				tree = NewAVLTree(key)
			} else {
				tree.Insert(key, 0)
			}
		}
		if tree.IsBalanced() == false{
			t.Errorf("%v create avl failed", test.num)
		}
		target := tree.Search(Int(test.key))
		if target == nil {
			if test.want != nil{
				t.Errorf("%d | %d search failed in %v", test.key, test.want, test.num)
			}
		}
	}
}

func TestAVLTree_Remove(t *testing.T) {
	var tree *AVLTree
	for _, test := range []struct{
		num		[]int
		key		int
		height	int
		want	interface{}
	}{
		{[]int{81, 887,847, 59, 81, 318, 425, 540, 456, 300}, 318, 4, 318},
	}{
		for _, value := range test.num {
			key := (Int)(value)
			if tree == nil {
				tree = NewAVLTree(key)
			} else {
				tree.Insert(key, 0)
			}
		}
		target := tree.Remove(Int(test.key))
		if target == nil {
			if test.want != nil{
				t.Errorf("%d | %d remove failed in %v", test.key, test.want, test.num)
			}
		}else{
			if int(target.Key.(Int)) != test.want.(int){
				t.Errorf("%d | %d remove failed in %v", target.Key, test.want, test.num)
			}
			if tree.Search(Int(test.key)) != nil{
				t.Errorf("%d remove failed in %v", test.key, test.num)
			}
		}
	}
}
