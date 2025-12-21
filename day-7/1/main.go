package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var visited [][]bool
var matrix [][]string

func countBeams(x, y int) int {
	if x >= len(matrix) || visited[x][y] {
		return 0
	}
	visited[x][y] = true
	if matrix[x][y] == "^" {
		if y > 0 && y < len(matrix[x])-1 {
			return 1 + countBeams(x, y-1) + countBeams(x, y+1)
		}
	}

	return countBeams(x+1, y)
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	matrix = [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		matrix = append(matrix, row)

	}
	startX, startY := 0, 0

	visited = make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}

	for i := range matrix {
		for j := range matrix[i] {
			cell := matrix[i][j]
			if cell == "S" {
				startX = i
				startY = j
			}
		}
	}
	fmt.Println(countBeams(startX, startY))

}
