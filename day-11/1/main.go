package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var memo = make(map[string]int)

// DFS
func countPaths(node string, graph map[string][]string) int {
	if node == "out" {
		return 1
	}

	// although not needed, I added memoization to improve performance
	if memo[node] != 0 {
		return memo[node]
	}

	count := 0
	for _, neighbor := range graph[node] {
		count += countPaths(neighbor, graph)
	}
	memo[node] = count

	return count
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	graph := map[string][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		graph[parts[0]] = strings.Split(parts[1], " ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(countPaths("you", graph))
}
