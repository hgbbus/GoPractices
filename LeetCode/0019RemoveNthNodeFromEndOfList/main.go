package main

import "fmt"
import "log"

type ListNode struct {
	Val int
	Next *ListNode
}

// Approach 1: Two-pass algorithm
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// figure out how many nodes are on the list
	sz := 0
	for curr := head; curr != nil; curr = curr.Next {
		sz++
	}

	if n > sz || sz == 0 {
		log.Fatal("Violation of the constraints!")
	}

	if n == sz {
		// removing head
		head = head.Next
	} else {
		n = sz - n
		prev := head
		for i := 0; i < n-1; i++ {
			prev = prev.Next
		}
		prev.Next = prev.Next.Next
	}

	return head
}

// Approach 2: One-pass algorithm
func removeNthFromEndOptimal(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy

	// move fast pointer n steps ahead
	for range n {
		fast = fast.Next
	}

	// move both pointers until fast reaches the end
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// remove the nth node from the end
	slow.Next = slow.Next.Next

	return dummy.Next
}

func main() {
	// test case 1
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	fmt.Print("Original list: ")
	printList(head) // Output: 1 -> 2 -> 3 -> 4 -> 5
	n := 2
	newHead := removeNthFromEndOptimal(head, n)
	fmt.Print("Modified list: ")
	printList(newHead) // Output: 1 -> 2 -> 3 -> 5
	fmt.Println()

	// test case 2
	head = &ListNode{Val: 1}
	n = 1
	fmt.Print("Original list: ")
	printList(head) // Output: 1
	newHead = removeNthFromEndOptimal(head, n)
	fmt.Print("Modified list: ")
	printList(newHead) // Output: (empty list)
	fmt.Println()

	// test case 3
	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	fmt.Print("Original list: ")
	printList(head) // Output: 1 -> 2
	n = 1
	newHead = removeNthFromEndOptimal(head, n)
	fmt.Print("Modified list: ")
	printList(newHead) // Output: 1
	fmt.Println()

	// test case 4
	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	fmt.Print("Original list: ")
	printList(head) // Output: 1 -> 2
	n = 2
	newHead = removeNthFromEndOptimal(head, n)
	fmt.Print("Modified list: ")
	printList(newHead) // Output: 2
	fmt.Println()
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}
