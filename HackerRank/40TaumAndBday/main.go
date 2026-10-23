package main

import (
	"fmt"
)

/*
 * Complete the 'taumBday' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER b
 *  2. INTEGER w
 *  3. INTEGER bc
 *  4. INTEGER wc
 *  5. INTEGER z
 */

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
    // Write your code here
	if bc > wc+z {
		bc = wc + z
	} else if wc > bc+z {
		wc = bc + z
	}
	return int64(b)*int64(bc) + int64(w)*int64(wc)
}

func main() {
	// read first line of input t
	var t int
	fmt.Scanln(&t)

	// read the next t lines of input
	for i := 0; i < t; i++ {
		var b, w, bc, wc, z int32
		fmt.Scanln(&b, &w)
		fmt.Scanln(&bc, &wc, &z)
		result := taumBday(b, w, bc, wc, z)
		fmt.Println(result)
	}
}
