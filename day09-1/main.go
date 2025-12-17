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
	file, err := os.Open("day09-1/input.txt")
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

	maxArea := 0
	for i := 0; i < len(points)-1; i++ {
		p1 := &points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := &points[j]
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
