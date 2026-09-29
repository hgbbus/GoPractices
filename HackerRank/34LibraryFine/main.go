package main

import "fmt"

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
    // Write your code here
	if y1 > y2 {
		return 10000
	} else if y1 < y2 {
		return 0
	} else {
		if m1 > m2 {
			return (m1 - m2) * 500
		} else if m1 < m2 {
			return 0
		} else {
			if d1 > d2 {
				return (d1 - d2) * 15
			} else {
				return 0
			}
		}
	}
}

func main() {
	var d1, m1, y1 int32
	fmt.Scanf("%d %d %d", &d1, &m1, &y1)

	var d2, m2, y2 int32
	fmt.Scanf("%d %d %d", &d2, &m2, &y2)

	fmt.Println(libraryFine(d1, m1, y1, d2, m2, y2))
}
