package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	warehouse, instructions := readInput("test.txt")
	thickWarehouse := make([]string, len(warehouse) + len(warehouse))

	getFinalWarehouse(warehouse, instructions)	
	coords := getGPSCoordinates(warehouse)
	fmt.Println(getCoordsSum(coords))

	thicken(thickWarehouse)
	
	// fmt.Println("")
	// for _, line := range warehouse {
	// 	fmt.Println(line)
	// }
	// fmt.Println("")
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
		y, x =makeMove(y, x, warehouse, instruction)
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

func makeMove(y, x int, warehouse []string, instruction rune) (int, int) {
	path := ""
	switch instruction {
	case '>':
		path = warehouse[y][x:]
		moveAvailable := checkDirection(path)
		if moveAvailable {
			moveRobot(path, warehouse, y, x, instruction)
			return y, x + 1
		}
	case 'v':
		for i := y; i < len(warehouse); i++ {
			path += string(warehouse[i][x])
		}
		moveAvailable := checkDirection(path)
		if moveAvailable {
			moveRobot(path, warehouse, y, x, instruction)
			return y + 1, x
		}
	case '<':
		path = reverseString(warehouse[y][:x+1])
		moveAvailable := checkDirection(path)
		if moveAvailable {
			moveRobot(path, warehouse, y, x, instruction)
			return y, x - 1
		}
	case '^':
		// TODO: check if this is correct
		for i := y; i >= 0; i-- {
			path += string(warehouse[i][x])
		}
		moveAvailable := checkDirection(path)
		if moveAvailable {
			moveRobot(path, warehouse, y, x, instruction)
			return y - 1, x
		}
	}

	return y, x
}

func checkDirection(path string) bool {
	re := regexp.MustCompile(`\@\.|\@O+\.`)
	return re.MatchString(path)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func moveRobot(path string, warehouse []string, y, x int, instruction rune) {
	dotPoint := 0
	for i, char := range path {
		if i == 0 {
			continue
		}
		if char == '#' {
			break
		}
		if char == '.' {
			dotPoint = i
			break
		}
	}

	pathRunes := []rune(path)
	if dotPoint != 0 {
		pathRunes[0] = '.'
		pathRunes[1] = '@'
		for i := 2; i <= dotPoint; i++ {
			pathRunes[i] = 'O'
		}
		path = string(pathRunes)
	}

	editWarehouse(warehouse, path, y, x, instruction)
}

func editWarehouse(warehouse []string, path string, y, x int, instruction rune) {
	switch instruction {
	case '>':
		runes := []rune(warehouse[y])
		for i := 0; i < len(path); i++ {
			runes[x+i] = rune(path[i])
		}
		warehouse[y] = string(runes)
	case 'v':
		for i := 0; i < len(path); i++ {
			runes := []rune(warehouse[y+i])
			runes[x] = rune(path[i])
			warehouse[y+i] = string(runes)
		}
	case '<':
		runes := []rune(warehouse[y])
		for i := 0; i < len(path); i++ {
			runes[x-i] = rune(path[i])
		}
		warehouse[y] = string(runes)
	case '^':
		for i := 0; i < len(path); i++ {
			runes := []rune(warehouse[y-i])
			runes[x] = rune(path[i])
			warehouse[y-i] = string(runes)
		}
	}
}

func getGPSCoordinates(warehouse []string) []int {
	coords := []int{}
	for i, row := range warehouse {
		for j, cell := range row {
			if cell == 'O' {
				coords = append(coords, i * 100 + j)
			}
		}
	}
	return coords
}

func getCoordsSum(coords []int) int {
	sum := 0
	for _, coord := range coords {
		sum += coord
	}
	return sum
}