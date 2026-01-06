package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

type CompressedPoint struct {
	X, Y   int
	Xi, Yi int
}

func sortInts(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func uniqueInts(arr []int) []int {
	sortInts(arr)
	result := make([]int, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		if i == 0 || arr[i] != arr[i-1] {
			result = append(result, arr[i])
		}
	}
	return result
}

func iMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func iMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	points := make([]Point, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		point := Point{X: x, Y: y}
		points = append(points, point)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	xi := []int{}
	yi := []int{}

	// add neighbors
	for _, point := range points {
		xi = append(xi, point.X, point.X-1, point.X+1)
		yi = append(yi, point.Y, point.Y-1, point.Y+1)
	}
	xi = uniqueInts(xi)
	yi = uniqueInts(yi)

	// compression
	xsMap := make(map[int]int)
	ysMap := make(map[int]int)
	for i, x := range xi {
		xsMap[x] = i
	}
	for i, y := range yi {
		ysMap[y] = i
	}

	compressedPoints := make([]CompressedPoint, len(points))
	for i, point := range points {
		compressedPoints[i] = CompressedPoint{
			X:  point.X,
			Y:  point.Y,
			Xi: xsMap[point.X],
			Yi: ysMap[point.Y],
		}
	}

	// build horizontall and vertical walls
	w := len(xi) - 1
	h := len(yi) - 1
	vWall := make([][]bool, len(xi))
	hWall := make([][]bool, len(yi))

	for i := range vWall {
		vWall[i] = make([]bool, h)
	}
	for i := range hWall {
		hWall[i] = make([]bool, w)
	}

	for i := 0; i < len(compressedPoints); i++ {
		a := compressedPoints[i]
		b := compressedPoints[(i+1)%len(compressedPoints)]
		if a.X == b.X {
			y0 := iMin(a.Yi, b.Yi)
			y1 := iMax(a.Yi, b.Yi)
			for y := y0; y < y1; y++ {
				vWall[a.Xi][y] = true
			}
		}
		if a.Y == b.Y {
			x0 := iMin(a.Xi, b.Xi)
			x1 := iMax(a.Xi, b.Xi)
			for x := x0; x < x1; x++ {
				hWall[a.Yi][x] = true
			}
		}
	}

	// flood fill
	outside := make([][]bool, w)
	for i := range outside {
		outside[i] = make([]bool, h)
	}

	queue := [][2]int{}
	push := func(x, y int) {
		if outside[x][y] {
			return
		}
		outside[x][y] = true
		queue = append(queue, [2]int{x, y})
	}
	for i := 0; i < w; i++ {
		push(i, 0)
		push(i, h-1)
	}
	for i := 0; i < h; i++ {
		push(0, i)
		push(w-1, i)
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		x, y := curr[0], curr[1]

		if x > 0 && !vWall[x][y] {
			push(x-1, y)
		}
		if x < w-1 && !vWall[x+1][y] {
			push(x+1, y)
		}
		if y > 0 && !hWall[y][x] {
			push(x, y-1)
		}
		if y < h-1 && !hWall[y+1][x] {
			push(x, y+1)
		}

	}

	// walkway cleanup because of flood fill
	for i := range points {
		a := points[i]
		b := points[(i+1)%len(points)]

		if a.X == b.X {
			x := a.X
			xi0 := xsMap[x]
			xi1 := xsMap[x+1]
			yi0 := ysMap[iMin(a.Y, b.Y)]
			yi1 := ysMap[iMax(a.Y, b.Y)+1]

			for xi := xi0; xi < xi1; xi++ {
				for yi := yi0; yi < yi1; yi++ {
					outside[xi][yi] = false
				}
			}
		}
		if a.Y == b.Y {
			y := a.Y
			yi0 := ysMap[y]
			yi1 := ysMap[y+1]
			xi0 := xsMap[iMin(a.X, b.X)]
			xi1 := xsMap[iMax(a.X, b.X)+1]

			for xi := xi0; xi < xi1; xi++ {
				for yi := yi0; yi < yi1; yi++ {
					outside[xi][yi] = false
				}
			}
		}
	}

	// prefix calculation
	prefix := make([][]int, w+1)
	for i := range prefix {
		prefix[i] = make([]int, h+1)
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			s := 0
			if outside[i][j] {
				s = 1
			}
			prefix[i+1][j+1] = prefix[i][j+1] + prefix[i+1][j] - prefix[i][j] + s
		}
	}

	// calculate area max
	maxArea := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			x1 := iMin(points[i].X, points[j].X)
			x2 := iMax(points[i].X, points[j].X)
			y1 := iMin(points[i].Y, points[j].Y)
			y2 := iMax(points[i].Y, points[j].Y)

			x1i := xsMap[x1]
			x2i := xsMap[x2+1]
			y1i := ysMap[y1]
			y2i := ysMap[y2+1]

			count := prefix[x2i][y2i] - prefix[x2i][y1i] - prefix[x1i][y2i] + prefix[x1i][y1i]
			if count == 0 {
				area := (x2 - x1 + 1) * (y2 - y1 + 1)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	fmt.Println(maxArea)

}
