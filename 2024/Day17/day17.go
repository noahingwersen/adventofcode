package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"
)

const filepath = "input.txt"

func parseInput() (map[byte]int, []int) {
	file, _ := os.ReadFile(filepath)
	parts := bytes.Split(file, []byte{'\n', '\n'})

	registers := make(map[byte]int)
	for _, register := range bytes.Split(parts[0], []byte{'\n'}) {
		value, _ := strconv.Atoi(string(register[12:]))
		registers[register[9]] = value
	}

	commands := bytes.Split(parts[1][9:], []byte{','})
	program := make([]int, len(commands))
	for i, value := range commands {
		// Subtract 48 from byte value to get integer value
		program[i] = int(value[0] - 48)
	}

	return registers, program
}

func combo(operand int, registers map[byte]int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return registers['A']
	case 5:
		return registers['B']
	case 6:
		return registers['C']
	default:
		return -1
	}
}

func simulateProgram(program []int, registers map[byte]int) [][]byte {
	var out [][]byte
	i := 0
	for i < len(program)-1 {
		jump := 2
		switch program[i] {
		case 0:
			// adv
			registers['A'] = registers['A'] / (1 << combo(program[i+1], registers))
		case 1:
			// bxl
			registers['B'] = registers['B'] ^ program[i+1]
		case 2:
			// bst
			registers['B'] = combo(program[i+1], registers) % 8
		case 3:
			// jnz
			if registers['A'] != 0 {
				i = program[i+1]
				jump = 0
			}
		case 4:
			// bxc
			registers['B'] = registers['B'] ^ registers['C']
		case 5:
			// out
			// Add 48 to get byte representation of integer
			value := combo(program[i+1], registers) % 8
			out = append(out, []byte{byte(value + 48)})
		case 6:
			// bdv
			registers['B'] = registers['A'] / (1 << combo(program[i+1], registers))
		case 7:
			// cdv
			registers['C'] = registers['A'] / (1 << combo(program[i+1], registers))
		}
		i += jump
	}
	return out
}

func part1() string {
	registers, program := parseInput()
	output := simulateProgram(program, registers)

	return string(bytes.Join(output, []byte{','}))
}

func copyProgram(program []int, registers map[byte]int, pointer int) int {
	if pointer == len(program) {
		return registers['A']
	}

}

func part2() int {
	_, program := parseInput()
	registers := map[byte]int{
		'A': 0,
		'B': 0,
		'C': 0,
	}

	return copyProgram(program, registers, 1)
}

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %v in %v\n", part1(), time.Since(start))
	half := time.Now()
	fmt.Printf("Part 2: %v in %v\n", part2(), time.Since(half))
}
