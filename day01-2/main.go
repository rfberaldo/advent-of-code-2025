package main

import (
	"aoc2025/assert"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day01-2/input.txt")
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

		for range dist {
			switch dir {
			case "L":
				pos--
			case "R":
				pos++
			}

			if pos < 0 {
				pos = 99
			}
			if pos > 99 {
				pos = 0
			}

			if pos == 0 {
				count++
			}
		}

		assert.True(pos >= 0 && pos <= 99)

		fmt.Println(pos)
	}

	fmt.Println("Result:", count)
}
