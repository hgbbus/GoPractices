package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter a line: ")
    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading:", err)
        return
    }

    // PAY ATTENTION: Remove trailing newline
    input = strings.TrimSpace(input)

    fmt.Println("You entered:", input)
}