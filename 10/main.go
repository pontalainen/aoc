package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := readFile()

	firstHalf(input)
	secondHalf(input)
}

func readFile() []string {
	line, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(line), "\n")
}

func firstHalf(input []string) {
	startingPointCoords := getStartingPointCoords(input)

	sum := 0
	for _, start := range startingPointCoords {
		foundEndsMap := map[[2]int]bool{}
		sum += followPath(input, start, 0, foundEndsMap, 1)
	}

	fmt.Println(sum)
}

func secondHalf(input []string) {
	startingPointCoords := getStartingPointCoords(input)

	sum := 0
	for _, start := range startingPointCoords {
		foundEndsMap := map[[2]int]bool{}
		sum += followPath(input, start, 0, foundEndsMap, 2)
	}

	fmt.Println(sum)
}

func getStartingPointCoords(input []string) [][]int {
	startingPointCoords := [][]int{}

	for i, line := range input {
		re := regexp.MustCompile(`0`)
		match := re.FindAllStringIndex(line, -1)
		if match != nil {
			for _, match := range match {
				startingPointCoords = append(startingPointCoords, []int{i, match[0]})
			}
		}
	}

	return startingPointCoords
}

func followPath(input []string, start []int, startValue int, foundEnds map[[2]int]bool, part int) int {
	key := [2]int{start[0], start[1]}
	if startValue == 9 {
		if part == 1 && !foundEnds[key] {
			foundEnds[key] = true
			return 1
		} else if part == 2 {
			return 1
		}
	}

	sum := 0

	topY, topX := start[0]-1, start[1]
	rightY, rightX := start[0], start[1]+1
	bottomY, bottomX := start[0]+1, start[1]
	leftY, leftX := start[0], start[1]-1

	if checkIfValid(input, startValue, topY, topX) {
		// fmt.Println("top", startValue+1)
		sum += followPath(input, []int{topY, topX}, startValue+1, foundEnds, part)
	}

	if checkIfValid(input, startValue, rightY, rightX) {
		// fmt.Println("right", startValue+1)
		sum += followPath(input, []int{rightY, rightX}, startValue+1, foundEnds, part)
	}

	if checkIfValid(input, startValue, bottomY, bottomX) {
		// fmt.Println("bottom", startValue+1)
		sum += followPath(input, []int{bottomY, bottomX}, startValue+1, foundEnds, part)
	}

	if checkIfValid(input, startValue, leftY, leftX) {
		// fmt.Println("left", startValue+1)
		sum += followPath(input, []int{leftY, leftX}, startValue+1, foundEnds, part)
	}

	return sum
}

func checkIfValid(input []string, startValue, y, x int) bool {
	stringStart := strconv.Itoa(startValue + 1)
	return y >= 0 && y < len(input) && x >= 0 && x < len(input[y]) && string(input[y][x]) == stringStart
}

