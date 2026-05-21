package main

import "fmt"

type Node struct {
	val		int
	left	*Node
	right	*Node
}

func inorder(root *Node) {
	stack := []*Node{}	// empty literal slice for empty stack
	current := root		// critical for inorder (not needed for preorder)

	for current != nil || len(stack) > 0 {
		// descend to the leftmost node
		for current != nil {
			stack = append(stack, current)
			current = current.left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println("Val:", current.val)

		current = current.right
	}
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

	inorder(root)
}