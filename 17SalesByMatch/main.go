package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInputUnbuffered() (int, []int) {
	var n int
	fmt.Scanln(&n)

	var ar []int = make([]int, n)
	for i := range n {
		fmt.Scan(&ar[i])
	}

	return n, ar
}

func readInputScanner() (int, []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	var ar []int = make([]int, n)
	for i := range n {
		ar[i], _ = strconv.Atoi(fields[i])
	}

	return n, ar
}

func readInputReader() (int, []int) {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	line, _ = reader.ReadString('\n')
	fields := strings.Fields(line)
	var ar []int = make([]int, n)
	for i := range n {
		ar[i], _ = strconv.Atoi(fields[i])
	}

	return n, ar
}

func sockMerchant( /*n*/ _ int, ar []int) int {
	// Precondition: 1 <= ar[i] <= 100
	var count [101]int

	for _, v := range ar {
		count[v]++
	}

	var ans int
	for _, v := range count {
		ans += v / 2
	}

	return ans
}

func sockMerchant2( /*n*/ _ int, ar []int) int {
	// This solution uses a map to count the occurrences of each sock color.
	count := make(map[int]int)

	for _, v := range ar {
		count[v]++
	}

	var ans int
	for _, v := range count {
		ans += v / 2
	}

	return ans
}

func main() {
	// Read n on the first line
	//n, ar := readInputUnbuffered()
	//n, ar := readInputScanner()
	n, ar := readInputReader()

	//fmt.Println(sockMerchant(n, ar))
	fmt.Println(sockMerchant2(n, ar))
}
