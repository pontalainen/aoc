package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := readInput("input.txt")

	gridSize := Position{x: 101, y: 103}
	positions := getRobotPositions(input, 100, gridSize)
	safetyFactor := getSafetyFactor(positions, gridSize)
	fmt.Println(safetyFactor)

	treeGrids := [][]string{}
	blinks := []int{}
	
	gridBase := []string{}
	for i := 0; i < gridSize.y; i++ {
		gridBase = append(gridBase, strings.Repeat(".", gridSize.x))
	}

	for i := 0; i < 10000; i++ {
		grid := make([]string, len(gridBase))
		copy(grid, gridBase)

		positions := getRobotPositions(input, i, gridSize)

		for _, position := range positions {
			grid[position.y] = grid[position.y][:position.x] + "#" + grid[position.y][position.x+1:]
		}

		for _, line := range grid {
			re := regexp.MustCompile(`###########`)
			if re.MatchString(line) {
				treeGrids = append(treeGrids, grid)
				blinks = append(blinks, i)
				break
			}
		}
	}

	for i, treeGrid := range treeGrids {
		for _, line := range treeGrid {
			fmt.Println(line)
		}
		fmt.Println(blinks[i])
	}

	fmt.Println(blinks)
}

func readInput(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return strings.Split(string(content), "\n")
}

type Position struct {
	x int
	y int
}

func getRobotPositions(lines []string, seconds int, gridSize Position) []Position {
	endingPositions := []Position{}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		start := strings.Split(strings.Split(parts[0], "=")[1], ",")
		velocity := strings.Split(strings.Split(parts[1], "=")[1], ",")

		startX, _ := strconv.Atoi(start[0])
		startY, _ := strconv.Atoi(start[1])
		velocityX, _ := strconv.Atoi(velocity[0])
		velocityY, _ := strconv.Atoi(velocity[1])

		endingPositions = append(endingPositions, Position{
			x: mod(startX+velocityX*seconds, gridSize.x),
			y: mod(startY+velocityY*seconds, gridSize.y),
		})
	}

	return endingPositions
}

func mod(a, b int) int {
    return (a % b + b) % b
}

func getSafetyFactor(positions []Position, gridSize Position) int {
	first := 0
	second := 0
	third := 0
	fourth := 0

	for _, position := range positions {
		if position.x == gridSize.x/2 || position.y == gridSize.y/2 {
			continue
		}

		isLeft := position.x < gridSize.x/2
		isTop := position.y < gridSize.y/2

		switch {
		case !isLeft && isTop:
			first++
		case !isLeft && !isTop:
			second++
		case isLeft && !isTop:
			third++
		case isLeft && isTop:
			fourth++
		}
	}

	safetyFactor := first * second * third * fourth

	return safetyFactor
}
