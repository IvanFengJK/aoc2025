package main

import (
	"fmt"
	"strings"
	"math"
)

func day3() {
	// Read input file
	data, err := fetchInput(3)
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

	input := string(data)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Part 1
	result1 := day3Part1(lines)
	fmt.Printf("Part 1: %d\n", result1)

	// Part 2
	result2 := day3Part2(lines)
	fmt.Printf("Part 2: %d\n", result2)
}

func argmaxDigit(s string, start, end int) (int, int) {
    maxDigit := -1
    maxIdx := -1
    
    for i := start; i < end; i++ {
		// Convert to actual number (0-9)
        digit := int(s[i] - '0')
        if digit > maxDigit {
            maxDigit = digit
            maxIdx = i
        }
    }
    
    return maxIdx, maxDigit
}

func findNMaxDigits(s string, n int, startPos int) int {
    if n == 0 {
        return 0
    }

    endPos := len(s) - n + 1
    pos, value := argmaxDigit(s, startPos, endPos)
    
    // Recursively find remaining (n-1) digits
    return value * int(math.Pow10(n-1)) + findNMaxDigits(s, n-1, pos+1)
}

func day3Part1(lines []string) int {
	total := 0
	for _, line := range lines {
		total += findNMaxDigits(line, 2, 0)
	}
	return total
}

func day3Part2(lines []string) int {
	total := 0
	for _, line := range lines {
		total += findNMaxDigits(line, 12, 0)
	}
	return total
}