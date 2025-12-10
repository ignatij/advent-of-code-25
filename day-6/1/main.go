package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func doOp(sign string, op1, op2 int64) int64 {
	switch {
	case sign == "+":
		return op1 + op2
	case sign == "*":
		return op1 * op2
	default:
		return 0
	}
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	parsedLines := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		sb := strings.Builder{}
		for i := 0; i < len(line); i++ {
			if line[i] == ' ' {
				sb.WriteString(" ")
				j := i
				for j+1 < len(line) {
					if line[j+1] != ' ' {
						break
					}
					j++
				}
				i = j
			} else {
				sb.WriteRune(rune(line[i]))
			}
		}
		parsedLines = append(parsedLines, strings.Split(strings.TrimSpace(sb.String()), " "))
	}
	total := int64(0)
	for j := 0; j < len(parsedLines[0]); j++ {
		sign := parsedLines[len(parsedLines)-1][j]
		op1, _ := strconv.Atoi(parsedLines[0][j])
		op2, _ := strconv.Atoi(parsedLines[1][j])
		acc := doOp(sign, int64(op1), int64(op2))
		for i := 2; i < len(parsedLines)-1; i++ {
			op3, _ := strconv.Atoi(parsedLines[i][j])
			acc = doOp(sign, acc, int64(op3))
		}
		total += acc
	}
	fmt.Println(total)
}
