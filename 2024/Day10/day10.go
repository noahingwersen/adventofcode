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

func isValid(current *[2]int, next *[2]int, trailMap []string) bool {
	if (*next)[0] < 0 || (*next)[0] >= len(trailMap[0]) || (*next)[1] < 0 || (*next)[1] >= len(trailMap) {
		return false
	}

	// Byte value difference will be the same as actual integer difference
	heightDifference := trailMap[(*next)[1]][(*next)[0]] - trailMap[(*current)[1]][(*current)[0]]
	return heightDifference == 1
}

func scoreTrail(start [2]int, trailMap []string, part1 bool) int {
	directions := [][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	queue := [][2]int{start}
	visited := make(map[[2]int]bool)
	score := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if part1 {
			if visited[current] {
				continue
			} else {
				visited[current] = true
			}
		}

		if string(trailMap[current[1]][current[0]]) == "9" {
			score++
			continue
		}

		for _, direciton := range directions {
			next := [2]int{current[0] + direciton[0], current[1] + direciton[1]}
			if isValid(&current, &next, trailMap) {
				queue = append(queue, next)
			}
		}
	}

	return score
}

func part1() int {
	trailMap := parseInput()
	score := 0
	for y, row := range trailMap {
		for x, char := range row {
			if char == '0' {
				score += scoreTrail([2]int{x, y}, trailMap, true)
			}
		}
	}

	return score
}

func part2() int {
	trailMap := parseInput()
	score := 0
	for y, row := range trailMap {
		for x, char := range row {
			if char == '0' {
				score += scoreTrail([2]int{x, y}, trailMap, false)
			}
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
