package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func preorder(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}

	result := []int{}

	result = append(result, node.value)
	result = append(result, preorder(node.left)...)
	result = append(result, preorder(node.right)...)

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

	fmt.Println(preorder(root))
}
