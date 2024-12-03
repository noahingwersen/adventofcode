package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const filepath = "input.txt"

func parseInput() string {
	data, ferr := os.ReadFile(filepath)
	check(ferr)

	return string(data)
}

func computeValue(instruction string) int {
	openIndex := strings.Index(instruction, "(")
	commaIndex := strings.Index(instruction, ",")
	closeIndex := strings.Index(instruction, ")")

	number1, err := strconv.Atoi(instruction[openIndex+1 : commaIndex])
	check(err)
	number2, err := strconv.Atoi(instruction[commaIndex+1 : closeIndex])
	check(err)

	return number1 * number2
}

func parseMemory(block string, part1 bool) int {
	pattern := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	matches := pattern.FindAllString(block, -1)

	total := 0
	do := true

	for _, match := range matches {
		if match == "do()" {
			do = true
		} else if match == "don't()" {
			do = false
		} else {
			if part1 || do {
				total += computeValue(match)
			}
		}
	}

	return total
}

func part1() int {
	memory := parseInput()

	return parseMemory(memory, true)
}

func part2() int {
	memory := parseInput()
	return parseMemory(memory, false)
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
