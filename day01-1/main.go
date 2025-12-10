package main

import (
	"aoc2025/assert"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day01-1/input.txt")
	assert.NoErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	pos := 50
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		dir := string(line[0])
		dist, err := strconv.Atoi(line[1:])
		assert.NoErr(err)

		if dir == "R" {
			pos += dist
		}

		if dir == "L" {
			pos -= dist
		}

		pos = pos - (pos / 100 * 100)

		if pos < 0 {
			pos = 100 + pos
		}

		if pos == 0 {
			count++
		}
	}

	fmt.Println("Result:", count)
}
