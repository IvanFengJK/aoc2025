package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <day>")
		fmt.Println("Example: go run . 12")
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid day number: %s\n", os.Args[1])
		os.Exit(1)
	}

	switch day {
	case 1:
		day1()
	case 2:
		day2()
	case 3:
		day3()
	case 4:
		day4()
	case 5:
		day5()
	case 6:
		day6()
	case 7:
		day7()
	case 8:
		day8()
	case 9:
		day9()
	case 10:
		day10()
	case 11:
		day11()
	case 12:
		day12()
	default:
		fmt.Printf("Day %d not implemented yet\n", day)
		os.Exit(1)
	}
}
