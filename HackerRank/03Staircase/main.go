package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := range n {

		// solution 1	
		// var line []byte
		// spaces := n - i - 1
		// sharps := i + 1

		// for range spaces {
		// 	line = append(line, ' ')
		// }
		// for range sharps {
		// 	line = append(line, '#')
		// }

		// fmt.Println(string(line))

		// solution 2
		// spaces := n - i - 1
		// sharps := i + 1

		// for j := range spaces {
		// 	fmt.Print(" ")
		// }
		// for j := range sharps {
		// 	fmt.Print("#")
		// }
		// fmt.Println()

		// solution 3
		fmt.Println(strings.Repeat(" ", n-i-1) + strings.Repeat("#", i+1))

	}
}