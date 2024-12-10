package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	floor := 0
	basement := -1
	for i, char := range input {
		if char == 40 {
			floor++
		} else if char == 41 {
			floor--
		}
		if basement == -1 && floor < 0 {
			basement = i + 1
		}
	}

	fmt.Println(floor)
	fmt.Println(basement)
}
