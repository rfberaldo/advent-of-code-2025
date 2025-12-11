package main

import (
	"aoc2025/assert"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day03-2/input.txt")
	assert.NoErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	const max = 12
	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		lastIdx := -1
		var digits []byte

		for len(digits) < max {
			var digit byte
			// fmt.Print("init:", lastIdx, ", max:", len(line)-max+len(digits))
			for i := lastIdx + 1; i <= len(line)-max+len(digits); i++ {
				if line[i] > digit {
					digit = line[i]
					lastIdx = i
				}
			}
			// fmt.Println(", found:", string(digit))
			digits = append(digits, digit)
		}

		n, err := strconv.Atoi(string(digits))
		assert.NoErr(err)
		count += n
	}

	fmt.Println("Result:", count)
}
