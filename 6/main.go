package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, length := readFile()

	firstHalf(input, length)
	// secondHalf(input)
}

func readFile() ([]string, int) {
	content, err := os.ReadFile("input.txt")
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
