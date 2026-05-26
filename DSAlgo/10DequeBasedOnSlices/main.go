package main

// Run: go run main.go deque.go

import (
	"fmt"
)

func main() {
    var q Deque[string]
    q.PushBack("foo")
    q.PushBack("bar")
    q.PushBack("baz")

    fmt.Println(q.Len())   // Prints: 3
    fmt.Println(q.Front()) // Prints: foo
    fmt.Println(q.Back())  // Prints: baz

    q.PopFront() // remove "foo"
    q.PopBack()  // remove "baz"

    q.PushFront("hello")
    q.PushBack("world")

	fmt.Println()

    // Consume deque and print elements.
    for q.Len() != 0 {
        fmt.Println(q.PopFront())
    }
}