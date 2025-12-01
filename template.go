package main

import (
	"fmt"
	"os"
	"strings"
)

func dayX() {
	// Read input file
	data, err := os.ReadFile("input/dayX.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	input := string(data)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Part 1
	result1 := dayXPart1(lines)
	fmt.Printf("Part 1: %d\n", result1)

	// Part 2
	result2 := dayXPart2(lines)
	fmt.Printf("Part 2: %d\n", result2)
}

func dayXPart1(lines []string) int {
	// TODO: Implement part 1 solution
	return 0
}

func dayXPart2(lines []string) int {
	// TODO: Implement part 2 solution
	return 0
}