package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func catAndMouse(x, y, z int) string {
	dx := abs(x - z)
	dy := abs(y - z)

	if dx < dy {
		return "Cat A"
	} else if dy < dx {
		return "Cat B"
	} else {
		return "Mouse C"
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// first line is a single integer, q
	line, _ := reader.ReadString('\n')
	q, _ := strconv.Atoi(strings.TrimSpace(line))

	for range q {
		line, _ = reader.ReadString('\n')
		parts := strings.Fields(line)
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		result := catAndMouse(x, y, z)
		fmt.Println(result)
	}
}