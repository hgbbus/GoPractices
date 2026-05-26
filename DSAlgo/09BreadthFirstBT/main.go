package main

import "fmt"

type Node struct {
	val		int
	left	*Node
	right	*Node
}

func dequeue(queue []*Node) ([]*Node, *Node) {
	element := queue[0]
	return queue[1:], element
}

func breadth_first(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		var node *Node
		queue, node = dequeue(queue)

		fmt.Println("Val:", node.val)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
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
    //               \
	//                11

	root := &Node{val: 10}

	root.left = &Node{val: 5, left: &Node{val: 3}}
	root.right = &Node{val: 15}
	root.right.left = &Node{val: 7}
	root.right.right = &Node{val: 9}
	root.right.right.right = &Node{val: 11}

	breadth_first(root)
}
