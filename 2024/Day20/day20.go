package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

const filepath = "input.txt"

func parseInput() [][]byte {
	file, _ := os.ReadFile(filepath)
	return bytes.Split(file, []byte{'\n'})
}

func isValid(position [2]int, racetrack [][]byte) bool {
	if position[0] < 0 || position[0] >= len(racetrack[0]) || position[1] < 0 || position[1] >= len(racetrack) {
		return false
	}

	return racetrack[position[1]][position[0]] != '#'
}

func bfs(start [2]int, end [2]int, racetrack [][]byte) map[[2]int]int {
	queue := [][2]int{start}
	predecessors := make(map[[2]int][2]int)
	directions := [][2]int{
		{1, 0}, {0, 1}, {-1, 0}, {0, -1},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == end {
			break
		}

		for _, direction := range directions {
			neighbor := [2]int{current[0] + direction[0], current[1] + direction[1]}
			_, visited := predecessors[neighbor]
			if !visited && isValid(neighbor, racetrack) && neighbor != start {
				queue = append(queue, neighbor)
				predecessors[neighbor] = current
			}
		}
	}

	timeToEnd := make(map[[2]int]int)
	position, seen, i := end, true, 0
	for seen {
		timeToEnd[position] = i
		position, seen = predecessors[position]
		i++
	}

	return timeToEnd
}

func difference(position1 [2]int, position2 [2]int) int {
	diff := 0
	for i := range 2 {
		thisDiff := position2[i] - position1[i]
		if thisDiff < 0 {
			thisDiff = -thisDiff
		}
		diff += thisDiff
	}
	return diff
}

func findTimeSaves(positions map[[2]int]int, maxCheat int) []int {
	var timesSaved []int
	for position, distance := range positions {
		for newPosition := range positions {
			newTime := positions[newPosition]
			diff := difference(position, newPosition)
			if diff <= maxCheat && newTime < distance {
				timeSaved := distance - newTime - diff
				if timeSaved > 0 {
					timesSaved = append(timesSaved, timeSaved)
				}
			}
		}
	}
	return timesSaved
}

func part1() int {
	racetrack := parseInput()
	var start [2]int
	var end [2]int
	for y, row := range racetrack {
		for x, tile := range row {
			if tile == 'S' {
				start = [2]int{x, y}
			} else if tile == 'E' {
				end = [2]int{x, y}
			}
		}
	}

	times := bfs(start, end, racetrack)
	total := 0
	for _, time := range findTimeSaves(times, 2) {
		if time >= 100 {
			total++
		}
	}
	return total
}

func part2() int {
	racetrack := parseInput()
	var start [2]int
	var end [2]int
	for y, row := range racetrack {
		for x, tile := range row {
			if tile == 'S' {
				start = [2]int{x, y}
			} else if tile == 'E' {
				end = [2]int{x, y}
			}
		}
	}

	times := bfs(start, end, racetrack)
	total := 0
	for _, time := range findTimeSaves(times, 20) {
		if time >= 100 {
			total++
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
