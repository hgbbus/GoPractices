package main

import (
    "fmt"
)

func isLeapYear(year int) bool {
    if year < 1918 {
        return year % 4 == 0
    } else if year > 1918 {
        return (year % 400 == 0) || (year % 4 == 0 && year % 100 != 0)
    } else {
        return false
    }
}

func dayOfProgrammer(year int) string {
    if year == 1918 {
        return "26.09.1918"
    } else if isLeapYear(year) {
        return fmt.Sprintf("12.09.%d", year)
    } else {
        return fmt.Sprintf("13.09.%d", year)
    }
}

func main() {
    var year int
    fmt.Scanln(&year)
    result := dayOfProgrammer(year)
    fmt.Println(result)
}
