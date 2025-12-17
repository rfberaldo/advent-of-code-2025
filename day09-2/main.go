package main

import (
	"aoc2025/assert"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	file, err := os.Open("day09-2/input.txt")
	assert.NoErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var points []Point

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		x, err := strconv.Atoi(fields[0])
		assert.NoErr(err)
		y, err := strconv.Atoi(fields[1])
		assert.NoErr(err)
		points = append(points, Point{x, y})
	}

	// add first point at the end to wrap
	points = append(points, points[0])

	maxArea := 0
	for i := 0; i < len(points)-1; i++ {
		p1 := &points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := &points[j]
			// shrink rect by 1 for easier intersection detection
			rect := []Point{
				// top left
				{min(p1.x, p2.x) + 1, min(p1.y, p2.y) + 1},
				// top right
				{max(p1.x, p2.x) - 1, min(p1.y, p2.y) + 1},
				// bottom right
				{max(p1.x, p2.x) - 1, max(p1.y, p2.y) - 1},
				// bottom left
				{min(p1.x, p2.x) + 1, max(p1.y, p2.y) - 1},
			}
			// top left again to wrap
			rect = append(rect, rect[0])

			if !isInside(rect, points) {
				continue
			}

			a := max(p1.x, p2.x) - min(p1.x, p2.x) + 1
			b := max(p1.y, p2.y) - min(p1.y, p2.y) + 1
			area := a * b
			if area > maxArea {
				maxArea = area
			}
		}
	}

	fmt.Println("Result:", maxArea)
}

func isInside(rect []Point, polygon []Point) bool {
	assert.True(len(rect) == 5)

	for i := 0; i < len(rect)-1; i++ {
		x1 := rect[i].x
		y1 := rect[i].y
		x2 := rect[i+1].x
		y2 := rect[i+1].y
		for j := 0; j < len(polygon)-1; j++ {
			x3 := polygon[j].x
			y3 := polygon[j].y
			x4 := polygon[j+1].x
			y4 := polygon[j+1].y
			if hasIntersection(x1, y1, x2, y2, x3, y3, x4, y4) {
				return false
			}
		}
	}

	return true
}

// https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection
func hasIntersection(x1, y1, x2, y2, x3, y3, x4, y4 int) bool {
	d := float64((x1-x2)*(y3-y4) - (y1-y2)*(x3-x4))
	if d == 0 {
		return false
	}

	t := float64((x1-x3)*(y3-y4) - (y1-y3)*(x3-x4))
	t /= d

	u := -float64((x1-x2)*(y1-y3) - (y1-y2)*(x1-x3))
	u /= d

	return 0 <= t && t <= 1 && 0 <= u && u <= 1
}
