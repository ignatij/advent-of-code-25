package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var matrix [][]string

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

	dp := make([][]int, len(matrix))
	for i := range matrix {
		dp[i] = make([]int, len(matrix[i]))
		for j := range matrix[i] {
			cell := matrix[i][j]
			if cell == "S" {
				startX = i
				startY = j
			}
		}
		// set terminal nodes to 1
		if i == len(matrix)-1 {
			for j := 0; j < len(matrix[i]); j++ {
				dp[i][j] = 1
			}
		}
	}

	for i := len(matrix) - 2; i >= 0; i-- {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "^" {
				dp[i][j] = dp[i+1][j-1] + dp[i+1][j+1]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}

	fmt.Println(dp[startX][startY])

}
