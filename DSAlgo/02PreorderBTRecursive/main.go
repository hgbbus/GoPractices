package main

import "fmt"

type Node struct {
	Val			int
	Left		*Node
	Right		*Node
}

func preorder(subtree *Node) {
	if subtree == nil {
		return
	}
	
	fmt.Println("Val: ", subtree.Val)
	preorder(subtree.Left)
	preorder(subtree.Right)
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