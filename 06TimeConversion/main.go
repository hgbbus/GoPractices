package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	var hours int
	hours, _ = strconv.Atoi(s[:2])

	if hours == 12 {
		hours = 0
	}

	if s[8:] == "PM" {
		hours += 12
	}

	return fmt.Sprintf("%02d", hours) + s[2:8]
}

func main() {
	var s string

	fmt.Scanln(&s)
	//s = strings.TrimSpace(s)
	fmt.Println(timeConversion(strings.TrimSpace(s)))
}
