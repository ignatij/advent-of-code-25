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
			k, _ := strconv.Atoi(strings.Split(line, "L")[1])
			base := curr % 100
			if base == 0 {
				total += k / 100
			} else if k >= base {
				total += 1 + (k-base)/100
			}
			curr = (curr - (k % 100) + 100) % 100
		} else if strings.Contains(line, "R") {
			k, _ := strconv.Atoi(strings.Split(line, "R")[1])
			base := (100 - curr) % 100
			if base == 0 {
				total += k / 100
			} else if k >= base {
				total += 1 + (k-base)/100
			}
			curr = (curr + k) % 100
		}
	}
	fmt.Println(total)
}
