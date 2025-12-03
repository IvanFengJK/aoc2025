package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
)

func day2() {
	// Read input file
	data, err := fetchInput(2)
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

	input := string(data)
	lines := strings.Split(strings.TrimSpace(input), ",")

	// Part 1
	result1 := day2Part1(lines)
	fmt.Printf("Part 1: %d\n", result1)

	// Part 1
	result1New := day2Part1New(lines)
	fmt.Printf("Part 1: %d\n", result1New)

	// Part 2
	result2 := day2Part2(lines)
	fmt.Printf("Part 2: %d\n", result2)
}

func roundToNearestEven(num int, sign int) int {
	digitCount := len(strconv.Itoa(num))
	if digitCount % 2 == 0 {
		return num
	}

	return int(math.Pow10(digitCount + sign)) + sign
}

func makeRepeated(half int) int {
	repeated, _ := strconv.Atoi(strings.Repeat(strconv.Itoa(half), 2))
	return repeated
}

func firstHalf(num int) int {
	numStr := strconv.Itoa(num)
	half := len(numStr) / 2
	firstHalf, _ := strconv.Atoi(numStr[:half])
	return firstHalf
}

func convertToRepeatedDigitNumber(num int, sign int) int {
    firstHalf := firstHalf(num)
    
    repeated := makeRepeated(firstHalf)
    for sign * repeated < sign * num {
        firstHalf += sign
        repeated = makeRepeated(firstHalf)
    }
    
    return repeated
}


func listRepeatedDigitNumbersInRange(start, end int) []int {
    firstHalfStart := firstHalf(start)
	firstHalfEnd := firstHalf(end)
	numbers := []int{}
	for i := firstHalfStart; i <= firstHalfEnd; i++ {
		numbers = append(numbers, makeRepeated(i))
	}
	return numbers
}

func day2Part1(lines []string) int {
	total := 0
	for _, line := range lines {
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		start = roundToNearestEven(start, 0)
		end = roundToNearestEven(end, -1)

		if end < start {
			continue
		}

		nextNumber := convertToRepeatedDigitNumber(start, 1)
		if nextNumber > end {
			continue
		}

		terminateNumber := convertToRepeatedDigitNumber(end, -1)
		for _, num := range listRepeatedDigitNumbersInRange(nextNumber, terminateNumber) {
			total += num
		}
	}
	return total
}

func checkRepeat(num int, windowSize int) bool {
    s := strconv.Itoa(num)
    n := len(s)

	if n % windowSize != 0 {
        return false
    }
    
    firstWindow := s[:windowSize]
    
    for i := windowSize; i < n; i += windowSize {
        window := s[i:i+windowSize]
        if window != firstWindow {
            return false
        }
    }
    
    return true
}

func getAllFactors(num int) []uint32 {
    digitCount := len(strconv.Itoa(num))
    
    factors := []uint32{}
    
    // Find all factors from 1 to digitCount-1
    for i := 1; i < digitCount; i++ {
        if digitCount % i == 0 {
            factors = append(factors, uint32(i))
        }
    }
    
    return factors
}

func sumRepeatedNumbers(start, rangeEnd int, windowSizes []uint32) int {
    total := 0
    for i := start; i < rangeEnd; i++ {
        for _, windowSize := range windowSizes {
            if checkRepeat(i, int(windowSize)) {
                total += i
                break
            }
        }
    }
    return total
}

func day2Part2(lines []string) int {
	total := 0
	for _, line := range lines {
		parts := strings.Split(line, "-")

		start, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		digitCount := len(strconv.Itoa(start))
		rangeEnd := int(math.Pow10(digitCount))
		factors := getAllFactors(start)

		if end > rangeEnd {
			total += sumRepeatedNumbers(start, rangeEnd, factors)
			factors = getAllFactors(end)
			total += sumRepeatedNumbers(rangeEnd, end+1, factors)
		} else {
			total += sumRepeatedNumbers(start, end+1, factors)
		}
	}
	return total
}

func day2Part1New(lines []string) int {
	total := 0
	for _, line := range lines {
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		digitCountStart := len(strconv.Itoa(start))
		
		initialValue := digitCountStart
		if digitCountStart%2 != 0 {
			initialValue = (digitCountStart + 1)
			start = int(math.Pow10(digitCountStart))
		} 
		
		digitCountEnd := len(strconv.Itoa(end))
		if digitCountEnd%2 != 0 {
			end = int(math.Pow10(digitCountEnd-1)) - 1
		}
		factors := []uint32{uint32(initialValue/2)}
		total += sumRepeatedNumbers(start, end+1, factors)
	}
	return total
}