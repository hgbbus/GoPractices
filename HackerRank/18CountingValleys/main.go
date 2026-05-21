package main

import "fmt"

func countingValleys(steps int, path string) int {
	altitude := 0
	valleys := 0

	for i := 0; i < steps; i++ {
		switch path[i] {
		case 'U':
			altitude++
		case 'D':
			if altitude == 0 {
				valleys++
			}
			altitude--
		}
	}

	return valleys
}

func main() {
	var steps int
	fmt.Scanln(&steps)

	var path string
	fmt.Scanln(&path)

	result := countingValleys(steps, path)
	fmt.Println(result)
}