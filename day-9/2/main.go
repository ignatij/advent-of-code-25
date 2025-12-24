package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func unique(arr []int) []int {
	seen := make(map[int]bool)
	for _, v := range arr {
		seen[v] = true
	}
	out := make([]int, 0, len(seen))
	for v := range seen {
		out = append(out, v)
	}
	slices.Sort(out)
	return out
}

func imin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func imax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	points := []Point{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, Point{X: x, Y: y})
	}

	// coordinate compression padding each red point with neighbors
	xs := []int{}
	ys := []int{}
	for _, p := range points {
		xs = append(xs, p.X, p.X-1, p.X+1)
		ys = append(ys, p.Y, p.Y-1, p.Y+1)
	}
	xs = unique(xs)
	ys = unique(ys)

	xsMap := make(map[int]int, len(xs))
	for i, v := range xs {
		xsMap[v] = i
	}
	ysMap := make(map[int]int, len(ys))
	for i, v := range ys {
		ysMap[v] = i
	}

	type CompressedPoint struct {
		X, Y   int
		Xi, Yi int
	}
	compressed := make([]CompressedPoint, len(points))
	for i, p := range points {
		compressed[i] = CompressedPoint{X: p.X, Y: p.Y, Xi: xsMap[p.X], Yi: ysMap[p.Y]}
	}

	w := len(xs) - 1
	h := len(ys) - 1

	hWall := make([][]bool, len(ys))
	for i := range hWall {
		hWall[i] = make([]bool, w)
	}
	vWall := make([][]bool, len(xs))
	for i := range vWall {
		vWall[i] = make([]bool, h)
	}

	for i := 0; i < len(points); i++ {
		a := compressed[i]
		b := compressed[(i+1)%len(points)]
		// mark walls along polygon edges (axis-aligned)
		if a.X == b.X {
			x := a.Xi
			y0 := imin(a.Yi, b.Yi)
			y1 := imax(a.Yi, b.Yi)
			for y := y0; y < y1; y++ {
				vWall[x][y] = true
			}
		}
		if a.Y == b.Y {
			y := a.Yi
			x0 := imin(a.Xi, b.Xi)
			x1 := imax(a.Xi, b.Xi)
			for x := x0; x < x1; x++ {
				hWall[y][x] = true
			}
		}
	}

	outside := make([][]bool, w)
	for i := 0; i < w; i++ {
		outside[i] = make([]bool, h)
	}

	queue := make([][2]int, 0)
	push := func(x, y int) {
		if x < 0 || x >= w || y < 0 || y >= h {
			return
		}
		if outside[x][y] {
			return
		}
		outside[x][y] = true
		queue = append(queue, [2]int{x, y})
	}

	for i := 0; i < w; i++ {
		push(i, 0)
		if h > 1 {
			push(i, h-1)
		}
	}
	for j := 0; j < h; j++ {
		push(0, j)
		if w > 1 {
			push(w-1, j)
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		x, y := cur[0], cur[1]
		if x > 0 && !vWall[x][y] {
			push(x-1, y)
		}
		if x+1 < w && !vWall[x+1][y] {
			push(x+1, y)
		}
		if y > 0 && !hWall[y][x] {
			push(x, y-1)
		}
		if y+1 < h && !hWall[y+1][x] {
			push(x, y+1)
		}
	}

	for i := 0; i < len(points); i++ {
		a := points[i]
		b := points[(i+1)%len(points)]
		// ensure walkway cells are treated as inside even if flood fill leaked
		if a.X == b.X {
			x := a.X
			xi0 := xsMap[x]
			xi1 := xsMap[x+1]
			yi0 := ysMap[imin(a.Y, b.Y)]
			yi1 := ysMap[imax(a.Y, b.Y)+1]
			for xi := xi0; xi < xi1 && xi < w; xi++ {
				for yi := yi0; yi < yi1 && yi < h; yi++ {
					outside[xi][yi] = false
				}
			}
		} else {
			y := a.Y
			xi0 := xsMap[imin(a.X, b.X)]
			xi1 := xsMap[imax(a.X, b.X)+1]
			yi0 := ysMap[y]
			yi1 := ysMap[y+1]
			for xi := xi0; xi < xi1 && xi < w; xi++ {
				for yi := yi0; yi < yi1 && yi < h; yi++ {
					outside[xi][yi] = false
				}
			}
		}
	}

	// build prefix sums of "outside" cells for O(1) rectangle queries
	prefix := make([][]int, w+1)
	for i := 0; i <= w; i++ {
		prefix[i] = make([]int, h+1)
	}
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			val := 0
			if outside[i][j] {
				val = 1
			}
			prefix[i+1][j+1] = prefix[i][j+1] + prefix[i+1][j] - prefix[i][j] + val
		}
	}

	// brute-force every pair of red corners in original coordinates
	maxArea := int64(0)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			x0 := imin(points[i].X, points[j].X)
			x1 := imax(points[i].X, points[j].X)
			y0 := imin(points[i].Y, points[j].Y)
			y1 := imax(points[i].Y, points[j].Y)

			xi0 := xsMap[x0]
			xi1 := xsMap[x1+1]
			yi0 := ysMap[y0]
			yi1 := ysMap[y1+1]

			count := prefix[xi1][yi1] - prefix[xi0][yi1] - prefix[xi1][yi0] + prefix[xi0][yi0]
			if count == 0 {
				width := int64(x1 - x0 + 1)
				height := int64(y1 - y0 + 1)
				area := width * height
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	fmt.Println(maxArea)
}
