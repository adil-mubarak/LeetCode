package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var count int
	var result int

	var InOrder func(node *TreeNode)
	InOrder = func(node *TreeNode) {
		if node == nil || count >= k {
			return
		}

		InOrder(node.Left)
		count++
		if count == k {
			result = node.Val
			return
		}

		InOrder(node.Right)
	}
	InOrder(root)
	return result
}

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Left.Left.Left = &TreeNode{Val: 1}

	k := 3
	fmt.Printf("The %d-th smallest element is: %d\n", k, kthSmallest(root, k))
}
