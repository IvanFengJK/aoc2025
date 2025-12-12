package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day6() {
	// Read input file
	data, err := fetchInput(6)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input := string(data)
	lines := strings.Split(input, "\n")

	// Part 1
	result1 := day6Part1(lines)
	fmt.Printf("Part 1: %d\n", result1)

	// Part 2
	result2 := day6Part2(lines)
	fmt.Printf("Part 2: %d\n", result2)
}

func day6Part1(lines []string) int {
	re := regexp.MustCompile(`\d+`)

	var homeworkInput [][]int

	for i := 0; i < len(lines)-1; i++ {
		matches := re.FindAllString(lines[i], -1)
		var numbers []int
		for _, match := range matches {
			num, _ := strconv.Atoi(match)
			numbers = append(numbers, num)
		}

		homeworkInput = append(homeworkInput, numbers)
	}

	re = regexp.MustCompile(`[*+\-/]`)
	operators := re.FindAllString(lines[len(lines)-1], -1)

	total := 0
	for i, op := range operators {
		// Extract column values
		var columnValues []int
		for _, row := range homeworkInput {
			columnValues = append(columnValues, row[i])
		}
		total += applyOperator(op, columnValues)
	}

	return total
}

func day6Part2(lines []string) int {
	numColumns := len(lines[0])
	numRows := len(lines)
	var columnInputs []int
	total := 0

	for col := numColumns - 1; col >= 0; col-- {
		// Extract column as string
		var columnChars []string
		for row := 0; row < numRows-1; row++ {
			if col < len(lines[row]) {
				columnChars = append(columnChars, string(lines[row][col]))
			}
		}
		columnStr := strings.TrimSpace(strings.Join(columnChars, ""))

		// Check if empty (operator column)
		if columnStr == "" {
			operator := string(lines[numRows-1][col+1])
			total += applyOperator(operator, columnInputs)
			columnInputs = []int{}
		} else {
			num, err := strconv.Atoi(columnStr)
			if err != nil {
				fmt.Println("Not a valid number:", err)
			} else {
				columnInputs = append(columnInputs, num)
			}
		}
	}

	// Process final group
	operator := string(lines[numRows-1][0])
	total += applyOperator(operator, columnInputs)

	return total
}

// applyOperator applies the given operator to a slice of values
func applyOperator(operator string, values []int) int {
	switch operator {
	case "+":
		sum := 0
		for _, val := range values {
			sum += val
		}
		return sum
	case "*":
		product := 1
		for _, val := range values {
			product *= val
		}
		return product
	default:
		panic("Unsupported operator: " + operator)
	}
}
