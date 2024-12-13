package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findValue(wire string, wires map[string]int, instructions map[string][]string) int {
	existingValue, found := wires[wire]
	if found {
		return existingValue
	}

	instruction := instructions[wire]
	var value int
	if len(instruction) == 1 {
		equals, err := strconv.Atoi(instruction[0])
		if err != nil {
			equals = findValue(instruction[0], wires, instructions)
		}
		value = equals
	} else if len(instruction) == 2 {
		value = ^findValue(instruction[1], wires, instructions)
	} else {
		lhs, lerr := strconv.Atoi(instruction[0])
		rhs, rerr := strconv.Atoi(instruction[2])
		if lerr != nil {
			lhs = findValue(instruction[0], wires, instructions)
		}
		if rerr != nil {
			rhs = findValue(instruction[2], wires, instructions)
		}

		switch instruction[1] {
		case "AND":
			value = lhs & rhs
		case "OR":
			value = lhs | rhs
		case "RSHIFT":
			value = lhs >> rhs
		case "LSHIFT":
			value = lhs << rhs
		}
	}
	if value < 0 {
		value += 65536
	}

	wires[wire] = value
	return value
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	wires := make(map[string]int)
	instructions := make(map[string][]string)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		instruction := strings.Split(parts[0], " ")
		instructions[parts[1]] = instruction
	}

	wireA := findValue("a", wires, instructions)
	fmt.Println(wireA)
	wires = make(map[string]int)
	wires["b"] = wireA
	fmt.Println(findValue("a", wires, instructions))
}
