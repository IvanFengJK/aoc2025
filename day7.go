package main

import (
	"fmt"
	"strings"
)

type Point struct {
    x, y int
}

func day7() {
	// Read input file
	data, err := fetchInput(7)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input := string(data)
	lines := strings.Split(input, "\n")

	// Part 1
	result1 := day7Part1(lines)
	fmt.Printf("Part 1: %d\n", result1)

	// Part 2
	result2 := day7Part2(lines)
	fmt.Printf("Part 2: %d\n", result2)
}

func day7Part1(lines []string) int {
	total := 0
	queue := []Point{}
	visited := make(map[Point]bool)

	pos := strings.Index(lines[0], "S")
	queue = append(queue, Point{1, pos})
	depth := len(lines)
	queueNumber := 0

	for queueNumber < len(queue) {
		front := queue[queueNumber]
		queueNumber++

		if front.x >= depth || visited[front] {
			continue
		}

		visited[front] = true
		next := lines[front.x][front.y] 
		if next == '.' {
			queue = append(queue, Point{front.x + 1, front.y})
		} else if next == '^' {
			total++
			queue = append(queue, Point{front.x, front.y - 1})
			queue = append(queue, Point{front.x, front.y + 1})
		}
	}

	return total
}

func dp(lines []string, point Point, memo map[Point]int) int {
	// Base case: exited
	if point.x >= len(lines) {
		return 1
	}

	if val, exists := memo[point]; exists{
		return val
	}

	next := lines[point.x][point.y]
	var result int

	if next == '.' || next == 'S' {
		result = dp(lines, Point{point.x+1, point.y}, memo)
	} else if next == '^' {
		// Split: left and right beams, both continue down
		leftResult := dp(lines, Point{point.x+1, point.y - 1}, memo)
		rightResult := dp(lines, Point{point.x+1, point.y + 1}, memo)
		result = leftResult + rightResult
	}

	memo[point] = result
	return result
}

func day7Part2(lines []string) int {
	memo := make(map[Point]int)
	pos := strings.Index(lines[0], "S")
	point := Point{0, pos}

	return dp(lines, point, memo)
}
