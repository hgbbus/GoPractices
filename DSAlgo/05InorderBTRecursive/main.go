package main

import "fmt"

// Define Node type for binary tree
type Node struct {
	Val		int
	Left	*Node
	Right	*Node
}

func inorder(subtree *Node) {
	if subtree == nil {
		return
	}

	inorder(subtree.Left)
	fmt.Println("Val:", subtree.Val)
	inorder(subtree.Right)
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

	root.Left = &Node{Val: 5, Left: &Node{Val: 3}}

	root.Right = &Node{Val: 15}
	root.Right.Left = &Node{Val: 7}
	root.Right.Right = &Node{Val: 9}

	inorder(root)
}