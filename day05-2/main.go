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
	file, err := os.Open("day05-2/input.txt")
	assert.NoErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	result := 0

	type Range struct {
		min, max int
	}

	var ranges []*Range

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		minStr, maxStr, ok := strings.Cut(line, "-")
		assert.True(ok)
		min, err := strconv.Atoi(minStr)
		assert.NoErr(err)
		max, err := strconv.Atoi(maxStr)
		assert.NoErr(err)
		ranges = append(ranges, &Range{min, max})
	}

	for {
		overlaps := false

		for i := 0; i < len(ranges); i++ {
			r1 := ranges[i]
			if r1 == nil {
				continue
			}

			for j := 0; j < len(ranges); j++ {
				r2 := ranges[j]
				if r2 == nil {
					continue
				}

				if i == j {
					continue // skip same
				}

				if r1.min > r2.max || r2.min > r1.max {
					continue // no overlap
				}

				overlaps = true
				add := &Range{min(r1.min, r2.min), max(r1.max, r2.max)}
				// fmt.Print(r1, " overlaps with ", r2)
				// fmt.Println(", replacing with", add)
				ranges[i] = add
				ranges[j] = nil
				break
			}
		}

		if !overlaps {
			break
		}
	}

	for _, r := range ranges {
		if r == nil {
			continue
		}
		result += r.max + 1 - r.min
	}

	fmt.Println("Result:", result)
}
