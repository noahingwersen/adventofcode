package main

import (
	"bufio"
	"fmt"
	"os"
)

func niceString(input string) bool {
	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}

	bad := map[string]bool{
		"ab": true,
		"cd": true,
		"pq": true,
		"xy": true,
	}

	doubleLetter := false
	vowelCount := 0
	for i := range len(input) {
		char := input[i]
		if vowels[string(char)] {
			vowelCount++
		}

		if i < len(input)-1 {
			if bad[string([]byte{char, input[i+1]})] {
				return false
			}
			if char == input[i+1] {
				doubleLetter = true
			}
		}
	}

	return doubleLetter && vowelCount >= 3
}

func niceString2(input string) bool {
	counter := make(map[[2]byte]int)
	repeat := false
	for i := range len(input) - 1 {
		letterPair := [2]byte{input[i], input[i+1]}
		if i < len(input)-2 && input[i] == input[i+2] {
			repeat = true
		}

		if repeat && counter[letterPair] >= 1 {
			return true
		}
		counter[letterPair]++
	}
	return false
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	p1 := 0
	p2 := 0
	for scanner.Scan() {
		if niceString(scanner.Text()) {
			p1++
		}
		if niceString2(scanner.Text()) {
			p2++
		}
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
