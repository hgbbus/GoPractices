package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	// given node is not the tail node
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	// test case 1
	node := &ListNode{Val: 5}
	node.Next = &ListNode{Val: 1}
	node.Next.Next = &ListNode{Val: 9}
	head := &ListNode{Val: 4, Next: node}
	deleteNode(node)
	fmt.Println(head.Val) // Output: 4
	fmt.Println(head.Next.Val) // Output: 1
	fmt.Println(head.Next.Next.Val) // Output: 9

	// test case 2
	head = &ListNode{Val: 4}
	head.Next = &ListNode{Val: 5}
	node = &ListNode{Val: 1}
	head.Next.Next = node
	node.Next = &ListNode{Val: 9}
	deleteNode(node)
	fmt.Println(head.Val) // Output: 4
	fmt.Println(head.Next.Val) // Output: 5
	fmt.Println(head.Next.Next.Val) // Output: 9
}
