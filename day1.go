package main

import (
	"fmt"
	"strings"
	"strconv"
)

func mod(a, b int) int {
    result := a % b
    if result < 0 {
        result += b
    }
    return result
}

func day1() {
	// Read input file
	data, err := fetchInput(1)
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

	input := string(data)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Part 1
	result1 := day1Part1(lines)
	fmt.Printf("Part 1: %d\n", result1)

	// Part 2
	result2 := day1Part2(lines)
	fmt.Printf("Part 2: %d\n", result2)
}

func day1Part1(lines []string) int {
	current_pos := 50
	count := 0
	for _, line := range lines {
		direction := string(line[0])
		magnitude, err := strconv.Atoi(line[1:])
		if err != nil {
			// Handle error (e.g., if it's not a valid number)
			fmt.Printf("Error converting to int: %v\n", err)
			continue
		}
		if direction == "R" {
			current_pos += magnitude
		} else if direction == "L" {
			current_pos -= magnitude
		}
		current_pos = mod(current_pos, 100)
		if current_pos == 0 {
			count += 1
		}
	}
	return count
}

func day1Part2(lines []string) int {
	current_pos := 50
	count := 0
	for _, line := range lines {
		direction := string(line[0])
		magnitude, err := strconv.Atoi(line[1:])
		if err != nil {
			// Handle error (e.g., if it's not a valid number)
			fmt.Printf("Error converting to int: %v\n", err)
			continue
		}
		count += magnitude / 100
		magnitude %= 100
		initial_pos := current_pos
		if direction == "R" {
			current_pos += magnitude
		} else if direction == "L" {
			current_pos -= magnitude
		}
		if current_pos >= 100 {
			count += 1
		}
		if current_pos <= 0 && initial_pos > 0 {
			count += 1
		}
		current_pos = mod(current_pos, 100)
	}
	return count
}