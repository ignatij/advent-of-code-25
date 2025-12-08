package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		bank := scanner.Text()
		max := 0
		for i := 0; i < len(bank); i++ {
			for j := i + 1; j < len(bank); j++ {
				curr, _ := strconv.Atoi(fmt.Sprintf("%c%c", bank[i], bank[j]))
				if curr > max {
					max = curr
				}
			}
		}
		total += max
	}
	fmt.Println(total)
}
