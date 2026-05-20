package main

import (
	"fmt"
)

/*
 * Complete the 'bonAppetit' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY bill
 *  2. INTEGER k
 *  3. INTEGER b
 */
func bonAppetit(bill []int, k int, b int) {
	var sum int
	for _, v := range bill {
		sum += v
	}
	share := sum - bill[k]
	if share/2 == b {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - share/2)
	}
}

func main() {
	// first line contains n and k
	var n, k int
	fmt.Scanln(&n, &k)

	// second line contains n space separated numbers
	var bill []int = make([]int, n)
	for i := range n {
		fmt.Scan(&bill[i])
	}

	// third line contains b
	var b int
	fmt.Scanln(&b)

	bonAppetit(bill, k, b)
}