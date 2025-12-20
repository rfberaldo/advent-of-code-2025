package main

import (
	"aoc2025/assert"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var nodes = make(map[string][]string)

func main() {
	file, err := os.Open("day11-2/input.txt")
	assert.NoErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

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

	fmt.Println("Result:", walk("svr", 0, 0))
}

var cache = make(map[string]int)

func walk(label string, dac, fft int) int {
	key := label + strconv.Itoa(dac) + strconv.Itoa(fft)
	if n, ok := cache[key]; ok {
		return n
	}

	switch label {
	case "out":
		if dac >= 1 && fft >= 1 {
			return 1
		}
		return 0
	case "dac":
		dac++
	case "fft":
		fft++
	}

	res := 0
	for _, to := range nodes[label] {
		res += walk(to, dac, fft)
	}

	cache[key] = res
	return res
}
