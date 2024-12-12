package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const filepath = "input.txt"

func parseInput() []string {
	file, _ := os.Open(filepath)

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	return lines
}

func isValid(position [2]int, match byte, grid []string) bool {
	if position[0] < 0 || position[0] >= len(grid[0]) || position[1] < 0 || position[1] >= len(grid) {
		return false
	}

	return grid[position[1]][position[0]] == match
}

func corners(neighbors map[[2]int]bool) int {
	/*
		There are two types of corners to check: convex and concave

		X X	  This is a convex corner, evaluated by checking the above and left neighbor
		X A   If they are both not valid, this is a corner. Repeat for all 4 corners

		X A   This is a concave corner, check the above, left and diagonal neighbors.
		A A   If the above and left neighbor are valid, but the diagonal is not then
			  it is considered a corner
	*/

	count := 0
	cornerChecks := [][3][2]int{
		{{-1, 0}, {0, 1}, {-1, 1}},
		{{0, 1}, {1, 0}, {1, 1}},
		{{1, 0}, {0, -1}, {1, -1}},
		{{-1, 0}, {0, -1}, {-1, -1}},
	}

	for _, check := range cornerChecks {

		if !neighbors[check[0]] && !neighbors[check[1]] {
			// Convex
			count++
		} else if neighbors[check[0]] && neighbors[check[1]] && !neighbors[check[2]] {
			// Concave
			count++
		}
	}
	return count
}

func findRegion(start [2]int, grid []string, seen map[[2]int]bool) (int, int, int) {
	// Diagonal directions are just used for corner checks, don't add to stack
	directions := map[[2]int]bool{
		{0, 1}: true, {1, 0}: true, {-1, 0}: true, {0, -1}: true,
		{1, -1}: false, {-1, -1}: false, {1, 1}: false, {-1, 1}: false,
	}

	var region [][2]int
	perimeter := 0
	sides := 0
	plant := grid[start[1]][start[0]]

	stack := [][2]int{start}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if seen[current] {
			continue
		}

		region = append(region, current)
		seen[current] = true
		neighbors := make(map[[2]int]bool)
		for direction, addToStack := range directions {
			neighbor := [2]int{current[0] + direction[0], current[1] + direction[1]}
			valid := isValid(neighbor, plant, grid)
			neighbors[direction] = valid
			if addToStack {
				if valid {
					stack = append(stack, neighbor)
				} else {
					perimeter++
				}
			}
		}
		sides += corners(neighbors)
	}

	return len(region), perimeter, sides
}

func part1() int {
	grid := parseInput()
	seen := make(map[[2]int]bool)
	total := 0
	for y := range len(grid) {
		for x := range len(grid[y]) {
			position := [2]int{x, y}
			if !seen[position] {
				area, perimeter, _ := findRegion(position, grid, seen)
				total += area * perimeter
			}
		}
	}

	return total
}

func part2() int {
	grid := parseInput()
	seen := make(map[[2]int]bool)
	total := 0
	for y := range len(grid) {
		for x := range len(grid[y]) {
			position := [2]int{x, y}
			if !seen[position] {
				area, _, sides := findRegion(position, grid, seen)
				total += area * sides
			}
		}
	}

	return total
}

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %v in %v\n", part1(), time.Since(start))
	half := time.Now()
	fmt.Printf("Part 2: %v in %v\n", part2(), time.Since(half))
}
