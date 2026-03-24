package main

import (
	"fmt"
)

/*
 * Complete the 'kangaroo' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER x1
 *  2. INTEGER v1
 *  3. INTEGER x2
 *  4. INTEGER v2
 */

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Given constraints: 0 <= x1 < x2 <= 10000
	//                    1 <= v1 <= 10000
	//					  1 <= v2 <= 10000
    // x1 + v1 * n = x2 + v2 * n
	// x1 - x2 = (v2 - v1) * n
	// n = (x1 - x2) / (v2 - v1) = (x2 - x1) / (v1 - v2)

	// since x2 > x1, v1 must be greater than v2
	if v1 <= v2 {
		return "NO"
	}

	if (x2 - x1) % (v1 - v2) == 0 {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	var x1, v1, x2, v2 int32

	fmt.Scanf("%d%d%d%d", &x1, &v1, &x2, &v2)

	fmt.Println(kangaroo(x1, v1, x2, v2))
}
