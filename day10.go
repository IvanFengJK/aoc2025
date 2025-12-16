package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day10() {
	// Read input file
	data, err := fetchInput(10)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input := string(data)
	lines := strings.Split(input, "\n")

	// Part 1
	result1 := day10Part1(lines)
	fmt.Printf("Part 1: %d\n", result1)

	// Part 2
	result2 := day10Part2(lines)
	fmt.Printf("Part 2: %d\n", result2)
}

func parseMachine(line string) (string, [][]string, string) {
	squareBracket := regexp.MustCompile(`\[([^\]]*)\]`)
	roundBrackets := regexp.MustCompile(`\(([^)]*)\)`)
	curlyBrackets := regexp.MustCompile(`\{([^}]*)\}`)

	squareMatch := squareBracket.FindStringSubmatch(line)
	roundMatches := roundBrackets.FindAllStringSubmatch(line, -1)
	curlyMatch := curlyBrackets.FindStringSubmatch(line)

	return squareMatch[1], roundMatches, curlyMatch[1]
}

func constructMatrix(desiredState string, buttons [][]string) [][]int {
	L := len(desiredState)
	B := len(buttons)
	M := make([][]int, L)
	for i := range M {
		M[i] = make([]int, B+1)
	}

	for row, state := range desiredState {
		if state == '#' {
			M[row][B] = 1
		} else {
			M[row][B] = 0
		}
	}

	for row, button := range buttons {
		positions := strings.Split(button[1], ",")
		for _, posStr := range positions {
			pos, err := strconv.Atoi(strings.TrimSpace(posStr))
			if err != nil {
				continue
			}
			if pos >= 0 && pos < L {
				M[pos][row] = 1
			}
		}
	}

	return M
}

func gaussianElimination(M [][]int) [][]int {
	pivotRow := 0
	L := len(M)
	B := len(M[0]) - 1
	
	for pivotCol := 0; pivotCol < B && pivotRow < L; pivotCol++ {
		pivotFound := false
		for i := pivotRow; i < L; i++ {
			if M[i][pivotCol] == 1 {
				M[i], M[pivotRow] = M[pivotRow], M[i]
				pivotFound = true
				break
			}
		}

		if pivotFound {
			for i := 0; i < L; i++ {
				if i != pivotRow && M[i][pivotCol] == 1 {
					for j := 0; j <= B; j++ {
						M[i][j] = M[i][j] ^ M[pivotRow][j]
					}
				}
			}
			pivotRow++
		}
	}

	for i := pivotRow; i < L; i++ {
		if M[i][B] == 1 {
			panic("System is inconsistent! No solution exists.")
		}
	}

	return M
}

func day10Part1(lines []string) int {
	total := 0
	for _, line := range lines {
		desiredState, buttons, _ := parseMachine(line)
		matrix := constructMatrix(desiredState, buttons)
		augmentedMatrix := gaussianElimination(matrix)

		L := len(augmentedMatrix)
		B := len(augmentedMatrix[0]) - 1
		pivotCols := make([]int, 0)
		freeCols := make([]int, 0)
		pivotRow := 0
		for col := 0; col < B; col++ {
			if pivotRow < L && augmentedMatrix[pivotRow][col] == 1 {
				pivotCols = append(pivotCols, col)
				pivotRow++
			} else {
				freeCols = append(freeCols, col)
			}
		}

		fmt.Printf("Pivot Columns: %v\n", pivotCols)
		fmt.Printf("Free Columns: %v\n", freeCols)

		numFree := len(freeCols)
		minPresses := B + 1
		// Enumerate all possible solutions
		for i := 0; i < (1 << numFree); i++ {
			x := make([]int, B)
			presses := 0
			for j, freeCol := range freeCols {
				if ((i >> j) & 1) == 1 {
					x[freeCol] = 1
					presses++
				} else {
					x[freeCol] = 0
				}
			}
			
			if presses >= minPresses {
				continue
			}

			for r, pivotCol := range pivotCols {
				desiredSolution := augmentedMatrix[r][B]

				// subtract contributions from free variables
				for _, freeCol := range freeCols {
					if augmentedMatrix[r][freeCol] == 1 {
						desiredSolution ^= x[freeCol]
					}
				}
				x[pivotCol] = desiredSolution
				if x[pivotCol] == 1 {
					presses++
					if presses >= minPresses {
						break
					}
				}
			}

			if presses < minPresses {
				minPresses = presses
			}
		}
		total += minPresses
	}
	return total
}

func day10Part2(lines []string) int {
	return 0
}
