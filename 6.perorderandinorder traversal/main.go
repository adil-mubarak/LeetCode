package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	var rootIndex int
	for i, val := range inorder {
		if val == rootVal {
			rootIndex = i
			break
		}
	}

	leftInorder := inorder[:rootIndex]
	rightInorder := inorder[rootIndex+1:]

	leftPreorder := preorder[1 : 1+len(leftInorder)]
	rightPreorder := preorder[1+len(leftInorder):]

	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)

	return root
}

func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	fmt.Printf("%d ", root.Val)
	inorderTraversal(root.Right)
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree(preorder, inorder)

	fmt.Println("Inorder traversal of the constructed tree:")
	inorderTraversal(root)
}