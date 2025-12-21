package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Box struct {
	X int
	Y int
	Z int
}

type Edge struct {
	A, B int
	dist int64
}

func find(parent []int, i int) int {
	if parent[i] != i {
		parent[i] = find(parent, parent[i])
	}
	return parent[i]
}

func union(parent []int, a, b int) {
	rootA := find(parent, a)
	rootB := find(parent, b)
	parent[rootB] = rootA
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	boxes := []Box{}
	edges := []Edge{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		box := Box{X: x, Y: y, Z: z}
		boxes = append(boxes, box)
	}

	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			boxA := boxes[i]
			boxB := boxes[j]
			dx := int64(boxB.X - boxA.X)
			dy := int64(boxB.Y - boxA.Y)
			dz := int64(boxB.Z - boxA.Z)
			dist := dx*dx + dy*dy + dz*dz
			edge := Edge{A: i, B: j, dist: dist}
			edges = append(edges, edge)
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	parent := make([]int, len(boxes))
	for i := 0; i < len(boxes); i++ {
		parent[i] = i
	}

	for i := 0; i < 1000; i++ {
		union(parent, edges[i].A, edges[i].B)
	}

	type KV struct {
		key   int
		value int
	}
	m := make([]KV, len(parent))

	for i := 0; i < len(parent); i++ {
		root := find(parent, parent[i])
		m[root].value++
	}
	sort.Slice(m, func(i, j int) bool {
		return m[i].value > m[j].value
	})
	m = m[:3]
	product := 1
	for _, v := range m {
		product *= v.value
	}
	fmt.Println(product)
}
