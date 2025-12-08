package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	scanner := bufio.NewScanner(f)
	curr := 50
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "L") {
			next, _ := strconv.Atoi(strings.Split(line, "L")[1])
			curr = (curr - next) % 100
		} else if strings.Contains(line, "R") {
			next, _ := strconv.Atoi(strings.Split(line, "R")[1])
			curr = (curr + next) % 100
		}
		if curr == 0 {
			total++
		}
	}
	fmt.Println(total)
}
