package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

const filepath = "input.txt"

func parseInput() ([][2]int, [][2]int, [][2]int) {
	input, _ := os.ReadFile(filepath)
	r := regexp.MustCompile(`(Button (?:A|B)|Prize): X(?:\+|=)(\d+), Y(?:\+|=)(\d+)`)
	matches := r.FindAllSubmatch(input, -1)

	var buttonA [][2]int
	var buttonB [][2]int
	var prize [][2]int
	for _, match := range matches {
		x, _ := strconv.Atoi(string(match[2]))
		y, _ := strconv.Atoi(string(match[3]))

		switch string(match[1]) {
		case "Button A":
			buttonA = append(buttonA, [2]int{x, y})
		case "Button B":
			buttonB = append(buttonB, [2]int{x, y})
		case "Prize":
			prize = append(prize, [2]int{x, y})
		}

	}

	return buttonA, buttonB, prize
}

func intersect(a [2]int, b [2]int, p [2]int) (int, int) {
	aPress := (p[0]*b[1] - p[1]*b[0]) / (a[0]*b[1] - a[1]*b[0])
	bPress := (a[0]*p[1] - a[1]*p[0]) / (a[0]*b[1] - a[1]*b[0])

	return aPress, bPress
}

func valid(a [2]int, aPress int, b [2]int, bPress int, prize [2]int) bool {
	return aPress*a[0]+bPress*b[0] == prize[0] && aPress*a[1]+bPress*b[1] == prize[1]
}

func part1() int {
	allButtonA, allButtonB, allPrizes := parseInput()
	total := 0
	for i, buttonA := range allButtonA {
		buttonB := allButtonB[i]
		prize := allPrizes[i]

		aPress, bPress := intersect(buttonA, buttonB, prize)
		if valid(buttonA, aPress, buttonB, bPress, prize) {
			total += aPress*3 + bPress*1
		}
	}

	return total
}

func part2() int {
	allButtonA, allButtonB, allPrizes := parseInput()
	total := 0
	for i, buttonA := range allButtonA {
		buttonB := allButtonB[i]
		prize := allPrizes[i]
		prize[0] += 10000000000000
		prize[1] += 10000000000000

		aPress, bPress := intersect(buttonA, buttonB, prize)
		if valid(buttonA, aPress, buttonB, bPress, prize) {
			total += aPress*3 + bPress*1
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
