package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	graph := make([][]string, 0)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		graph = append(graph, make([]string, len(line)))
		for j := 0; j < len(line); j++ {
			graph[i][j] = string(line[j])
		}
		i++
	}
	count := 0
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			if graph[i][j] == "@" {
				adjacent := 0
				if i > 0 && graph[i-1][j] == "@" {
					adjacent++
				}
				if i < len(graph)-1 && graph[i+1][j] == "@" {
					adjacent++
				}
				if j > 0 && graph[i][j-1] == "@" {
					adjacent++
				}
				if j < len(graph[i])-1 && graph[i][j+1] == "@" {
					adjacent++
				}
				if i > 0 && j > 0 && graph[i-1][j-1] == "@" {
					adjacent++
				}
				if i > 0 && j < len(graph[i])-1 && graph[i-1][j+1] == "@" {
					adjacent++
				}
				if i < len(graph)-1 && j > 0 && graph[i+1][j-1] == "@" {
					adjacent++
				}
				if i < len(graph)-1 && j < len(graph[i])-1 && graph[i+1][j+1] == "@" {
					adjacent++
				}
				if adjacent < 4 {
					count++
				}
			}

		}
	}
	fmt.Println(count)
}
