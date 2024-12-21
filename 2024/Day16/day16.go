package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"time"
)

const filepath = "input.txt"

type node struct {
	position  [2]int
	direction [2]int
	score     int
	visited   [][2]int
}

func parseInput() [][]byte {
	file, _ := os.ReadFile(filepath)

	return bytes.Split(file, []byte{'\n'})
}

func isValid(position [2]int, maze [][]byte) bool {
	return maze[position[1]][position[0]] != '#'
}

// TODO: Try a priority queue
func bestPath(start [2]int, end [2]int, maze [][]byte) (int, int) {
	minScore := -1
	queue := []node{{position: start, direction: [2]int{1, 0}}}
	seen := make(map[[2][2]int]int)
	paths := make(map[int]map[[2]int]bool)

	directions := [4][2]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		current.visited = append(current.visited, current.position)

		key := [2][2]int{current.position, current.direction}
		previousScore, inSeen := seen[key]
		if inSeen && current.score > previousScore {
			continue
		}

		if minScore != -1 && current.score > minScore {
			continue
		}

		if current.position == end {
			if minScore == -1 {
				minScore = current.score
			} else {
				minScore = min(minScore, current.score)
			}

			if minScore == current.score {
				if paths[current.score] == nil {
					paths[current.score] = make(map[[2]int]bool)
				}
				for _, tile := range current.visited {
					paths[current.score][tile] = true
				}
			}
		}

		seen[key] = current.score
		for _, direction := range directions {
			newPosition := [2]int{current.position[0] + direction[0], current.position[1] + direction[1]}
			if isValid(newPosition, maze) {
				newScore := current.score + 1
				if direction != current.direction {
					newScore += 1000
				}
				queue = append(queue, node{position: newPosition, direction: direction, score: newScore, visited: slices.Clone(current.visited)})
			}
		}

	}

	bestPathTiles := 0
	for range paths[minScore] {
		bestPathTiles++
	}

	return minScore, bestPathTiles
}

func part1() int {
	maze := parseInput()
	var start [2]int
	var end [2]int
	for y, row := range maze {
		for x, char := range row {
			if char == 'S' {
				start = [2]int{x, y}
			} else if char == 'E' {
				end = [2]int{x, y}
			}
		}
	}

	score, _ := bestPath(start, end, maze)
	return score
}

func part2() int {
	maze := parseInput()
	var start [2]int
	var end [2]int
	for y, row := range maze {
		for x, char := range row {
			if char == 'S' {
				start = [2]int{x, y}
			} else if char == 'E' {
				end = [2]int{x, y}
			}
		}
	}

	_, tiles := bestPath(start, end, maze)
	return tiles
}

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %v in %v\n", part1(), time.Since(start))
	half := time.Now()
	fmt.Printf("Part 2: %v in %v\n", part2(), time.Since(half))
}
