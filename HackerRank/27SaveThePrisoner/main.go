package main

import "fmt"

func saveThePrisoner(n int, m int, s int) int {
	return (m+s-2)%n + 1
}

func main() {
	// first line contains t, number of test cases
	var t int
	_, err := fmt.Scanln(&t)
	if err != nil {
		panic(err)
	}

	// each test case contains n, m, and s on a separate line
	for range t {
		var n, m, s int
		_, err := fmt.Scanln(&n, &m, &s)
		if err != nil {
			panic(err)
		}
		fmt.Println(saveThePrisoner(n, m, s))
	}
}
