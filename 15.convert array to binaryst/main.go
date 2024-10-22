package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func sortedArrayToBST(nums []int)*TreeNode{
	if nums == nil{
		return nil
	}

	return BuildBST(nums,0,len(nums)-1)
}
func BuildBST(nums []int,start, end int)*TreeNode{
	if start > end{
		return nil
	}

	mid := (start + end)/2

	node := &TreeNode{value: nums[mid]}

	node.left = BuildBST(nums,start,mid-1)
	node.right = BuildBST(nums,mid+1,end)
	return node
}

func InOrder(node *TreeNode) {
	if node == nil {
		return 
	}
	InOrder(node.left)
	fmt.Print(node.value," ")
	InOrder(node.right)

}

func main() {
	nums := []int{-10, -3, 0, 5, 9}

	root := sortedArrayToBST(nums)
	InOrder(root)
}