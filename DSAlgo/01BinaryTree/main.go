package main

import "fmt"

// first define a tree node type using struct
type TreeNode struct {
	Value		int
	Left		*TreeNode
	Right		*TreeNode
}

func main() {
	// Now, manually build a binary tree
	//
    //        10
    //       /  \
    //      5    15
    //     / \
    //    3   7
    //

	// Start with root pointer, which points to the root node
	root := &TreeNode{Value: 10}

	// Add two children for the root node
	root.Left = &TreeNode{Value: 5}
	root.Right = &TreeNode{Value: 15}

	// Add two children for the left child of the root
	root.Left.Left = &TreeNode{Value: 3}
	root.Left.Right = &TreeNode{Value: 7}

	// Let's print a few values in different nodes
	fmt.Println(root.Value)
	fmt.Println(root.Right.Value)
	fmt.Println(root.Left.Right.Value)
}