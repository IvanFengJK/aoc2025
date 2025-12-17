package main

import (
	"fmt"
	"strings"
	"strconv"
)

func day12() {
	// Read input file
	data, err := fetchInput(12)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input := string(data)
	lines := strings.Split(input, "\n\n")

	// Part 1
	result1 := day12Part1(lines)
	fmt.Printf("Part 1: %d\n", result1)
}

type Shape struct {
	area     int
	grid     [][]bool
	variants [][]Coord // All rotations and flips of the shape
}

type Coord struct {
	x, y int
}

func inputToShape(input []string) Shape {
	grid := make([][]bool, len(input))
	area := 0
	for i := 0; i < len(input); i++ {
		grid[i] = make([]bool, len(input[0]))
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == '#' {
				grid[i][j] = true
				area++
			} else {
				grid[i][j] = false
			}
		}
	}

	// Generate all variants (rotations and flips)
	variants := generateVariants(grid)

	return Shape{area: area, grid: grid, variants: variants}
}

// Extract coordinates from a grid
func gridToCoords(grid [][]bool) []Coord {
	coords := []Coord{}
	for y, row := range grid {
		for x, cell := range row {
			if cell {
				coords = append(coords, Coord{x, y})
			}
		}
	}
	return coords
}

// Rotate coordinates 90 degrees clockwise
func rotateCoords(coords []Coord) []Coord {
	rotated := make([]Coord, len(coords))
	for i, c := range coords {
		rotated[i] = Coord{x: -c.y, y: c.x}
	}
	return gridToCoords(coordsToGrid(rotated))
}

// Flip coordinates horizontally
func flipCoords(coords []Coord) []Coord {
	flipped := make([]Coord, len(coords))
	for i, c := range coords {
		flipped[i] = Coord{x: -c.x, y: c.y}
	}
	return gridToCoords(coordsToGrid(flipped))
}

// Convert coords back to grid for normalization
func coordsToGrid(coords []Coord) [][]bool {
	if len(coords) == 0 {
		return [][]bool{}
	}
	maxX, maxY := coords[0].x, coords[0].y
	minX, minY := coords[0].x, coords[0].y
	for _, c := range coords {
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
		if c.x < minX {
			minX = c.x
		}
		if c.y < minY {
			minY = c.y
		}
	}

	grid := make([][]bool, maxY-minY+1)
	for i := range grid {
		grid[i] = make([]bool, maxX-minX+1)
	}

	for _, c := range coords {
		grid[c.y-minY][c.x-minX] = true
	}
	return grid
}

// Generate all unique rotations and flips
func generateVariants(grid [][]bool) [][]Coord {
	baseCoords := gridToCoords(grid)
	variantSet := make(map[string][]Coord)

	// Generate 4 rotations
	current := baseCoords
	for i := 0; i < 4; i++ {
		key := coordsKey(current)
		variantSet[key] = current
		current = rotateCoords(current)
	}

	// Generate 4 rotations of flipped version
	flipped := flipCoords(baseCoords)
	for i := 0; i < 4; i++ {
		key := coordsKey(flipped)
		variantSet[key] = flipped
		flipped = rotateCoords(flipped)
	}

	// Convert map to slice
	variants := make([][]Coord, 0, len(variantSet))
	for _, v := range variantSet {
		variants = append(variants, v)
	}

	return variants
}

// Create a string key for deduplication
func coordsKey(coords []Coord) string {
	key := ""
	for _, c := range coords {
		key += fmt.Sprintf("%d,%d;", c.x, c.y)
	}
	return key
}

// Try to place a shape variant at position (startX, startY) in the grid
func canPlace(grid [][]bool, variant []Coord, startX, startY int) bool {
	height, width := len(grid), len(grid[0])
	for _, c := range variant {
		y, x := startY+c.y, startX+c.x
		if y < 0 || y >= height || x < 0 || x >= width || grid[y][x] {
			return false
		}
	}
	return true
}

// Place a shape variant at position (startX, startY) in the grid
func place(grid [][]bool, variant []Coord, startX, startY int) {
	for _, c := range variant {
		grid[startY+c.y][startX+c.x] = true
	}
}

// Remove a shape variant from position (startX, startY) in the grid
func unplace(grid [][]bool, variant []Coord, startX, startY int) {
	for _, c := range variant {
		grid[startY+c.y][startX+c.x] = false
	}
}

// Backtracking solver to place all presents
func solve(grid [][]bool, presents []int, shapes []Shape, idx int) bool {
	if idx >= len(presents) {
		return true // All presents placed
	}

	// Skip if no presents of this shape needed
	if presents[idx] == 0 {
		return solve(grid, presents, shapes, idx+1)
	}

	height, width := len(grid), len(grid[0])

	// Try to place one present of shape 'idx'
	for variant := range shapes[idx].variants {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if canPlace(grid, shapes[idx].variants[variant], x, y) {
					place(grid, shapes[idx].variants[variant], x, y)
					presents[idx]--

					if solve(grid, presents, shapes, idx) {
						return true
					}

					// Backtrack
					presents[idx]++
					unplace(grid, shapes[idx].variants[variant], x, y)
				}
			}
		}
	}

	return false
}

func canFitPresents(width, height int, presents []int, shapes []Shape) bool {
	// Quick area check first
	totalArea := 0
	for i, count := range presents {
		if i < len(shapes) {
			totalArea += count * shapes[i].area
		}
	}
	if width*height < totalArea {
		return false
	}

	// Create grid and try to solve
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}

	presentsCopy := make([]int, len(presents))
	copy(presentsCopy, presents)

	return solve(grid, presentsCopy, shapes, 0)
}

func day12Part1(lines []string) int {
	// Pre-allocate shapes slice with known capacity
	shapes := make([]Shape, 0, len(lines)-1)
	for i := 0; i < len(lines)-1; i++ {
		parts := strings.Split(lines[i], "\n")
		shape := inputToShape(parts[1:])
		shapes = append(shapes, shape)
	}

	trimmed := strings.TrimSpace(lines[len(lines)-1])
	regions := strings.Split(trimmed, "\n")
	total := 0

	for _, region := range regions {
		parts := strings.Split(region, ":")
		if len(parts) != 2 {
			continue
		}

		dimensions := strings.Split(parts[0], "x")
		if len(dimensions) != 2 {
			continue
		}

		width, err := strconv.Atoi(dimensions[0])
		if err != nil {
			fmt.Printf("Error converting width to int: %v\n", err)
			continue
		}
		height, err := strconv.Atoi(dimensions[1])
		if err != nil {
			fmt.Printf("Error converting height to int: %v\n", err)
			continue
		}

		counts := strings.Fields(parts[1])
		presents := make([]int, len(counts))
		for j, countStr := range counts {
			count, err := strconv.Atoi(countStr)
			if err != nil {
				fmt.Printf("Error converting count to int: %v\n", err)
				continue
			}
			presents[j] = count
		}

		if canFitPresents(width, height, presents, shapes) {
			total++
		}
	}
	return total
}
