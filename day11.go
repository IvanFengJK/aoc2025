package main

import (
	"fmt"
	"strings"
)

func day11() {
	// Read input file
	data, err := fetchInput(11)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input := string(data)
	lines := strings.Split(input, "\n")

	// Part 1
	result1 := day11Part1(lines)
	fmt.Printf("Part 1: %d\n", result1)

	// Part 2
	result2 := day11Part2(lines)
	fmt.Printf("Part 2: %d\n", result2)
}

func parseGraph(lines []string) map[string][]string {
	graph := make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			continue
		}
		node := parts[0]
		neighbors := strings.Split(parts[1], " ")
		graph[node] = neighbors
	}
	return graph
}

func countPaths(graph map[string][]string, node string, memo map[string]int) int {
	// Base case: reached the "out" node
	if node == "out" {
		return 1
	}

	if val, exists := memo[node]; exists {
		return val
	}
	total := 0
	neighbors := graph[node]
	for _, neighbor := range neighbors {
		total += countPaths(graph, neighbor, memo)
	}
	memo[node] = total
	return total
}

func day11Part1(lines []string) int {
	graph := parseGraph(lines)
	memo := make(map[string]int)
	return countPaths(graph, "you", memo)
}

type Requirements struct {
	node       string
	visitedDac bool
	visitedFft bool
}

func day11Part2(lines []string) int {
	graph := parseGraph(lines)
	memo := make(map[Requirements]int)
	startReq := Requirements{
		node:       "svr",
		visitedDac: false,
		visitedFft: false,
	}
	return countPathsWithRequirements(graph, startReq, memo)
}

func countPathsWithRequirements(graph map[string][]string, req Requirements, memo map[Requirements]int) int {
	// Base case: reached the "out" node
	if req.node == "out" {
		if req.visitedDac && req.visitedFft {
			return 1
		}
		return 0
	}

	if req.node == "dac" {
		req.visitedDac = true
	}

	if req.node == "fft" {
		req.visitedFft = true
	}

	if val, exists := memo[req]; exists {
		return val
	}

	total := 0
	neighbors := graph[req.node]
	for _, neighbor := range neighbors {
		total += countPathsWithRequirements(graph, Requirements{
			node:       neighbor,
			visitedDac: req.visitedDac,
			visitedFft: req.visitedFft,
		}, memo)
	}
	memo[req] = total
	return total
}
