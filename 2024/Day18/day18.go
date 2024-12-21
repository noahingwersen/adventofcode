package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"
)

const filepath = "input.txt"

type node struct {
	position [2]int
	steps    int
}

func parseInput(maxBytes int) [][2]int {
	file, _ := os.ReadFile(filepath)

	var bytePositions [][2]int
	for b, position := range bytes.Split(file, []byte{'\n'}) {
		if b == maxBytes {
			break
		}
		var positionValue [2]int
		for i, p := range bytes.Split(position, []byte{','}) {
			value, _ := strconv.Atoi(string(p))
			positionValue[i] = value
		}
		bytePositions = append(bytePositions, positionValue)
	}

	return bytePositions
}

func isValid(position [2]int, grid [][]int) bool {
	if position[0] < 0 || position[0] >= len(grid[0]) || position[1] < 0 || position[1] >= len(grid) {
		return false
	}

	return grid[position[1]][position[0]] >= 0
}

func bfs(start [2]int, end [2]int, grid [][]int) int {
	directions := [][2]int{
		{1, 0}, {0, 1}, {-1, 0}, {0, -1},
	}

	seen := make(map[[2]int]bool)
	queue := []node{
		{position: start, steps: 0},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.position == end {
			return current.steps
		}

		if seen[current.position] {
			continue
		}

		seen[current.position] = true
		for _, direction := range directions {
			neighbor := [2]int{current.position[0] + direction[0], current.position[1] + direction[1]}
			if isValid(neighbor, grid) {
				queue = append(queue, node{position: neighbor, steps: current.steps + 1})
			}
		}
	}
	return -1
}

func part1() int {
	size := 70
	grid := make([][]int, size+1)
	for i := range grid {
		grid[i] = make([]int, size+1)
	}

	corrupted := parseInput(1024)
	for _, position := range corrupted {
		grid[position[1]][position[0]] = -1
	}

	return bfs([2]int{0, 0}, [2]int{size, size}, grid)
}

func part2() string {
	size := 70
	grid := make([][]int, size+1)
	for i := range grid {
		grid[i] = make([]int, size+1)
	}

	corrupted := parseInput(-1)
	for _, position := range corrupted {
		grid[position[1]][position[0]] = -1
	}

	blockingByte := len(corrupted)
	for bfs([2]int{0, 0}, [2]int{size, size}, grid) == -1 {
		blockingByte--
		bytePosition := corrupted[blockingByte]
		grid[bytePosition[1]][bytePosition[0]] = 0
	}

	firstBlock := corrupted[blockingByte]
	return strconv.Itoa(firstBlock[0]) + "," + strconv.Itoa(firstBlock[1])
}

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %v in %v\n", part1(), time.Since(start))
	half := time.Now()
	fmt.Printf("Part 2: %v in %v\n", part2(), time.Since(half))
}
