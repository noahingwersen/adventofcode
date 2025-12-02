package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	r := regexp.MustCompile(`-?\d+`)
	matches := r.FindAll(input, -1)

	total := 0
	for _, match := range matches {
		value, _ := strconv.Atoi(string(match))
		total += value
	}

	fmt.Println(total)
}
