package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start int64
	End   int64
}

func binarySearch(ranges []Range, target int64) bool {
	left, right := 0, len(ranges)-1
	for left <= right {
		mid := (left + right) / 2
		if ranges[mid].Start <= target && ranges[mid].End >= target {
			return true
		} else if ranges[mid].Start > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	rangeMode := true
	ranges := make([]Range, 0)
	fruits := make([]int64, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			rangeMode = false
			continue
		}
		if rangeMode {
			parts := strings.Split(line, "-")
			startInt, _ := strconv.Atoi(parts[0])
			endInt, _ := strconv.Atoi(parts[1])
			ranges = append(ranges, Range{Start: int64(startInt), End: int64(endInt)})
		} else {
			fruit, _ := strconv.Atoi(line)
			fruits = append(fruits, int64(fruit))
		}
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})
	merged := make([]Range, 0, len(ranges))
	if len(ranges) > 0 {
		curr := ranges[0]
		for i := 1; i < len(ranges); i++ {
			r := ranges[i]
			if r.Start <= curr.End {
				if r.End > curr.End {
					curr.End = r.End
				}
			} else {
				merged = append(merged, curr)
				curr = r
			}
		}
		merged = append(merged, curr)
	}

	ranges = merged
	count := 0
	for _, fruit := range fruits {
		if binarySearch(ranges, fruit) {
			count++
		}
	}
	fmt.Println(count)
}
