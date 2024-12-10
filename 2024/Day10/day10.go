// TODO: Write a faster solution. BFS is faster for part1, but ideally there's a method that is effecient for both parts
package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"time"
)

const filepath = "test.txt"

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

func isValid(current *[2]int, next *[2]int, trailMap []string) bool {
	if (*next)[0] < 0 || (*next)[0] >= len(trailMap[0]) || (*next)[1] < 0 || (*next)[1] >= len(trailMap) {
		return false
	}

	// Byte value difference will be the same as actual integer difference
	heightDifference := trailMap[(*next)[1]][(*next)[0]] - trailMap[(*current)[1]][(*current)[0]]
	return heightDifference == 1
}

func scoreTrail(current [2]int, target [2]int, visited map[[2]int]bool, trailMap []string, part1 bool) int {
	if current == target {
		return 1
	}

	if visited[current] {
		return 0
	}

	visited[current] = true
	directions := [][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	score := 0
	for _, direction := range directions {
		next := [2]int{current[0] + direction[0], current[1] + direction[1]}
		if isValid(&current, &next, trailMap) {
			visitedCopy := make(map[[2]int]bool)
			maps.Copy(visitedCopy, visited)
			forkScore := scoreTrail(next, target, visitedCopy, trailMap, part1)
			if part1 && forkScore == 1 {
				return 1
			}

			score += forkScore
		}
	}

	return score
}

func part1() int {
	trailMap := parseInput()
	var trailStarts [][2]int
	var trailEnds [][2]int
	for y, row := range trailMap {
		for x, char := range row {
			if char == '0' {
				trailStarts = append(trailStarts, [2]int{x, y})
			} else if char == '9' {
				trailEnds = append(trailEnds, [2]int{x, y})
			}
		}
	}

	score := 0
	for _, start := range trailStarts {
		for _, end := range trailEnds {
			score += scoreTrail(start, end, make(map[[2]int]bool), trailMap, true)
		}
	}

	return score
}

func part2() int {
	trailMap := parseInput()
	var trailStarts [][2]int
	var trailEnds [][2]int
	for y, row := range trailMap {
		for x, char := range row {
			if char == '0' {
				trailStarts = append(trailStarts, [2]int{x, y})
			} else if char == '9' {
				trailEnds = append(trailEnds, [2]int{x, y})
			}
		}
	}

	score := 0
	for _, start := range trailStarts {
		for _, end := range trailEnds {
			score += scoreTrail(start, end, make(map[[2]int]bool), trailMap, false)
		}
	}

	return score
}

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %v in %v\n", part1(), time.Since(start))
	half := time.Now()
	fmt.Printf("Part 2: %v in %v\n", part2(), time.Since(half))
}
