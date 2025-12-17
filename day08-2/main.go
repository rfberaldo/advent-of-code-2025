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

type Point struct {
	x, y, z, circ int
}

type Pair struct {
	p1, p2 *Point
	dist   float64
}

func main() {
	file, err := os.Open("day08-2/input.txt")
	assert.NoErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var points []Point

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		x, err := strconv.Atoi(fields[0])
		assert.NoErr(err)
		y, err := strconv.Atoi(fields[1])
		assert.NoErr(err)
		z, err := strconv.Atoi(fields[2])
		assert.NoErr(err)
		points = append(points, Point{x, y, z, 0})
	}

	var pairs []Pair
	for i := 0; i < len(points)-1; i++ {
		p1 := &points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := &points[j]
			pairs = append(pairs, Pair{p1, p2, euclideanDist(*p1, *p2)})
		}
	}

	// total pairs = N*(N-1)/2
	assert.True(len(pairs) == len(points)*(len(points)-1)/2)

	// sort by dist
	slices.SortFunc(pairs, func(p1, p2 Pair) int {
		return cmp.Compare(p1.dist, p2.dist)
	})

	pointsByCircuit := make(map[int]int)

	for i := range pairs {
		pair := &pairs[i]
		circ1 := pair.p1.circ
		circ2 := pair.p2.circ

		// same circuit
		if circ1 > 0 && circ1 == circ2 {
			continue
		}

		// merge two circuits
		if circ1 > 0 && circ2 > 0 {
			for j := range points {
				if points[j].circ == circ2 {
					points[j].circ = circ1
				}
			}

			pointsByCircuit[circ1] += pointsByCircuit[circ2]
			delete(pointsByCircuit, circ2)

		} else {

			// simple connection
			circ := max(circ1, circ2)
			if circ == 0 {
				circ = i + 1
			}
			pair.p1.circ = circ
			pair.p2.circ = circ

			if circ1 != circ {
				pointsByCircuit[circ]++
			}
			if circ2 != circ {
				pointsByCircuit[circ]++
			}
		}

		if len(pointsByCircuit) == 1 && getMapValue(pointsByCircuit) == len(points) {
			fmt.Println("Result:", pair.p1.x*pair.p2.x)
			break
		}
	}
}

func getMapValue[K comparable, V any](m map[K]V) V {
	assert.True(len(m) == 1)
	for _, v := range m {
		return v
	}
	panic("unreacheable")
}

func euclideanDist(p1, p2 Point) float64 {
	x := float64(p2.x - p1.x)
	y := float64(p2.y - p1.y)
	z := float64(p2.z - p1.z)
	return math.Sqrt(x*x + y*y + z*z)
}
