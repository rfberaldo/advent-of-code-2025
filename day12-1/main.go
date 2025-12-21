package main

import (
	"aoc2025/assert"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Grid struct {
	area   int
	shapes int
}

func main() {
	file, err := os.Open("day12-1/input.txt")
	assert.NoErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var grids []Grid

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "x") {
			continue
		}
		size, shapes, ok := strings.Cut(line, ":")
		assert.True(ok)
		sw, sh, ok := strings.Cut(size, "x")
		assert.True(ok)
		w, err := strconv.Atoi(sw)
		assert.NoErr(err)
		h, err := strconv.Atoi(sh)
		assert.NoErr(err)
		nums := strings.Fields(shapes)
		count := 0
		for _, s := range nums {
			n, err := strconv.Atoi(s)
			assert.NoErr(err)
			count += n
		}
		grids = append(grids, Grid{w * h, count})
	}

	result := 0
	for _, grid := range grids {
		if grid.area >= grid.shapes*9 {
			result++
		}
	}

	fmt.Println("Result:", result)
}
