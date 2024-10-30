package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func newNode(value int) *TreeNode {
	return &TreeNode{value: value}
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.left == nil && root.right == nil {
		return 1
	}
	left := int(^uint(0) >> 1)
	right := int(^uint(0) >> 1)

	if root.left != nil {
		left = minDepth(root.left)
	}
	if root.right != nil {
		right = minDepth(root.right)
	}

	return min(left, right) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	root := newNode(1)
	root.left = newNode(2)
	root.right = newNode(3)
	root.left.left = newNode(4)
	root.left.right = newNode(5)

	fmt.Println("Minimum depth of the binary tree:", minDepth(root))
}
