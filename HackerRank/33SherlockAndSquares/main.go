package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'squares' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER a
 *  2. INTEGER b
 */

func squares(a int32, b int32) int32 {
    // Write your code here
	var sqa, sqb int32
	sqa = int32(math.Sqrt(float64(a)))
	sqb = int32(math.Sqrt(float64(b)))

	var result int32 = sqb - sqa
	if sqa*sqa == a {
		result++
	}
	return result
}

func main() {
	var t int32
	fmt.Scanln(&t)

	for range t {
		var a, b int32
		fmt.Scanln(&a, &b)
		fmt.Println(squares(a, b))
	}
}
