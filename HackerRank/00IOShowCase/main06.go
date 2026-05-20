package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)

	for {
		line, err := scanner.ReadString('\n')
		if err != nil {
			break
		}

		fmt.Println("Line:", strings.TrimSpace(line))
	}
}