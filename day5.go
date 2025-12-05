package main

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
)

func day5() {
	// Read input file
	data, err := fetchInput(5)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input := string(data)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Part 1
	result1 := day5Part1(lines)
	fmt.Printf("Part 1: %d\n", result1)

	// Part 2
	result2 := day5Part2(lines)
	fmt.Printf("Part 2: %d\n", result2)
}

type Range struct {
    Start int64
    End   int64
}

func parseRanges(lines []string) ([]Range, int) {
	ranges := []Range{}
	inputStart := 0

	for i, line := range lines {
		if line == "" {
			inputStart = i + 1
			break
		}
		parts := strings.Split(line, "-")
		start, _ := strconv.ParseInt(parts[0], 10, 64)
		end, _ := strconv.ParseInt(parts[1], 10, 64)
		ranges = append(ranges, Range{Start: start, End: end})
	}

	return ranges, inputStart
}

func mergeRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return ranges
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	merged := []Range{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		lastIdx := len(merged) - 1
		if ranges[i].Start <= merged[lastIdx].End + 1 {
			merged[lastIdx].End = max(merged[lastIdx].End, ranges[i].End)
		} else {
			merged = append(merged, ranges[i])
		}
	}

	return merged
}

func buildIntervalArray(ranges []Range) []int64 {
	intervals := []int64{}
	for _, r := range ranges {
		intervals = append(intervals, r.Start)
		intervals = append(intervals, r.End)
	}
	return intervals
}

func day5Part1(lines []string) int {
	ranges, inputStart := parseRanges(lines)
	mergedRanges := mergeRanges(ranges)
	intervals := buildIntervalArray(mergedRanges)

	total := 0
	for i := inputStart; i < len(lines); i++ {
		ingredientID, _ := strconv.ParseInt(lines[i], 10, 64)

		// Binary search to find the position of ingredientID in intervals
		index := sort.Search(len(intervals), func(i int) bool {
			return intervals[i] >= ingredientID
		})

		if index%2 != 0 {
			total++
		}
	}

	return total
}

func day5Part2(lines []string) int {
	ranges, _ := parseRanges(lines)
	mergedRanges := mergeRanges(ranges)
	total := 0
	for _, r := range mergedRanges {
		total += int(r.End - r.Start + 1)
	}
	return total
}