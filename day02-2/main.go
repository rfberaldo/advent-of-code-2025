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
	file, err := os.Open("day02-2/input.txt")
	assert.NoErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan() // only 1 line
	line := scanner.Text()

	count := 0

	for strrange := range strings.SplitSeq(line, ",") {
		ids := strings.Split(strrange, "-")
		assert.True(len(ids) == 2)

		init, err := strconv.Atoi(ids[0])
		assert.NoErr(err)
		end, err := strconv.Atoi(ids[1])
		assert.NoErr(err)

		// fmt.Println(init, end)

	outer:
		for v := init; v <= end; v++ {
			s := strconv.Itoa(v)
			for i := 1; i <= len(s)/2; i++ {
				if strings.Count(s, s[:i])*i == len(s) {
					count += v
					continue outer
				}
			}
		}
	}

	fmt.Println("Result:", count)
}
