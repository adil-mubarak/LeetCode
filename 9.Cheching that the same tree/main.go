package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	tree1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	tree2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println("Are tree1 and tree2 the same?", isSameTree(tree1, tree2))

	tree3 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}

	tree4 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}

	fmt.Println("Are tree3 and tree4 the same?", isSameTree(tree3, tree4))
}
