package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func angryProfessor(k int, a []int) string {
	var count int

	for _, v := range a {
		if v <= 0 {
			count++
		}
	}

	if (count >= k) {
		return "NO"
	}
	return "YES"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// first line contains t, the number of test cases
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	// go over each test case
	for range t {
		// first line of each test case contains n and k
		scanner.Scan()
		nk := strings.Fields(scanner.Text())
		n, _ := strconv.Atoi(nk[0])
		k, _ := strconv.Atoi(nk[1])

		// second line of each test case contains n arrival times of students
		scanner.Scan()
		ars := strings.Fields(scanner.Text())
		a := make([]int, n);
		for i := range n {
			a[i], _ = strconv.Atoi(ars[i])
		}

		fmt.Println(angryProfessor(k, a))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
