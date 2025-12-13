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
	file, err := os.Open("day06-1/input.txt")
	assert.NoErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	result := 0

	type Operation struct {
		num []int
		op  string
	}

	opByColumn := make(map[int]Operation)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		for i, v := range fields {
			op := opByColumn[i]
			if v == "*" || v == "+" {
				op.op = v
				opByColumn[i] = op
				continue
			}
			num, err := strconv.Atoi(v)
			assert.NoErr(err)
			op.num = append(op.num, num)
			opByColumn[i] = op
		}
	}

	for _, op := range opByColumn {
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
