package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GCD using the Euclidean Algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	    }
	return a
}

// LCM using the formula: (a * b) / GCD(a, b)
func lcm(a, b int) int {
return (a * b) / gcd(a, b)
}

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func getTotalX(a []int, b []int) int {
    // find lcm of array a
	lcm_a := a[0]
	for _, num := range a[1:] {
		lcm_a = lcm(lcm_a, num)
	}

	// find gcd of array b
	gcd_b := b[0]
	for _, num := range b[1:] {
		gcd_b = gcd(gcd_b, num)
	}

	count := 0
	for multi := lcm_a; multi <= gcd_b; multi += lcm_a {
		if gcd_b % multi == 0 {
			count++
		}
	}

	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Line 1 has two numbers n and m
	scanner.Scan()	// this reads a line of text into buffer
	var n, m int
	fmt.Sscan(scanner.Text(), &n, &m)

	// Line 2 has n numbers
	scanner.Scan()	// this returns true or false: I ignored the error
	fields := strings.Fields(scanner.Text())
	a := make([]int, 0, n)	// preallocate slice for efficiency
	for i := 0; i < n && i < len(fields); i++ {
		val, _ := strconv.Atoi(fields[i])
		a = append(a, val)
	}

	// Line 3 has m numbers
	scanner.Scan()
	fields2 := strings.Fields(scanner.Text())
	b := make([]int, 0, m)	// preallocate slice for efficiency
	for i := 0; i < m && i < len(fields2); i++ {
		val, _ := strconv.Atoi(fields2[i])
		b = append(b, val)
	}

	total := getTotalX(a, b)
	fmt.Println(total)
}
