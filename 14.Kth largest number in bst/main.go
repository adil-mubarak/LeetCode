package main

import (
	"container/list"
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(node *TreeNode, k int) int64 {
	if node == nil {
		return -1
	}

	queue := list.New()
	queue.PushBack(node)

	var levelSums []int64

	for queue.Len() > 0 {
		levelSize := queue.Len()
		var levelSum int64 = 0

		for i := 0; i < levelSize; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			levelSum += int64(node.Val)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		levelSums = append(levelSums, levelSum)
	}

	sort.Slice(levelSums, func(i, j int) bool {
		return levelSums[i] > levelSums[j]
	})

	if k > len(levelSums) {
		return -1
	}

	return levelSums[k-1]

}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	k := 2
	result := kthLargestLevelSum(root, k)
	fmt.Println(result)
}
