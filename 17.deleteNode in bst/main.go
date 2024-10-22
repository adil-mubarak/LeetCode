package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DeleteNode(node *TreeNode, key int) *TreeNode {
	if node == nil {
		return nil
	}

	if key < node.Val {
		node.Left = DeleteNode(node.Left, key)
	} else if key > node.Val {
		node.Right = DeleteNode(node.Right, key)
	} else {
		if node.Left == nil && node.Right == nil {
			return nil
		} else if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		} else {
			node.Val = FindMin(node.Right)
			node.Right = DeleteNode(node.Right, node.Val)
		}
	}
	return node
}

func FindMin(node *TreeNode) int {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current.Val
}
func insertNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: key}
	}

	if key < root.Val {
		root.Left = insertNode(root.Left, key)
	} else {
		root.Right = insertNode(root.Right, key)
	}
	return root
}
func inorderTraversal(root *TreeNode) {
	if root != nil {
		inorderTraversal(root.Left)
		fmt.Printf("%d ", root.Val)
		inorderTraversal(root.Right)
	}
}

func main() {
	root := &TreeNode{Val: 50}
	root = insertNode(root, 30)
	root = insertNode(root, 20)
	root = insertNode(root, 40)
	root = insertNode(root, 70)
	root = insertNode(root, 60)
	root = insertNode(root, 80)

	fmt.Println("In-order traversal of the original tree:")
	inorderTraversal(root)
	fmt.Println()

	root = DeleteNode(root, 50)

	fmt.Println("In-order traversal after deleting node 50:")
	inorderTraversal(root)
	fmt.Println()
}
