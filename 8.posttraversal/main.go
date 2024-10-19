package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func posrOrder(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}

	result := []int{}

	result = append(result, posrOrder(node.left)...)
	result = append(result, posrOrder(node.right)...)
	result = append(result, node.value)

	return result
}

func main() {
	root := &TreeNode{
		value: 1,
		right: &TreeNode{
			value: 2,
			left: &TreeNode{
				value: 3,
			},
		},
	}

	fmt.Println(posrOrder(root))
}
