package main

import (
	"aoc2025/assert"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day03-1/input.txt")
	assert.NoErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		b1 := line[0]
		i1 := 0

		// bigger number that is not the last digit
		for i := 0; i < len(line); i++ {
			b := line[i]
			if b > b1 && i < len(line)-1 {
				i1 = i
				b1 = b
			}
		}

		b2 := line[i1+1]

		// bigger number after b1
		for i := i1 + 1; i < len(line); i++ {
			b := line[i]
			if b > b2 {
				b2 = b
			}
		}

		// fmt.Println(string(b1), string(b2))

		n, err := strconv.Atoi(string(b1) + string(b2))
		assert.NoErr(err)
		count += n
	}

	fmt.Println("Result:", count)
}
