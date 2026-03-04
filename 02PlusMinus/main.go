package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)

	var a []int = make([]int, n)
	for i := range n {
		fmt.Scan(&a[i])
		//fmt.Println(a[i])
	}

	pnum := 0
	mnum := 0
	zeros := 0
	for _, x := range a {
		if x > 0 {
			pnum++
		} else if x < 0 {
			mnum++
		} else {
			zeros++
		}
	}

	// golang does not support implicit conversion
	fmt.Printf("%.6f\n", float64(pnum)/float64(n))
	fmt.Printf("%.6f\n", float64(mnum)/float64(n))
	fmt.Printf("%.6f\n", float64(zeros)/float64(n))
}
