package main

import (
	"fmt"
	"os"
)

func lookAndSay(number []byte) []byte {
	index := 0
	var newNumber []byte
	for index < len(number) {
		counter := 1
		pointer := index + 1
		for pointer < len(number) && number[pointer] == number[index] {
			counter++
			pointer++
		}
		newNumber = append(newNumber, byte(counter+48))
		newNumber = append(newNumber, number[index])
		index = pointer
	}

	return newNumber
}

func main() {
	number, _ := os.ReadFile("input.txt")
	for iteration := range 50 {
		number = lookAndSay(number)
		if iteration == 39 {
			fmt.Println(len(number))
		}
	}
	fmt.Println(len(number))
}
