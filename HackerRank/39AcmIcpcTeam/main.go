package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'acmTeam' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts STRING_ARRAY topic as parameter.
 */

func acmTeam(topic []string) []int32 {
    var maxTopics int32 = 0
	var maxTeams int32 = 0

	n := len(topic) // n >= 1
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			// count the number of topics by i and j
			var count int32 = 0
			for k := 0; k < len(topic[i]); k++ {
				if topic[i][k] == '1' || topic[j][k] == '1' {
					count++
				}
			}

			if count > maxTopics {
				maxTopics = count
				maxTeams = 1
			} else if count == maxTopics {
				maxTeams++
			}
		}
	}

    return []int32{maxTopics, maxTeams}
}

func main() {
	// set up buffered input using scanner
	scanner := bufio.NewScanner(os.Stdin)

	// read the first line of input (n and  m)
	var n int
	if scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		n, _ = strconv.Atoi(fields[0]) // number of attendees
		// m, _ := strconv.Atoi(fields[1]) // number of topics
	}

	// read the next n lines of input
	topics := make([]string, n)
	for i := range n {
		if scanner.Scan() {
			topics[i] = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}

	result := acmTeam(topics)
	fmt.Printf("%v\n%v\n", result[0], result[1])
}