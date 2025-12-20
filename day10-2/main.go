package main

import (
	"aoc2025/assert"
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Machine struct {
	joltage []int
	btns    [][]int
}

// press returns the new joltage, returns false if any digit is < zero
func (m *Machine) press(pattern []int) ([]int, bool) {
	joltage := slices.Clone(m.joltage)
	for i, n := range pattern {
		joltage[i] -= n
		if joltage[i] < 0 {
			return joltage, false
		}
	}
	return joltage, true
}

// done returns whether joltage is all zero
func (m *Machine) done() bool {
	for _, n := range m.joltage {
		if n != 0 {
			return false
		}
	}
	return true
}

func (m *Machine) pattern(btns [][]int) []int {
	pattern := make([]int, len(m.joltage))
	for _, btn := range btns {
		for _, n := range btn {
			pattern[n]++
		}
	}
	return pattern
}

// https://en.wikipedia.org/wiki/Power_set
func (m *Machine) powersets() [][][]int {
	// total powerset = 2^N
	total := int(math.Pow(2, float64(len(m.btns))))
	powersets := make([][][]int, 0, total)

	for i := 0; i < (1 << len(m.btns)); i++ {
		var btns [][]int
		for j, btn := range m.btns {
			if i&(1<<j) > 0 {
				btns = append(btns, btn)
			}
		}
		powersets = append(powersets, btns)
	}

	assert.True(len(powersets) == total)

	// sort by length (cost)
	slices.SortFunc(powersets, func(a, b [][]int) int {
		return cmp.Compare(len(a), len(b))
	})

	return powersets
}

func (m *Machine) key() string {
	return fmt.Sprintf("%v", *m)
}

type Pattern struct {
	pattern []int
	cost    int
}

// https://www.reddit.com/r/adventofcode/comments/1pk87hl/2025_day_10_part_2_bifurcate_your_way_to_victory/
func (m *Machine) solve() int {
	var patterns []Pattern
	for _, btns := range m.powersets() {
		pattern := m.pattern(btns)

		// discard same pattern with higher cost
		if slices.ContainsFunc(patterns, func(p Pattern) bool {
			return slices.Equal(p.pattern, pattern)
		}) {
			continue
		}

		patterns = append(patterns, Pattern{pattern, len(btns)})
	}

	res := solve(*m, patterns)
	assert.True(slices.Max(m.joltage) <= res && res < 1_000_000, "machine: ", *m)
	return res
}

var cacheSolve = make(map[string]int)

func solve(m Machine, patterns []Pattern) int {
	key := m.key()
	if n, ok := cacheSolve[key]; ok {
		return n
	}

	if m.done() {
		return 0
	}

	count := 1_000_000
	for _, pat := range patterns {
		joltage, ok := m.press(pat.pattern)
		if !ok || hasAnyOdd(joltage) {
			continue
		}

		m2 := m
		m2.joltage = divide(joltage, 2)
		count = min(count, pat.cost+2*solve(m2, patterns))
	}

	cacheSolve[key] = count
	return count
}

func hasAnyOdd(s []int) bool {
	for _, n := range s {
		if n%2 != 0 {
			return true
		}
	}
	return false
}

func divide(s []int, d int) []int {
	if d == 1 {
		return s
	}

	s2 := make([]int, len(s))
	for i := range s {
		s2[i] = s[i] / d
	}
	return s2
}

func main() {
	file, err := os.Open("day10-2/input.txt")
	assert.NoErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var machines []Machine

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		indic := strings.Split(strings.Trim(fields[len(fields)-1], "{}"), ",")
		joltage := make([]int, len(indic))
		for i, s := range indic {
			n, err := strconv.Atoi(string(s))
			assert.NoErr(err)
			joltage[i] = n
		}
		var btns [][]int
		for _, f := range fields[1 : len(fields)-1] {
			var btn []int
			f = strings.Trim(f, "()")
			for s := range strings.SplitSeq(f, ",") {
				n, err := strconv.Atoi(s)
				assert.NoErr(err)
				btn = append(btn, n)
			}
			btns = append(btns, btn)
		}

		machines = append(machines, Machine{
			joltage: joltage,
			btns:    btns,
		})
	}

	result := 0

	for _, m := range machines {
		result += m.solve()
	}

	fmt.Println("Result:", result)
}
