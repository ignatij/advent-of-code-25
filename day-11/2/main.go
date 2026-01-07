package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MemoEntry struct {
	node    string
	seenDac bool
	seenFft bool
}

var memo = make(map[MemoEntry]int)

// DFS
func countPaths(node string, graph map[string][]string, seenDac, seenFft bool) int {
	if node == "dac" {
		seenDac = true
	}
	if node == "fft" {
		seenFft = true
	}
	if node == "out" {
		if seenDac && seenFft {
			return 1
		}
		return 0
	}

	// here it's definitely needed, otherwise it doesn't finish
	if _, ok := memo[MemoEntry{node, seenDac, seenFft}]; ok {
		return memo[MemoEntry{node, seenDac, seenFft}]
	}

	count := 0
	for _, neighbor := range graph[node] {
		count += countPaths(neighbor, graph, seenDac, seenFft)
	}
	memo[MemoEntry{node, seenDac, seenFft}] = count

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

	fmt.Println(countPaths("svr", graph, false, false))
}
