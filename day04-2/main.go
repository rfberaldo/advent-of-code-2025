package main

import (
	"aoc2025/assert"
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("day04-2/input.txt")
	assert.NoErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	result := 0

	var grid [][]string

	for scanner.Scan() {
		line := scanner.Text()
		var row []string
		for i := range line {
			row = append(row, string(line[i]))
		}
		grid = append(grid, row)
	}

	grid = addGridMargin(grid, ".")

	type Point struct {
		x, y int
	}

	adj := []Point{
		{-1, -1}, // up left
		{0, -1},  // up
		{+1, -1}, // up right
		{+1, 0},  // right
		{+1, +1}, // down right
		{0, +1},  // down
		{-1, +1}, // down left
		{-1, 0},  // left
	}

	for {
		var toRemove []Point

		for y := 1; y < len(grid)-1; y++ {
			for x := 1; x < len(grid[y])-1; x++ {
				if grid[y][x] != "@" {
					continue
				}

				count := 0
				for _, p := range adj {
					if grid[y+p.y][x+p.x] == "@" {
						count++
					}
				}

				if count < 4 {
					result++
					toRemove = append(toRemove, Point{x, y})
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		for _, p := range toRemove {
			grid[p.y][p.x] = "."
		}
	}

	fmt.Println("Result:", result)
}

func addGridMargin(grid [][]string, s string) [][]string {
	var row []string
	for range len(grid[0]) {
		row = append(row, s)
	}

	grid = append([][]string{row}, grid...)
	grid = append(grid, row)

	for i := range grid {
		grid[i] = append([]string{s}, grid[i]...)
		grid[i] = append(grid[i], s)
	}

	return grid
}
