package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func Insert(head *listNode, value int) *listNode {
	newNode := &listNode{Val: value}

	if head == nil {
		return newNode
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	return head

}

func reverseList(head *listNode) *listNode {
	if head == nil {
		return nil
	}
	var curr, prev *listNode
	curr = head
	prev = nil
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
func print(head *listNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	head := &listNode{Val: 1}
  head.Next = &listNode{Val: 2}
  head.Next.Next = &listNode{Val: 3}
  fmt.Println("Original list: ")
  print(head)

  fmt.Println("Reersed list: ")
  print(reverseList(head))
}