package main

import (
	"fmt"
)

// Approach 1: Union-Find
func findCircleNum1(isConnected [][]int) int {
	n := len(isConnected)
	parent := make([]int, n)
	for i := range n {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if (parent[x] != x) {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	var union func(int, int)
	union = func(x int, y int) {
		rootX := find(x)
		rootY := find(y)
		if rootX != rootY {
			parent[rootY] = rootX
		}
	}

	for i := range n {
		for j := range n {
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}

	count := 0
	for i := range n {
		if parent[i] == i {
			count++
		}
	}

	return count
}

// Approach 2: DFS
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)

	var dfs func(int)
	dfs = func(i int) {
		visited[i] = true
		for j := range n {
			if isConnected[i][j] == 1 && !visited[j] {
				dfs(j)
			}
		}
	}

	count := 0
	for i := range n {
		if !visited[i] {
			count++
			dfs(i)
		}
	}

	return count
}

func main() {
	// test case 1
	isConnected1 := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(findCircleNum(isConnected1)) // expected output: 2

	// test case 2
	isConnected2 := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(findCircleNum(isConnected2)) // expected output: 3
}