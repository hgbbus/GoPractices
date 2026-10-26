package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * Complete the 'kaprekarNumbers' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER p
 *  2. INTEGER q
 */

func kaprekarNumbers(p int32, q int32) {
	var result []int32
	for i := p; i <= q; i++ {
		if isKNum(int(i)) {
			result = append(result, i)
		}
	}

	if len(result) == 0 {
		fmt.Println("INVALID RANGE")
	} else {
		fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
	}
}

func isKNum(n int) bool {
	var result bool

	nsq := n * n
	s_nsq := strconv.Itoa(nsq)
	ls := s_nsq[:len(s_nsq)/2]
	rs := s_nsq[len(s_nsq)/2:]

	l_num, _ := strconv.Atoi(ls)
	r_num, _ := strconv.Atoi(rs)
	if n == l_num+r_num {
		result = true
	}

	return result
}

func main() {
	var p, q int32
	fmt.Scanln(&p)
	fmt.Scanln(&q)

	kaprekarNumbers(p, q)
}
