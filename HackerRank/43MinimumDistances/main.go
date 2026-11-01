package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'minimumDistances' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func minimumDistances(a []int32) int32 {
    // Write your code here
    lastSeen := make(map[int32]int)
    minDist := int32(-1)
    for i, v := range a {
        if j, ok := lastSeen[v]; ok {
            dist := int32(i - j)
            if minDist == -1 || dist < minDist {
                minDist = dist
            }
        }
        lastSeen[v] = i
    }
    return minDist
}

/* Brute force (O(n^2)) solution
func minimumDistances(a []int32) int32 {
    // Write your code here
    minDist := int32(-1)
    for i := range a {
        for j := i + 1; j < len(a); j++ {
            if a[i] == a[j] {
                dist := int32(j - i)
                if minDist == -1 || dist < minDist {
                    minDist = dist
                }
            }
        }
    }
    return minDist
}
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// first line one number n
	n, _ := strconv.Atoi(scanner.Text())

	// second line n numbers
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	a := make([]int32, n)
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(fields[i])
		a[i] = int32(num)
	}

	result := minimumDistances(a)
	fmt.Println(result)
}