package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const filepath = "input.txt"

func parseInput() ([]string, [2]int) {
	file, ferr := os.Open(filepath)
	check(ferr)

	scanner := bufio.NewScanner(file)
	var grid []string
	var guard [2]int

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		guardColumn := strings.Index(line, "^")
		if guardColumn != -1 {
			guard = [2]int{guardColumn, y}
		}
		grid = append(grid, line)
		y++
	}
	file.Close()

	return grid, guard
}

func outOfBounds(position [2]int, grid []string) bool {
	return position[0] < 0 || position[0] >= len(grid[0]) || position[1] < 0 || position[1] >= len(grid)
}

func mapRoute(guard [2]int, grid []string) (int, bool) {
	turn := map[[2]int][2]int{
		{0, -1}: {1, 0},
		{1, 0}:  {0, 1},
		{0, 1}:  {-1, 0},
		{-1, 0}: {0, -1},
	}

	direction := [2]int{0, -1}
	seen := make(map[[2]int][2]int)
	loop := false
	for !loop && !outOfBounds(guard, grid) {
		if seen[guard] == direction {
			loop = true
			break
		}

		seen[guard] = direction
		newPosition := [2]int{guard[0] + direction[0], guard[1] + direction[1]}
		for !outOfBounds(newPosition, grid) && string(grid[newPosition[1]][newPosition[0]]) == "#" {
			direction = turn[direction]
			newPosition = [2]int{guard[0] + direction[0], guard[1] + direction[1]}
		}
		guard = newPosition
	}

	return len(seen), loop
}

func part1() int {
	grid, guard := parseInput()
	seen, _ := mapRoute(guard, grid)

	return seen
}

func part2() int {
	grid, guard := parseInput()
	loops := 0
	for y, row := range grid {
		for x, char := range row {
			if string(char) != "#" {
				newGrid := make([]string, len(grid))
				copy(newGrid, grid)
				newGrid[y] = newGrid[y][:x] + "#" + newGrid[y][x+1:]
				_, loop := mapRoute(guard, newGrid)
				if loop {
					loops++
				}
			}
		}
	}

	return loops
}

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %v in %v\n", part1(), time.Since(start))
	half := time.Now()
	fmt.Printf("Part 2: %v in %v\n", part2(), time.Since(half))
}
