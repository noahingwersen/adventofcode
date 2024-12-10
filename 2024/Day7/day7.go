package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const filepath = "input.txt"

func parseInput() map[int][]int {
	file, ferr := os.Open(filepath)
	check(ferr)

	scanner := bufio.NewScanner(file)
	equations := make(map[int][]int)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")
		answer, _ := strconv.Atoi(parts[0])
		var numbers []int
		for _, str := range strings.Split(parts[1], " ") {
			number, _ := strconv.Atoi(str)
			numbers = append(numbers, number)
		}
		equations[answer] = numbers

	}
	file.Close()
	return equations
}

func endsWith(a int, b int) bool {
	return (a-b)%int(math.Pow10(len(strconv.Itoa(b)))) == 0
}

// Works backwards to determine solvability
func solvable(numbers *[]int, index int, current int, part2 bool) bool {
	if index < 0 || current <= 0 {
		return false
	}

	if index == 0 {
		return current == (*numbers)[0]
	}

	// Multiplication is valid if current value is perfectly divisble by previous
	if current%(*numbers)[index] == 0 {
		if solvable(numbers, index-1, current/(*numbers)[index], part2) {
			return true
		}
	}

	// Concatenation operator is valid if the current number ends with the previous
	if part2 && endsWith(current, (*numbers)[index]) {
		if solvable(numbers, index-1, current/int(math.Pow10(len(strconv.Itoa((*numbers)[index])))), part2) {
			return true
		}
	}

	// Addition is always possible
	return solvable(numbers, index-1, current-(*numbers)[index], part2)

}

func part1() int {
	equations := parseInput()
	total := 0
	for answer, numbers := range equations {
		if solvable(&numbers, len(numbers)-1, answer, false) {
			total += answer
		}

	}
	return total
}

func part2() int {
	equations := parseInput()
	total := 0
	for answer, numbers := range equations {
		if solvable(&numbers, len(numbers)-1, answer, true) {
			total += answer
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
