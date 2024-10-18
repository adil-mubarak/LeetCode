package main

import "fmt"

type Tree struct {
	value int
	left  *Tree
	right *Tree
}

func maxDepth(root *Tree) int {
	if root == nil {
		return 0
	}

	var result int

	left := maxDepth(root.left)
	right := maxDepth(root.right)

	if left > right {
		result = left+1
	} else {
		result = right+1
	}
	return result
}

func main() {
	root := &Tree{
		value: 3,
		left: &Tree{
			value: 9,
		},
		right: &Tree{
			value: 20,
			left: &Tree{
				value: 15,
			},
			right: &Tree{
				value: 7,
			},
		},
	}

	fmt.Println(maxDepth(root))
}