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
		stack := make([]byte, 0, len(bank))
		k := len(bank) - 12
		for i := 0; i < len(bank); i++ {
			if len(stack) == 0 {
				stack = append(stack, bank[i])
				continue
			}
			d := bank[i]
			for k > 0 && len(stack) > 0 && stack[len(stack)-1] < d {
				stack = stack[:len(stack)-1]
				k--
			}
			stack = append(stack, d)
		}
		stack = stack[:len(stack)-k]

		max, _ := strconv.Atoi(string(stack))
		total += max
	}

	fmt.Println(total)
}
