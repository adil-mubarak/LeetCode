//binary tree inorder traversal
package main

import "fmt"

type Tree struct {
	value int
	left  *Tree
	right *Tree
}

func InorderTraversal(root *Tree) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}

	result = append(result, InorderTraversal(root.left)...)
	result = append(result, root.value)
	result = append(result, InorderTraversal(root.right)...)

	return result
}

func main() {

	root := &Tree{
		value: 1,
		right: &Tree{
			value: 2,
			left: &Tree{
				value: 3,
			},
		},
	}

	fmt.Println(InorderTraversal(root))
}
