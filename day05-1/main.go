package main

import (
	"aoc2025/assert"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day05-1/input.txt")
	assert.NoErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	result := 0

	type Range struct {
		min, max int
	}

	var ranges []Range

	readingRange := true
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			readingRange = false
			continue
		}

		if readingRange {
			minStr, maxStr, ok := strings.Cut(line, "-")
			assert.True(ok)
			min, err := strconv.Atoi(minStr)
			assert.NoErr(err)
			max, err := strconv.Atoi(maxStr)
			assert.NoErr(err)
			ranges = append(ranges, Range{min, max})
			continue
		}

		id, err := strconv.Atoi(line)
		assert.NoErr(err)

		for _, r := range ranges {
			if r.min <= id && id <= r.max {
				result++
				break
			}
		}
	}

	fmt.Println("Result:", result)
}
