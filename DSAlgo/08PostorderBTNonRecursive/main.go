package main

import "fmt"

type Node struct {
	val		int
	left	*Node
	right	*Node
}

func pop(stack []*Node) ([]*Node, *Node) {
	n := len(stack)
	if n == 0 {
		return stack, nil
	}

	node := stack[n-1]
	stack[n-1] = nil	// Optional

	return stack[:n-1], node
}

func main() {
	// Now, manually build a binary tree
	//
    //        10
    //       /  \
    //      5    15
    //     /    /  \
    //    3    7    9
    //               \
	//                11

	root := &Node{val: 10}

	root.left = &Node{val: 5, left: &Node{val: 3}}
	root.right = &Node{val: 15}
	root.right.left = &Node{val: 7}
	root.right.right = &Node{val: 9}
	root.right.right.right = &Node{val: 11}

	postorderOneStack(root)
}

// Approach 1: Using two stacks - essentially a reversed preorder traversal
func postorderTwoStacks(root *Node) {
	if root == nil {
		return
	}

	stack1 := []*Node{root}
	stack2 := []*Node{}

	for len(stack1) > 0 {
		var curr *Node
		stack1, curr = pop(stack1)
		stack2 = append(stack2, curr)

		if curr.left != nil {
			stack1 = append(stack1, curr.left)
		}
		if curr.right != nil {
			stack1 = append(stack1, curr.right)
		}
	}

	for len(stack2) > 0 {
		var curr *Node
		stack2, curr = pop(stack2)
		fmt.Println("Val:", curr.val)
	}
}

// Approach 2: Using one stack and a pointer to track the last visited node
func postorderOneStack(root *Node) {
	if root == nil {
		return
	}

	stack := []*Node{}
	var lastVisited *Node
	curr := root

	for curr != nil || len(stack) > 0 {
		// traverse down the left subtree
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		}

		// peek the top node from the stack
		n := len(stack)
		top := stack[n-1]

		// if the right child is nil or already visited, visit the node
		if top.right == nil || top.right == lastVisited {
			fmt.Println("Val:", top.val)
			lastVisited = top
			stack = stack[:n-1] // pop the node from the stack
		} else {
			// otherwise, traverse the right subtree
			curr = top.right
		}
	}
}