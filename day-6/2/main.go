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
	rawLines := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		sb := strings.Builder{}
		rawLines = append(rawLines, []byte(line))
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
	// total := int64(0)
	// for j := 0; j < len(parsedLines[0]); j++ {
	// 	sign := parsedLines[len(parsedLines)-1][j]
	// 	op1, _ := strconv.Atoi(parsedLines[0][j])
	// 	op2, _ := strconv.Atoi(parsedLines[1][j])
	// 	acc := doOp(sign, int64(op1), int64(op2))
	// 	for i := 2; i < len(parsedLines)-1; i++ {
	// 		op3, _ := strconv.Atoi(parsedLines[i][j])
	// 		acc = doOp(sign, acc, int64(op3))
	// 	}
	// 	total += acc
	// }

	sign := ""
	total := int64(0)
	arr := []int64{}
	for j := 0; j < len(rawLines[0]); j++ {
		acc := int64(0)
		_ = acc
		if rawLines[len(rawLines)-1][j] == '*' || rawLines[len(rawLines)-1][j] == '+' {
			sign = string(rawLines[len(rawLines)-1][j])
		}
		sb := strings.Builder{}
		for i := 0; i < len(rawLines)-1; i++ {
			sb.WriteByte(rawLines[i][j])
		}
		if strings.TrimSpace(sb.String()) == "" {
			acc := int64(arr[0])
			for i := 1; i < len(arr); i++ {
				acc = doOp(sign, acc, arr[i])
			}
			total += acc
			arr = []int64{}
		} else {
			curr, _ := strconv.Atoi(strings.TrimSpace(sb.String()))
			arr = append(arr, int64(curr))
		}

	}
	acc := int64(arr[0])
	for i := 1; i < len(arr); i++ {
		acc = doOp(sign, acc, arr[i])
	}
	total += acc
	fmt.Println(total)
}
