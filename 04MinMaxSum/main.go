package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func minMaxSum(arr []int) {
	max := arr[0]
	min := arr[0]
	sum := 0

	for _, x := range arr {
		sum += x
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}

	fmt.Println(sum-max, sum-min)
}

func main() {
	var nums []int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	for tok := range strings.FieldsSeq(line) {
		n, _ := strconv.Atoi(tok)
		nums = append(nums, n)
	}
	//fmt.Println(nums)
	minMaxSum(nums)
}
