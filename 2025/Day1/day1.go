package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"
)

const filepath = "input.txt"

func parseInput() [][]byte {
	file, _ := os.ReadFile(filepath)
	return bytes.Split(file, []byte("\n"))
}

func part1() int {
	input := parseInput()
	dial := 50
	var zeroCounter int
	for _, instruction := range input {
		sign := instruction[0]
		value, _ := strconv.Atoi(string(instruction[1:]))
		switch sign {
		case 'R':
			dial += value
		case 'L':
			dial -= value
		default:
			panic("Invalid instruction")
		}

		if dial%100 == 0 {
			zeroCounter++
		}
	}
	return zeroCounter
}

func part2() int {
	input := parseInput()
	dial := 50
	var zeroCounter int
	wasAtZero := false
	for _, instruction := range input {
		sign := instruction[0]
		value, _ := strconv.Atoi(string(instruction[1:]))
		rotations := value / 100
		remainder := value % 100
		if rotations < 0 {
			rotations *= -1
		}
		zeroCounter += rotations
		switch sign {
		case 'R':
			dial += remainder
		case 'L':
			dial -= remainder
		default:
			panic("Invalid instruction")
		}

		if dial <= 0 || dial >= 100 {
			// We counted being at 0 the last iteration, don't double count
			if !wasAtZero {
				zeroCounter++
			}

			dial = dial % 100
			if dial < 0 {
				dial += 100
			}
		}
		wasAtZero = dial == 0
	}
	return zeroCounter
}

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %v in %v\n", part1(), time.Since(start))
	half := time.Now()
	fmt.Printf("Part 2: %v in %v\n", part2(), time.Since(half))
}
