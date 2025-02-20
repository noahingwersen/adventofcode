package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"time"
)

const filepath = "test.txt"

func parseInput() ([][]byte, []byte) {
	file, _ := os.ReadFile(filepath)
	parts := bytes.Split(file, []byte{'\n', '\n'})

	return bytes.Split(parts[0], []byte{'\n'}), parts[1]
}

func print(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func scoreGps(warehouse [][]byte) int {
	total := 0
	for y, row := range warehouse {
		for x, char := range row {
			if char == 'O' {
				total += x + (100 * y)
			}
		}
	}
	return total
}

func isValid(position [2]int, direction [2]int, grid [][]byte) bool {
	if grid[position[1]][position[0]] == '#' {
		return false
	}

	for slices.Contains([]byte{'O', '[', ']'}, grid[position[1]][position[0]]) {
		newPosition := [2]int{position[0] + direction[0], position[1] + direction[1]}
		if grid[newPosition[1]][newPosition[0]] == '#' {
			return false
		}
		position = newPosition
	}

	return true
}

func moveBoxes(position [2]int, direction [2]int, warehouse [][]byte) [][2]int {
	var boxLocations [][2]int
	for warehouse[position[1]][position[0]] == 'O' {
		nextPosition := [2]int{position[0] + direction[0], position[1] + direction[1]}
		boxLocations = append(boxLocations, nextPosition)
		position = nextPosition
	}
	return boxLocations
}

func isVerticalValid(position [2]int, direction [2]int, grid [][]byte) bool {
	if direction[0] != 0 {
		return false
	}
}

func part1() int {
	warehouse, moves := parseInput()
	var robot [2]int
	for y, row := range warehouse {
		for x, char := range row {
			if char == '@' {
				robot = [2]int{x, y}
			}
		}
	}

	directions := map[byte][2]int{
		'^': {0, -1},
		'>': {1, 0},
		'v': {0, 1},
		'<': {-1, 0},
	}

	for _, move := range moves {
		direction := directions[move]
		newPosition := [2]int{robot[0] + direction[0], robot[1] + direction[1]}
		if isValid(newPosition, direction, warehouse) {
			for _, position := range moveBoxes(newPosition, direction, warehouse) {
				warehouse[position[1]][position[0]] = 'O'
			}
			warehouse[robot[1]][robot[0]] = '.'
			warehouse[newPosition[1]][newPosition[0]] = '@'
			robot = newPosition
		}
	}

	return scoreGps(warehouse)
}

func doubleWarehouse(initial [][]byte) [][]byte {
	var newWarehouse [][]byte
	for _, row := range initial {
		var newRow []byte
		for _, char := range row {
			switch char {
			case '#':
				newRow = append(newRow, []byte{'#', '#'}...)
			case 'O':
				newRow = append(newRow, []byte{'[', ']'}...)
			case '.':
				newRow = append(newRow, []byte{'.', '.'}...)
			case '@':
				newRow = append(newRow, []byte{'@', '.'}...)
			}
		}
		newWarehouse = append(newWarehouse, newRow)
	}

	return newWarehouse
}

func part2() int {
	warehouse, moves := parseInput()
	warehouse = doubleWarehouse(warehouse)
	var robot [2]int
	for y, row := range warehouse {
		for x, char := range row {
			if char == '@' {
				robot = [2]int{x, y}
			}
		}
	}

	print(warehouse)
	directions := map[byte][2]int{
		'^': {0, -1},
		'>': {1, 0},
		'v': {0, 1},
		'<': {-1, 0},
	}

	for _, move := range moves {
		direction := directions[move]
		newPosition := [2]int{robot[0] + direction[0], robot[1] + direction[1]}
		if direction[1] == 0 {
			// Horizontal movement

		}

	}

	return scoreGps(warehouse)
	return 0
}

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %v in %v\n", part1(), time.Since(start))
	half := time.Now()
	fmt.Printf("Part 2: %v in %v\n", part2(), time.Since(half))
}
