package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const filepath = "input.txt"

func parseInput() []string {
	file, ferr := os.Open(filepath)
	check(ferr)

	scanner := bufio.NewScanner(file)
	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	file.Close()

	return grid
}

func isValid(position [2]int, grid []string) bool {
	if position[0] < 0 || position[0] >= len(grid) {
		return false
	}
	if position[1] < 0 || position[1] >= len(grid[0]) {
		return false
	}

	return true
}

func xmasScan(grid []string, position [2]int) int {
	directions := [][]int{
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}

	goal := "XMAS"
	total := 0
	// This should always be an X
	start := string(grid[position[1]][position[0]])
	for _, direction := range directions {
		newPosition := [2]int{position[0] + direction[0], position[1] + direction[1]}
		word := start
		for len(word) < 4 && goal[:len(word)] == word && isValid(newPosition, grid) {
			word += string(grid[newPosition[1]][newPosition[0]])
			if word == goal {
				total++
				break
			}

			newPosition = [2]int{newPosition[0] + direction[0], newPosition[1] + direction[1]}
		}
	}

	return total
}

func x_masScan(grid []string, position [2]int) int {
	cornerPairs := [][][2]int{
		{{position[0] - 1, position[1] + 1}, {position[0] + 1, position[1] - 1}},
		{{position[0] - 1, position[1] - 1}, {position[0] + 1, position[1] + 1}},
	}

	// If "M" (77) is in one corner, "S" (83) needs to be in the opposite corner
	match := map[byte]byte{77: 83, 83: 77}

	valid := true
	for _, cornerPair := range cornerPairs {
		if !isValid(cornerPair[0], grid) || !isValid(cornerPair[1], grid) {
			valid = false
			break
		}

		corner1Value := grid[cornerPair[0][1]][cornerPair[0][0]]
		corner2Value := grid[cornerPair[1][1]][cornerPair[1][0]]
		if match[corner1Value] != corner2Value {
			valid = false
			break
		}
	}

	if valid {
		return 1
	}

	return 0
}

func part1() int {
	wordSearch := parseInput()
	xmasCount := 0
	for y, row := range wordSearch {
		for x, char := range row {
			// X in byte
			if char == 88 {
				xmasCount += xmasScan(wordSearch, [2]int{x, y})
			}
		}
	}

	return xmasCount
}

func part2() int {
	wordSearch := parseInput()
	x_masCount := 0
	for y, row := range wordSearch {
		for x, char := range row {
			// A in byte
			if char == 65 {
				x_masCount += x_masScan(wordSearch, [2]int{x, y})
			}
		}
	}

	return x_masCount
}

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %v in %v\n", part1(), time.Since(start))
	half := time.Now()
	fmt.Printf("Part 2: %v in %v\n", part2(), time.Since(half))
}
