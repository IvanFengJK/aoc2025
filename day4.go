package main

import (
	"fmt"
	"strings"
)

const maxAdjacentRolls = 4

func day4() {
	// Read input file
	data, err := fetchInput(4)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input := string(data)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Part 1
	result1 := day4Part1(lines)
	fmt.Printf("Part 1: %d\n", result1)

	// Part 2
	result2 := day4Part2(lines)
	fmt.Printf("Part 2: %d\n", result2)
}

func parseGrid(lines []string) [][]byte {
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}
	return grid
}

func countAdjacentRolls(grid [][]byte, row, col int) int {
	rows := len(grid)
	cols := len(grid[0])

	// 8 directions: up, down, left, right, and 4 diagonals
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
        {0, -1},           {0, 1},
        {1, -1},  {1, 0},  {1, 1},
	}

	count := 0
	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		// Check bounds
		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
			if grid[newRow][newCol] == '@' {
				count++
			}
		}
	}
	return count
}

func isAccessible(grid [][]byte, row, col int) bool {
	return countAdjacentRolls(grid, row, col) < maxAdjacentRolls
}

func day4Part1(lines []string) int {
	grid := parseGrid(lines)
	total := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '@' && isAccessible(grid, i, j) {
				total++
			}
		}
	}
	return total
}

func day4Part2(lines []string) int {
	grid := parseGrid(lines)
	total := 0

	for {
		removed := 0
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] == '@' && isAccessible(grid, i, j) {
					grid[i][j] = '.'
					removed++
				}
			}
		}
		total += removed
		if removed == 0 {
			break
		}
	}

	return total
}