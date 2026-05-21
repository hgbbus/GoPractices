package main

import "fmt"

type Node struct {
	Val			int
	Left		*Node
	Right		*Node
}

func preorder(root *Node) {
	// non-recursive version of preorder traversal using stack
	if root == nil {
		return
	}

	// create a stack and put the root into it
	stack := []*Node{root}	// slice literal
	for len(stack) > 0 {
		// pop the top node from the stack
		n := len(stack)
		node := stack[n-1]
		stack = stack[:n-1]

		// print the value of the node
		fmt.Println("Val: ", node.Val)

		// push the right child first to be processed later
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
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
    //

	root := &Node{Val: 10}
	root.Left = &Node{Val: 5, 
						Left: &Node{Val: 3}}
	root.Right = &Node{Val: 15, 
						Left: &Node{Val: 7},
						Right: &Node{Val: 9}}
	
	preorder(root)
}