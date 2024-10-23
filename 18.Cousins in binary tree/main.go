package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	q := []*TreeNode{root}
	root.Val = 0

	for len(q) > 0 {
		l := len(q)
		sum := 0

		for i := range l {
			cur := q[i]
			if cur.Left != nil {
				sum += cur.Left.Val
			}

			if cur.Right != nil {
				sum += cur.Right.Val
			}
		}

		for i := range l {
			cur := q[i]
			left, right := 0, 0
			if cur.Left != nil {
				left = cur.Left.Val
			}

			if cur.Right != nil {
				right = cur.Right.Val
			}
			if cur.Left != nil {
				cur.Left.Val = sum - left - right
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				cur.Right.Val = sum - left - right
				q = append(q, cur.Right)
			}
		}
		q = q[l:]
	}
	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			cur := q[i]
			fmt.Print(cur.Val, " ")
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		fmt.Println()
		q = q[l:]
	}
}

func main() {
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root.Left.Left = &TreeNode{4, nil, nil}
	root.Left.Right = &TreeNode{5, nil, nil}
	root.Right.Left = &TreeNode{6, nil, nil}
	root.Right.Right = &TreeNode{7, nil, nil}

	fmt.Println("Original Tree:")

	printTree(root)

	root = replaceValueInTree(root)

	fmt.Println("Modified Tree:")
	printTree(root)
}
