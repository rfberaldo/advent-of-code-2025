package main

import (
	"aoc2025/assert"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("day11-1/input.txt")
	assert.NoErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	nodes := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()
		label := line[:3]
		fields := strings.Fields(line[5:])
		var out []string
		for _, f := range fields {
			out = append(out, f)
		}
		nodes[label] = out
	}

	fmt.Println("Result:", walk("you", nodes))
}

func walk(label string, nodes map[string][]string) int {
	if label == "out" {
		return 1
	}

	count := 0
	for _, v := range nodes[label] {
		count += walk(v, nodes)
	}

	return count
}
