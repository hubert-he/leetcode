package Tree

type TreeNode struct {
	Val interface{}
	Children []*TreeNode
}

func PrintTree(root *TreeNode, t int) []interface{} {
	switch t {
	case PreOrder:
		return preOrderNTree(root)
	case PostOrder:
		return postOrderNTree(root)
	case MidOrder:
		return nil
	case LayerOrder:
		return layerOrderNTree(root)
	default:
		return nil
	}
}

func preOrderNTree(root *TreeNode) []interface{} {
	serial := []interface{}{}
	if root == nil {
		return nil
	}
	serial = append(serial, root.Val)
	for _, item := range(root.Children) {
		serial = append(serial, preOrderNTree(item)...)
	}
	return serial
}

func PreOrderNTreeIter(root *TreeNode) []interface{} {
	serial := []interface{}{}
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		item := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		serial = append(serial, item.Val)
		for top := len(item.Children) - 1; top >= 0; top-- {
			stack = append(stack, item.Children[top])
		}
	}
	return serial
}

func postOrderNTree(root *TreeNode) []interface{} {
	serial := []interface{}{}
	if root == nil {
		return nil
	}
	for _, item := range(root.Children) {
		serial = append(serial, postOrderNTree(item)...)
	}
	serial = append(serial, root.Val)
	return nil
}

func postOrderNTreeIter(root *TreeNode) []interface{}{
	res := []interface{}{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res = append([]interface{}{node.Val}, res...)

		for _, v := range node.Children {
			stack = append(stack, v)
		}
	}
	return res
}
func layerOrderNTree(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}
	serial := []interface{}{}
	Queue := []interface{}{root}
	for len(Queue) > 0 {
		serial = append(serial, Queue[0])
		for _, i := range(root.Children){
			Queue = append(Queue, i)
		}
		Queue = Queue[1:]
	}
	return serial
}

func deserial(s string) *TreeNode {
	const (
		SPACE = iota
		LBRACKET
		RBRACKET
		NUM
	)
	const (
		START = iota

	)
	var root *TreeNode = nil
	return root
}
