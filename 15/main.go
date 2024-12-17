package main

import (
	"os"
	"regexp"
	"strings"
)

func main() {
	warehouse, instructions := readInput("mini.txt")

	getFinalWarehouse(warehouse, instructions)
	// coords := getGPSCoordinates(warehouse)
	// sum := getCoordsSum(coords)

	// fmt.Println(warehouse)
	// fmt.Println(instructions)
	// fmt.Println(coords)
	// fmt.Println(sum)
}

func readInput(filename string) ([]string, string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	contentParts := strings.Split(string(content), "\n\n")
	warehouse := strings.Split(contentParts[0], "\n")
	instructions := contentParts[1]
	return warehouse, instructions
}

func getFinalWarehouse(warehouse []string, instructions string) {
	y, x := getStartPosition(warehouse)
	for _, instruction := range instructions {
		path := ""
		switch instruction {
		case '>':
			path = warehouse[y][x:]
		case '<':
			path = warehouse[y][:x]
		case '^':
			// TODO: check if this is correct
			for i := 0; i < y; i++ {
				path += string(warehouse[i][x])
			}
		case 'v':
			// TODO: check if this is correct
			for i := y + 1; i < len(warehouse); i++ {
				path += string(warehouse[i][x])
			}
		}

		moveAvailable := checkDirection(path)
		if !moveAvailable {
			continue
		}

		warehouse[y] = string(moveRobot(y, x, path, '>'))
	}
}

func getStartPosition(warehouse []string) (int, int) {
	for i, row := range warehouse {
		for j, cell := range row {
			if cell == '@' {
				return i, j
			}
		}
	}

	return -1, -1
}

func checkDirection(path string) bool {
	re := regexp.MustCompile(`.`)
	return re.MatchString(path)
}

func moveRobot(y, x int, path string, direction rune) []rune {
	// TODO: fix this so that moves are actually made
	pathRunes := []rune(path)
	pathRunes[x] = direction
	return pathRunes
}

// func getGPSCoordinates(warehouse []string) []int {

// }

// func getCoordsSum(coords []int) int {

// }
