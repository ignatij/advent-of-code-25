package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInvalidId(id int64) bool {
	curr := strconv.FormatInt(id, 10)
	half := len(curr) / 2
	for k := 1; k <= half; k++ {
		if len(curr)%k != 0 {
			continue
		}
		ok := true
		s := curr[0:k]
		for i := k; i < len(curr); i += k {
			if s != curr[i:i+k] {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		for idRange := range strings.SplitSeq(scanner.Text(), ",") {
			idRangeParts := strings.Split(idRange, "-")
			start, _ := strconv.Atoi(idRangeParts[0])
			end, _ := strconv.Atoi(idRangeParts[1])
			for i := start; i <= end; i++ {
				if isInvalidId(int64(i)) {
					total += i
				}
			}
		}
	}
	fmt.Println(total)
}
