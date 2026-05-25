package main

import "fmt"

type Node struct {
	Val		int
	Left	*Node
	Right	*Node
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

	postorder(root)
}

func postorder(subtree *Node) {
	if subtree == nil {
		return
	}

	postorder(subtree.Left)
	postorder(subtree.Right)
	fmt.Println("Val:", subtree.Val)
}