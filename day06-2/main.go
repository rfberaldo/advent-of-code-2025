package main

import (
	"aoc2025/assert"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// For this to work, input file must be saved without formatting,
// meaning it must keep trailing whitespace.

func main() {
	file, err := os.Open("day06-2/input.txt")
	assert.NoErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	type Operation struct {
		num []int
		op  string
	}

	var ops []Operation
	result := 0
	op := Operation{}

	for x := len(lines[0]) - 1; x >= 0; x-- {
		var val string
		var done = false
		for y := 0; y < len(lines); y++ {
			v := string(lines[y][x])
			if v == "*" || v == "+" {
				op.op = v
				done = true
				break
			}
			val += v
		}

		val = strings.TrimSpace(val)
		num, err := strconv.Atoi(val)
		assert.NoErr(err)
		op.num = append(op.num, num)

		if done {
			ops = append(ops, op)
			op = Operation{}
			done = false
			x--
		}
	}
	// fmt.Printf("%+v\n", ops)

	for _, op := range ops {
		r := op.num[0]
		for i := 1; i < len(op.num); i++ {
			v := op.num[i]
			switch op.op {
			case "+":
				r += v
			case "*":
				r *= v
			default:
				panic("unreachable")
			}
		}
		result += r
	}

	fmt.Println("Result:", result)
}
