package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'workbook' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER_ARRAY arr
 */

func workbook(n int32, k int32, arr []int32) int32 {
    var result int32
	var page int32

	for i := range n {
		for j := int32(1); j <= arr[i]; j += k {
			page++
			if j <= page && page <= min(j + k - 1, arr[i]) {
				result++
			}
		}
	}

	return result
}

func main() {
	// Create a buffered reader to read input from standard input
	reader := bufio.NewReader(os.Stdin)
	
	// Read the first line containing n and k
	firstLine, _ := reader.ReadString('\n')
	parts := strings.Fields(firstLine)
	n, _ := strconv.Atoi(parts[0])
	k, _ := strconv.Atoi(parts[1])

	// Read the second line containing the number of problems in each chapter
	arrLine, _ := reader.ReadString('\n')
	arrParts := strings.Fields(arrLine)
	arr := make([]int32, n)
	for i := range n {
		a, _ := strconv.Atoi(arrParts[i])
		arr[i] = int32(a)
	}

	// Call the workbook function and print the result
	result := workbook(int32(n), int32(k), arr)
	fmt.Println(result)
}