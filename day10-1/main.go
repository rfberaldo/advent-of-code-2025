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
	onState []int
	btns    [][]int
	state   []int
}

func (m *Machine) reset() {
	m.state = make([]int, len(m.onState))
}

func (m *Machine) toggle(n int) {
	m.state[n] ^= 1
}

func (m *Machine) press(btns [][]int) {
	for _, btn := range btns {
		for _, n := range btn {
			m.toggle(n)
		}
	}
}

func (m *Machine) isOn() bool {
	return slices.Equal(m.state, m.onState)
}

// https://en.wikipedia.org/wiki/Power_set
func (m *Machine) powerset() [][][]int {
	// total powerset = 2^N - 1
	total := int(math.Pow(2, float64(len(m.btns)))) - 1
	powerset := make([][][]int, 0, total)

	for i := 0; i < (1 << len(m.btns)); i++ {
		var btns [][]int
		for j, btn := range m.btns {
			if i&(1<<j) > 0 {
				btns = append(btns, btn)
			}
		}
		if len(btns) > 0 {
			powerset = append(powerset, btns)
		}
	}

	assert.True(len(powerset) == total)

	// sort by length
	slices.SortFunc(powerset, func(a, b [][]int) int {
		return cmp.Compare(len(a), len(b))
	})

	return powerset
}

func main() {
	file, err := os.Open("day10-1/input.txt")
	assert.NoErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var machines []Machine

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		indic := strings.Trim(fields[0], "[]")
		onState := make([]int, len(indic))
		for i, s := range indic {
			n := 0
			if s == '#' {
				n = 1
			}
			onState[i] = n
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
		machines = append(machines, Machine{onState: onState, btns: btns})
	}

	result := 0

	for _, m := range machines {
		for _, btns := range m.powerset() {
			m.reset()
			m.press(btns)
			if m.isOn() {
				result += len(btns)
				break
			}
		}
	}

	fmt.Println("Result:", result)
}
