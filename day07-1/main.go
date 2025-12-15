package main

import (
	"aoc2025/assert"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("day07-1/input.txt")
	assert.NoErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var grid [][]string

	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		grid = append(grid, chars)
	}

	type Point struct {
		x, y int
	}

	// find start point
	start := Point{}
	for i, v := range grid[0] {
		if v == "S" {
			start.x = i
			break
		}
	}

	queue := []Point{start}
	result := 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.y+1 >= len(grid) {
			continue
		}

		switch grid[curr.y][curr.x] {
		case "S", ".":
			queue = append(queue, Point{curr.x, curr.y + 1})
			grid[curr.y][curr.x] = "|"

		case "^":
			queue = append(queue, Point{curr.x - 1, curr.y + 1})
			queue = append(queue, Point{curr.x + 1, curr.y + 1})
			result++
		}

		// fmt.Println(curr)
		// for _, v := range grid {
		// 	fmt.Println(v)
		// }
		// fmt.Print("\n\n")
	}

	fmt.Println("Result:", result)
}
