package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// use meet-in-the-middle algorithm to find the minimum number of presses
func minPresses(target uint64, buttons []uint64) uint64 {
	m := len(buttons) / 2
	bestLeft := make(map[uint64]uint64, 1<<m)
	bestRight := make(map[uint64]uint64, 1<<m)
	bestLeft[0] = 0
	bestRight[0] = 0

	for i := 0; i < m; i++ {
		curr := buttons[i]
		keys := make([]uint64, 0, len(bestLeft))
		for k := range bestLeft {
			keys = append(keys, k)
		}
		for _, k := range keys {
			newEff := curr ^ k
			// total presses are the number of presses needed for k + 1
			total := bestLeft[k] + 1
			if _, ok := bestLeft[newEff]; !ok || total < bestLeft[newEff] {
				bestLeft[newEff] = total
			}
		}
	}

	for i := m; i < len(buttons); i++ {
		curr := buttons[i]
		keys := make([]uint64, 0, len(bestRight))
		for k := range bestRight {
			keys = append(keys, k)
		}
		for _, k := range keys {
			newEff := curr ^ k
			total := bestRight[k] + 1
			if _, ok := bestRight[newEff]; !ok || total < bestRight[newEff] {
				bestRight[newEff] = total
			}
		}
	}

	min := uint64(1<<64 - 1)
	for right := range bestRight {
		left := right ^ target
		if _, ok := bestLeft[left]; ok {
			curr := bestLeft[left] + bestRight[right]
			if curr < min {
				min = curr
			}
		}
	}

	return min
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum uint64
	for scanner.Scan() {
		var target uint64
		var buttons []uint64
		line := scanner.Text()
		for i, s := range strings.Split(line, " ") {
			if i == 0 {
				// find target
				for i := 1; i < len(s)-1; i++ {
					if s[i] == '#' {
						// subtracting 1 to get the correct index of the bitmap
						target |= 1 << uint64(i-1)
					}
				}
				continue
			}
			if i == len(strings.Split(line, " "))-1 {
				// find min presses for machine line
				sum += minPresses(target, buttons)
				continue
			}
			// construct the bitmask array of buttons
			var bitmask uint64
			s = s[1 : len(s)-1]
			parts := strings.Split(s, ",")
			for _, part := range parts {
				x, _ := strconv.Atoi(part)
				bitmask |= 1 << uint64(x)
			}
			buttons = append(buttons, bitmask)
		}
	}
	fmt.Println(sum)
}
