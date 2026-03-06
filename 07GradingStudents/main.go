package main

import (
	"fmt"
)

func gradingStudents(grades []int32) []int32 {
	for i, grade := range grades {
		if grade < 38 {
			continue
		}

		if grade/5 < (grade + 2)/5 {
			grades[i] = (grade + 2)/5 * 5
		}
	}

	return grades
}

func main() {
	var n int32
	fmt.Scanln(&n)

	grades := make([]int32, n)

	for i := 0; i < int(n); i++ {
		fmt.Scanln(&grades[i])
	}

	result := gradingStudents(grades)

	for i := 0; i < int(n); i++ {
		fmt.Println(result[i])
	}
}