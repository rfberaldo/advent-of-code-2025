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
	file, err := os.Open("day02-1/input.txt")
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

		for v := init; v <= end; v++ {
			s := strconv.Itoa(v)
			if len(s)%2 != 0 {
				// must not be odd length
				continue
			}

			s1 := s[:len(s)/2]
			s2 := s[len(s)/2:]

			if s1 == s2 {
				count += v
			}
		}
	}

	fmt.Println("Result:", count)
}
