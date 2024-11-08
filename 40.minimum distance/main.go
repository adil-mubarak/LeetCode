package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	minDis := math.MaxInt
	prev := -1

	var InOrder func(tree *TreeNode)
	InOrder = func(tree *TreeNode) {
		if tree == nil {
			return
		}
		InOrder(tree.Left)

		if prev != -1 {
			diff := int(math.Abs(float64(tree.Val - prev)))
			if minDis > diff {
				minDis = diff
			}
		}
		prev = tree.Val

		InOrder(tree.Right)
	}
	InOrder(root)
	return minDis
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func main() {
	root := NewTreeNode(27)
	root.Right = NewTreeNode(34)
	root.Right.Right = NewTreeNode(58)
	root.Right.Right.Left = NewTreeNode(50)
	root.Right.Right.Left.Left = NewTreeNode(44)
	result := minDiffInBST(root)

	fmt.Println("Minimum difference between any two nodes:", result)
}
