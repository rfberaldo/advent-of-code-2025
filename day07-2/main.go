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

func (p *Point) key() string {
	return strconv.Itoa(p.x) + "-" + strconv.Itoa(p.y)
}

var grid [][]string
var memo map[string]int

func main() {
	file, err := os.Open("day07-2/input.txt")
	assert.NoErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		grid = append(grid, chars)
	}

	// find start point
	start := Point{}
	for i, v := range grid[0] {
		if v == "S" {
			start.x = i
			break
		}
	}

	memo = make(map[string]int)

	fmt.Println("Result:", walk(start, 0))
}

func walk(curr Point, r int) int {
	if v, ok := memo[curr.key()]; ok {
		return r + v
	}

	if curr.y > len(grid)-1 {
		return r + 1
	}

	switch grid[curr.y][curr.x] {
	case "S", ".":
		return walk(Point{curr.x, curr.y + 2}, r)

	case "^":
		v := walk(Point{curr.x - 1, curr.y + 2}, r) +
			walk(Point{curr.x + 1, curr.y + 2}, r)
		memo[curr.key()] = v
		return v
	}

	panic("unreacheable")
}
