package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
	"strings"
)

const (
	cellSize   = 10 // pixels per cell
	frameDelay = 5 // delay between frames in 100ths of a second (50 = 0.5s)
)

// Colors for visualization
var (
	colorEmpty     = color.RGBA{240, 240, 240, 255} // light gray for empty
	colorRoll      = color.RGBA{60, 60, 60, 255}    // dark gray for remaining rolls
	colorRemoving  = color.RGBA{255, 80, 80, 255}   // red for rolls being removed
	colorGrid      = color.RGBA{200, 200, 200, 255} // grid lines
)

func day4Visualize() {
	// Read input file
	data, err := fetchInput(4)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input := string(data)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Generate the animated GIF
	if err := generateAnimatedGIF(lines, "day4_animation.gif"); err != nil {
		fmt.Println("Error generating GIF:", err)
		return
	}

	fmt.Println("✓ Animation saved to day4_animation.gif")
}

func generateAnimatedGIF(lines []string, filename string) error {
	grid := parseGrid(lines)
	rows := len(grid)
	cols := len(grid[0])

	// Image dimensions
	width := cols * cellSize
	height := rows * cellSize

	var images []*image.Paletted
	var delays []int

	// Create color palette
	palette := color.Palette{
		colorEmpty,
		colorRoll,
		colorRemoving,
		colorGrid,
	}

	iteration := 0

	// Initial state frame
	img := createFrame(grid, width, height, palette, 0, 0)
	images = append(images, img)
	delays = append(delays, frameDelay*2) // Hold initial state longer

	// Process each iteration
	for {
		removed := 0
		toRemove := make([][]bool, rows)
		for i := range toRemove {
			toRemove[i] = make([]bool, cols)
		}

		// Find all cells to remove this iteration
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] == '@' && isAccessible(grid, i, j) {
					toRemove[i][j] = true
					removed++
				}
			}
		}

		if removed == 0 {
			break
		}

		iteration++

		// Frame showing cells about to be removed (highlighted in red)
		img = createFrameWithHighlight(grid, toRemove, width, height, palette, iteration, removed)
		images = append(images, img)
		delays = append(delays, frameDelay)

		// Actually remove the cells
		for i := range grid {
			for j := range grid[i] {
				if toRemove[i][j] {
					grid[i][j] = '.'
				}
			}
		}

		// Frame showing state after removal
		img = createFrame(grid, width, height, palette, iteration, removed)
		images = append(images, img)
		delays = append(delays, frameDelay)
	}

	// Final state frame (hold longer)
	img = createFrame(grid, width, height, palette, iteration, 0)
	images = append(images, img)
	delays = append(delays, frameDelay*3)

	// Save GIF
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}

func createFrame(grid [][]byte, width, height int, palette color.Palette, iteration, removed int) *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, width, height), palette)

	// Fill background
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, colorEmpty)
		}
	}

	// Draw cells
	for i := range grid {
		for j := range grid[i] {
			drawCell(img, i, j, grid[i][j], false)
		}
	}

	return img
}

func createFrameWithHighlight(grid [][]byte, toRemove [][]bool, width, height int, palette color.Palette, iteration, removed int) *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, width, height), palette)

	// Fill background
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, colorEmpty)
		}
	}

	// Draw cells with highlighting
	for i := range grid {
		for j := range grid[i] {
			highlight := toRemove[i][j]
			drawCell(img, i, j, grid[i][j], highlight)
		}
	}

	return img
}

func drawCell(img *image.Paletted, row, col int, cell byte, highlight bool) {
	x0 := col * cellSize
	y0 := row * cellSize
	x1 := x0 + cellSize
	y1 := y0 + cellSize

	var fillColor color.Color
	if cell == '@' {
		if highlight {
			fillColor = colorRemoving
		} else {
			fillColor = colorRoll
		}
	} else {
		fillColor = colorEmpty
	}

	// Fill cell
	for y := y0; y < y1; y++ {
		for x := x0; x < x1; x++ {
			img.Set(x, y, fillColor)
		}
	}

	// Draw grid lines (1px border)
	for x := x0; x < x1; x++ {
		img.Set(x, y0, colorGrid)
		img.Set(x, y1-1, colorGrid)
	}
	for y := y0; y < y1; y++ {
		img.Set(x0, y, colorGrid)
		img.Set(x1-1, y, colorGrid)
	}
}
