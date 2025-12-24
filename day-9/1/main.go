package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	points := []Point{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		i, _ := strconv.Atoi(parts[0])
		j, _ := strconv.Atoi(parts[1])
		points = append(points, Point{X: i, Y: j})
	}

	max := int64(0)

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			length := int64(math.Abs(float64(points[j].X-points[i].X)) + 1)
			width := int64(math.Abs(float64(points[j].Y-points[i].Y)) + 1)
			area := int64(length * width)
			if area > max {
				max = area
			}
		}
	}

	fmt.Println(max)
}
