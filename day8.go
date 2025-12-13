package main

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
)

const (
	numConnection   = 1000
)

type Point3D struct {
    x, y, z int
}

type Edge struct {
	i, j     int
	distance int
}

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

func day8() {
	// Read input file
	data, err := fetchInput(8)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input := string(data)
	lines := strings.Split(input, "\n")

	// Part 1
	result1 := day8Part1(lines)
	fmt.Printf("Part 1: %d\n", result1)

	// Part 2
	result2 := day8Part2(lines)
	fmt.Printf("Part 2: %d\n", result2)
}

func distance(p1, p2 Point3D) int {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	dz := p1.z - p2.z
	return dx*dx + dy*dy + dz*dz
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool{
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	
	// If the root nodes are the same, they are already connected
	if rootX == rootY {
		return false
	}

	// Union by rank
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}

	return true
}

func parsePoints(lines []string) []Point3D {
	points := []Point3D{}
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		z, err := strconv.Atoi(parts[2])
		if err != nil {
			continue
		}

		points = append(points, Point3D{x, y, z})
	}
	return points
}

func createSortedEdges(points []Point3D) []Edge {
	edges := []Edge{}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := distance(points[i], points[j])
			edges = append(edges, Edge{i, j, dist})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	return edges
}

func day8Part1(lines []string) int {
	points := parsePoints(lines)
	edges := createSortedEdges(points)

	uf := NewUnionFind(len(points))
	for i := 0; i < numConnection; i++ {
		uf.Union(edges[i].i, edges[i].j)
	}

	circuits := make(map[int]int)
	for i := 0; i < len(points); i++ {
		circuits[uf.Find(i)]++
	}

	top3 := [3]int{0, 0, 0}

	for _, size := range circuits {
		if size > top3[2] {
			if size > top3[0] {
				top3[2] = top3[1]
				top3[1] = top3[0]
				top3[0] = size
			} else if size > top3[1] {
				top3[2] = top3[1]
				top3[1] = size
			} else {
				top3[2] = size
			}
		}
	}
	
	return top3[0] * top3[1] * top3[2]
}

func day8Part2(lines []string) int {
	points := parsePoints(lines)
	edges := createSortedEdges(points)

	uf := NewUnionFind(len(points))
	var p1, p2 Point3D
	for i := 0; i < len(edges); i++ {
		if uf.Union(edges[i].i, edges[i].j) {
			p1 = points[edges[i].i]
			p2 = points[edges[i].j]
		}
	}

	return p1.x * p2.x
}
