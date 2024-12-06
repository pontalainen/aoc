package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, length := readFile()

	// firstHalf(input, length)
	secondHalf(input, length)
}

func readFile() ([]string, int) {
	content, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	return lines, len(content)
}

func firstHalf(input []string, totalLength int) {
	x := 0
	y := 0

	// Get starting position
	for i, line := range input {
		arrowIndex := strings.Index(line, "^")
		if arrowIndex == -1 {
			continue
		}

		x = arrowIndex
		y = i
	}

	// Walk around
	direction := "up"
	sum := 1 // Counting start position
	for i := 0; i < totalLength; i++ {
		newX, newY := move(direction, x, y)

		if newX < 0 || newY < 0 || newX >= len(input[0]) || newY >= len(input) {
			break
		}

		newPos := string(input[newY][newX])

		if newPos == "#" {
			direction = changeDirection(direction)
			continue
		}

		x = newX
		y = newY
		if newPos == "." {
			sum += 1
		}

		runes := []rune(input[y])
		runes[x] = 'X'
		input[y] = string(runes)
	}

	fmt.Println(sum)
}

func secondHalf(input []string, totalLength int) {
	startX := 0
	startY := 0

	// Get starting position
	for i, line := range input {
		arrowIndex := strings.Index(line, "^")
		if arrowIndex == -1 {
			continue
		}

		startX = arrowIndex
		startY = i
	}

	// Get all placements of "X" (possible placements of new "X")
	xPlacements := getXPlacements(input, totalLength, startX, startY)

	sum := 0
	for _, idxs := range xPlacements {
		blockX := idxs[0]
		blockY := idxs[1]

		runes := []rune(input[blockY])

		if runes[blockX] == '^' {
			continue
		}

		prevRune := runes[blockX]
		runes[blockX] = '#'
		input[blockY] = string(runes)

		visitedPositions := make(map[[2]int]string)
		
		dir := "up"
		x := startX
		y := startY
		for i := 0; i < totalLength; i++ {
			newX, newY := move(dir, x, y)
	
			if newX < 0 || newY < 0 || newX >= len(input[0]) || newY >= len(input) {
				break
			}

			newPos := string(input[newY][newX])
			if newPos == "#" {
				dir = changeDirection(dir)
				continue
			}
			
			if visitedPositions[[2]int{newX, newY}] == dir {
				sum += 1
				break
			}

			visitedPositions[[2]int{newX, newY}] = dir
	
			x = newX
			y = newY
		}

		runes[blockX] = prevRune
		input[blockY] = string(runes)
	}

	fmt.Println(sum)
}

func getXPlacements(input []string, totalLength int, x int, y int) [][]int {
	direction := "up"
	placements := [][]int{}
	uniquePlacements := make(map[[2]int]bool)

	// Initial placement (is always unique)
	placements = append(placements, []int{x, y})
	uniquePlacements[[2]int{x, y}] = true

	for i := 0; i < totalLength; i++ {
		newX, newY := move(direction, x, y)

		if newX < 0 || newY < 0 || newX >= len(input[0]) || newY >= len(input) {
			break
		}

		newPos := string(input[newY][newX])

		if newPos == "#" {
			direction = changeDirection(direction)
			continue
		}

		x = newX
		y = newY

		coord := [2]int{x, y}
		if !uniquePlacements[coord] {
			uniquePlacements[coord] = true
			placements = append(placements, []int{x, y})
		}
	}

	return placements
}

func changeDirection(direction string) string {
	newDirection := direction

	switch direction {
	case "up":
		newDirection = "right"
	case "right":
		newDirection = "down"
	case "down":
		newDirection = "left"
	case "left":
		newDirection = "up"
	}

	return newDirection
}

func move(direction string, x int, y int) (int, int) {
	switch direction {
	case "up":
		y--
	case "down":
		y++
	case "left":
		x--
	case "right":
		x++
	}

	return x, y
}
