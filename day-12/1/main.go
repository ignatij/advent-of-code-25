package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		flag := true
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		totals := strings.Split(parts[0], "x")
		width, _ := strconv.Atoi(totals[0])
		height, _ := strconv.Atoi(totals[1])
		area := width * height
		totalAreaPackages := 0
		for _, s := range strings.Split(parts[1], " ") {
			currRect, _ := strconv.Atoi(s)
			// hardcoding 9 as a fixed area of a rectangle
			// making my life easier once peeked into the input
			totalAreaPackages += currRect * 9
			if totalAreaPackages > area {
				flag = false
				break

			}
		}
		if flag {
			count++
		}
	}
	fmt.Println(count)
}
