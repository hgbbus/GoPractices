package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'beautifulTriplets' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER d
 *  2. INTEGER_ARRAY arr
 */

func beautifulTriplets(d int32, arr []int32) int32 {
	freq := make(map[int32]int32, len(arr))
	for _, value := range arr {
		freq[value]++
	}

	var count int32
	for value, occurrences := range freq {
		count += occurrences * freq[value+d] * freq[value+2*d]
	}
	// Missing key in Go map returns 0

	return count
}

/* Brute force (O(n^3)) solution
func beautifulTriplets(d int32, arr []int32) int32 {
	var count int32

	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			firstDifference := arr[j] - arr[i]
			if firstDifference > d {
				break
			}
			if firstDifference != d {
				continue
			}

			for k := j + 1; k < len(arr); k++ {
				secondDifference := arr[k] - arr[j]
				if secondDifference > d {
					break
				}
				if secondDifference == d {
					count++
				}
			}
		}
	}

	return count
}
*/

func main() {
	// use buffered input to handle large amount of inputs
	scanner := bufio.NewScanner(os.Stdin)

	// first line contains n and d
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(fields[0])
	d, _ := strconv.Atoi(fields[1])

	// second line contains n numbers for the array
	scanner.Scan()
	fields = strings.Fields(scanner.Text())
	var ar []int32 = make([]int32, n)
	for i := range n {
		a, _ := strconv.Atoi(fields[i])
		ar[i] = int32(a)
	}

	fmt.Println(beautifulTriplets(int32(d), ar))
}
