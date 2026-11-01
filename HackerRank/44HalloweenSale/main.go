package main

import (
	"fmt"
)

/*
 * Complete the 'howManyGames' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER p
 *  2. INTEGER d
 *  3. INTEGER m
 *  4. INTEGER s
 */

func howManyGames(p int32, d int32, m int32, s int32) int32 {
	P, D, M, S := int64(p), int64(d), int64(m), int64(s)

	// n games are bought at decreasing prices; after that every game costs m
	n := (P-M)/D + 1
	cost := func(j int64) int64 { return j*P - D*j*(j-1)/2 }

	// largest j in [0, n] with cost(j) <= S
	lo, hi := int64(0), n
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if cost(mid) <= S {
			lo = mid
		} else {
			hi = mid - 1
		}
	}

	if lo < n {
		return int32(lo)
	}
	return int32(n + (S-cost(n))/M)
}

func main() {
	var p, d, m, s int32
	fmt.Scanf("%d %d %d %d", &p, &d, &m, &s)

	result := howManyGames(p, d, m, s)

	fmt.Println(result)
}