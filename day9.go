package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
)

func day9() {
	// Read input file
	data, err := fetchInput(9)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input := string(data)
	lines := strings.Split(input, "\n")

	// Part 1
	result1 := day9Part1(lines)
	fmt.Printf("Part 1: %d\n", result1)

	// Part 2
	result2 := day9Part2(lines)
	fmt.Printf("Part 2: %d\n", result2)
}

func area2D(p1, p2 Point) int {
	dx := int(math.Abs(float64(p1.x - p2.x))) + 1
	dy := int(math.Abs(float64(p1.y - p2.y))) + 1
	return dx * dy
}

func parse2DPoints(lines []string) []Point {
	points := []Point{}
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			fmt.Println("Error in input")
			continue
		}
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}

		points = append(points, Point{x, y})
	}
	return points
}

func day9Part1(lines []string) int {
	points := parse2DPoints(lines)
	maxArea := 0
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			area := area2D(points[i], points[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func isRectangleValid(p1, p2 Point, polygon []Point) bool {
    // Phase 1: Check all 4 corners are inside or on polygon
    minX, maxX := min(p1.x, p2.x), max(p1.x, p2.x)
    minY, maxY := min(p1.y, p2.y), max(p1.y, p2.y)

	// Check no polygon vertex is strictly inside rectangle
	for _, point := range polygon {
        if point.x > minX && point.x < maxX && point.y > minY && point.y < maxY {
            return false
        }
    }

    corners := []Point{
        {minX, minY}, {minX, maxY},
        {maxX, minY}, {maxX, maxY},
    }

	// Check all corners are inside or on polygon
    for _, corner := range corners {
        if !isInsideOrOnPolygon(corner, polygon) {
            return false
        }
    }

    // Check if any polygon edge properly crosses the boundary of the rectangle
    rectEdges := []struct{ p1, p2 Point }{
        {p1: Point{minX, minY}, p2: Point{maxX, minY}},
        {p1: Point{maxX, minY}, p2: Point{maxX, maxY}},
        {p1: Point{maxX, maxY}, p2: Point{minX, maxY}},
        {p1: Point{minX, maxY}, p2: Point{minX, minY}},
    }

    n := len(polygon)
    for i := 0; i < n; i++ {
        j := (i + 1) % n
        polyP1, polyP2 := polygon[i], polygon[j]

        for _, rectEdge := range rectEdges {
            if segmentsProperlyIntersect(rectEdge.p1, rectEdge.p2, polyP1, polyP2) {
                return false
            }
        }
    }

    return true
}

// Check if two segments properly intersect (cross each other, not just touch)
func segmentsProperlyIntersect(p1, q1, p2, q2 Point) bool {
    o1 := orientation(p1, q1, p2)
    o2 := orientation(p1, q1, q2)
    o3 := orientation(p2, q2, p1)
    o4 := orientation(p2, q2, q1)

    return (o1 != o2 && o3 != o4) && !(o1 == 0 || o2 == 0 || o3 == 0 || o4 == 0)
}

func orientation(p, q, r Point) int {
    val := (q.y-p.y)*(r.x-q.x) - (q.x-p.x)*(r.y-q.y)
    if val == 0 {
        return 0 // collinear
    }
    if val > 0 {
        return 1 // clockwise
    }
    return 2 // counterclockwise
}

func isOnSegment(point, p1, p2 Point) bool {
    // Check collinearity using cross product
    // (p2 - p1) × (point - p1) should be 0
    cross := (p2.y-p1.y)*(point.x-p1.x) - (p2.x-p1.x)*(point.y-p1.y)
    return cross == 0
}

func isInsideOrOnPolygon(point Point, polygon []Point) bool {
    count := 0
    n := len(polygon)
    
    for i := 0; i < n; i++ {
        j := (i + 1) % n
        p1, p2 := polygon[i], polygon[j]

		// On Polygon Edge
		if isOnSegment(point, p1, p2) {
            return true
        }

		// Inside Polygon
		// raycasting algorithm
        // Ray going right (positive x direction)
        // Check if edge crosses the horizontal ray
        if (p1.y > point.y) != (p2.y > point.y) {
            // Calculate x coordinate of intersection
            xIntersect := p1.x + (p2.x-p1.x)*(point.y-p1.y)/(p2.y-p1.y)
            if xIntersect > point.x {
                count++
            }
        }
    }
    
    // Point is inside if count is odd
    return count%2 == 1
}

func day9Part2(lines []string) int {
	points := parse2DPoints(lines)
	maxArea := 0
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			area := area2D(points[i], points[j])
			if area > maxArea && isRectangleValid(points[i], points[j], points) {
				maxArea = area
			}
		}
	}
	return maxArea
}
