package main

import (
	"fmt"
	"os"
	"slices"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const filepath = "input.txt"

func parseInput() string {
	file, ferr := os.ReadFile(filepath)
	check(ferr)
	return string(file)
}

func triangularDifference(block [2]int) int {
	// https://en.wikipedia.org/wiki/Triangular_number
	sum1 := ((block[0] - 1) * block[0]) / 2
	sum2 := ((block[1] - 1) * block[1]) / 2

	return sum2 - sum1
}

func checksum(files map[int][][2]int) int {
	total := 0
	for id, locations := range files {
		for _, block := range locations {

			total += triangularDifference(block) * id
		}
	}

	return total
}

func checksum2(files [][2]int) int {
	total := 0
	for id, location := range files {
		total += triangularDifference(location) * id
	}

	return total
}

func separateDisk(disk string) ([][2]int, [][2]int) {
	currentBlock := 0
	var fileBlocks [][2]int
	var freeSpaces [][2]int
	for i, itemRune := range disk {
		// https://stackoverflow.com/questions/21322173/convert-rune-to-int
		item := int(itemRune - '0')
		if i%2 == 0 {
			fileBlocks = append(fileBlocks, [2]int{currentBlock, currentBlock + item})
		} else {
			freeSpaces = append(freeSpaces, [2]int{currentBlock, currentBlock + item})
		}
		currentBlock += item
	}

	return fileBlocks, freeSpaces

}

func part1() int {
	disk := parseInput()
	fileBlocks, freeSpaces := separateDisk(disk)

	fileLocations := make(map[int][][2]int)
	currentFile := len(fileBlocks) - 1
	currentSpace := 0

	for currentSpace < len(freeSpaces) && currentFile >= 0 {
		if freeSpaces[currentSpace][0] > fileBlocks[currentFile][1] {
			break
		}

		spaceLength := freeSpaces[currentSpace][1] - freeSpaces[currentSpace][0]
		fileLength := fileBlocks[currentFile][1] - fileBlocks[currentFile][0]

		if fileLength <= spaceLength {
			spaceStart := freeSpaces[currentSpace][0]
			fileLocations[currentFile] = append(fileLocations[currentFile], [2]int{spaceStart, spaceStart + fileLength})

			if fileLength == spaceLength {
				currentSpace++
			} else {
				freeSpaces[currentSpace] = [2]int{spaceStart + fileLength, freeSpaces[currentSpace][1]}
			}
			currentFile--
		} else {
			fileLocations[currentFile] = append(fileLocations[currentFile], freeSpaces[currentSpace])

			fileBlocks[currentFile] = [2]int{fileBlocks[currentFile][0], fileBlocks[currentFile][1] - spaceLength}
			currentSpace++
		}

	}

	for fileId := range currentFile + 1 {
		fileLocations[fileId] = append(fileLocations[fileId], fileBlocks[fileId])
	}

	return checksum(fileLocations)
}

func part2() int {
	disk := parseInput()
	fileBlocks, freeSpaces := separateDisk(disk)

	for index := range len(fileBlocks) {
		fileId := len(fileBlocks) - 1 - index
		for spaceIndex, freeSpace := range freeSpaces {
			// Free space is past file
			if freeSpace[1] > fileBlocks[fileId][0] {
				break
			}

			fileLength := fileBlocks[fileId][1] - fileBlocks[fileId][0]
			spaceLength := freeSpace[1] - freeSpace[0]
			if fileLength <= spaceLength {
				if fileLength == spaceLength {
					fileBlocks[fileId] = freeSpace
					freeSpaces = slices.Delete(freeSpaces, spaceIndex, spaceIndex+1)
				} else {
					fileBlocks[fileId] = [2]int{freeSpace[0], freeSpace[0] + fileLength}
					freeSpaces[spaceIndex] = [2]int{freeSpace[0] + fileLength, freeSpace[1]}
				}
				break
			}
		}
	}

	return checksum2(fileBlocks)
}

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %v in %v\n", part1(), time.Since(start))
	half := time.Now()
	fmt.Printf("Part 2: %v in %v\n", part2(), time.Since(half))
}
