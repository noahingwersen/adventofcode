package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	// ^ > v <
	directions := map[byte][2]int{
		94:  {0, 1},
		62:  {1, 0},
		118: {0, -1},
		60:  {-1, 0},
	}

	position := [2]int{0, 0}
	santa := position
	roboSanta := position

	p1 := map[[2]int]bool{position: true}
	p2 := map[[2]int]bool{position: true}
	for i, char := range input {
		direction := directions[char]
		position = [2]int{position[0] + direction[0], position[1] + direction[1]}
		p1[position] = true

		if i%2 == 0 {
			santa = [2]int{santa[0] + direction[0], santa[1] + direction[1]}
			p2[santa] = true
		} else {
			roboSanta = [2]int{roboSanta[0] + direction[0], roboSanta[1] + direction[1]}
			p2[roboSanta] = true
		}
	}

	fmt.Println(len(p1))
	fmt.Println(len(p2))
}
