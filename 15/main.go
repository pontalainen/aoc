package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	warehouse, instructions := readInput("input.txt")
	thickWarehouse := make([]string, len(warehouse))
	copy(thickWarehouse, warehouse)

	// coords := getGPSCoordinates(warehouse)
	// fmt.Println(getCoordsSum(coords))

	thickWarehouse = thicken(thickWarehouse)
	
	y, x := getStartPosition(thickWarehouse)
	for _, instruction := range instructions {
		y, x = makeThickMove(y, x, thickWarehouse, instruction)
	}

	boxes := getBoxes(thickWarehouse)
	fmt.Println(getBoxesSum(boxes))
}

func readInput(filename string) ([]string, string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	normalizedContent := strings.ReplaceAll(string(content), "\r\n", "\n")	
	contentParts := strings.Split(normalizedContent, "\n\n")
	warehouse := strings.Split(contentParts[0], "\n")
	instructions := contentParts[1]
	return warehouse, instructions
}

func getFinalWarehouse(warehouse []string, instructions string) {
	y, x := getStartPosition(warehouse)
	for _, instruction := range instructions {
		y, x = makeMove(y, x, warehouse, instruction)
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

func thicken(warehouse []string) []string {
	wareSlice := [][]byte{}
	for i, row := range warehouse {
		wareSlice = append(wareSlice, []byte{})
		for _, cell := range row {
			switch cell {
			case '#':
				wareSlice[i] = append(wareSlice[i], '#')
				wareSlice[i] = append(wareSlice[i], '#')
			case 'O':
				wareSlice[i] = append(wareSlice[i], '[')
				wareSlice[i] = append(wareSlice[i], ']')
			case '.':
				wareSlice[i] = append(wareSlice[i], '.')
				wareSlice[i] = append(wareSlice[i], '.')
			case '@':
				wareSlice[i] = append(wareSlice[i], '@')
				wareSlice[i] = append(wareSlice[i], '.')
			}
		}
	}

	thickWarehouse := []string{}
	for _, row := range wareSlice {
		thickWarehouse = append(thickWarehouse, string(row))
	}

	return thickWarehouse
}

func makeThickMove(y, x int, warehouse []string, instruction rune) (int, int) {
	path := ""
	switch instruction {
	case '>':
		path = warehouse[y][x:]
		moveAvailable := checkThickHorizontal(path)
		if moveAvailable != "" {
			moveRobotHorizontal(path, moveAvailable, warehouse, y, x, instruction)
			return y, x + 1
		}
	case 'v':
		movesAvailable, err := checkThickVertical(y, x, warehouse, 1)
		if !err {
			makeVertMoves(Pos{y, x}, movesAvailable, warehouse, 1)
			return y + 1, x
		}
	case '<':
		path = reverseString(warehouse[y][:x+1])
		moveAvailable := checkThickHorizontal(path)
		if moveAvailable != "" {
			moveRobotHorizontal(path, moveAvailable, warehouse, y, x, instruction)
			return y, x - 1
		}
	case '^':
		movesAvailable, err := checkThickVertical(y, x, warehouse, -1)
		if !err {
			makeVertMoves(Pos{y, x}, movesAvailable, warehouse, -1)
			return y - 1, x
		}
	}

	return y, x
}

func checkThickHorizontal(path string) string {
	re := regexp.MustCompile(`\@\.|\@(\[|])+\.`)
	matches := re.FindString(path)
	return matches
}

type Pos struct {
	y int
	x int
}

func checkThickVertical(y, x int, warehouse []string, direction int) ([]Pos, bool) {
	err := false
	availableMoves := []Pos{}
	visited := make(map[Pos]bool)
	stack := []Pos{{y + direction, x}}

	for len(stack) > 0 {
		current := stack[0]
		stack = stack[1:]

		outside := current.y < 0 || current.y >= len(warehouse) || current.x < 0 || current.x >= len(warehouse[current.y])
		if outside || visited[current] {
			continue
		}

		visited[current] = true

		posChar := warehouse[current.y][current.x]

		switch posChar {
		case '#':
			err = true
			break
		case '@':
			stack = append(stack, Pos{current.y + direction, current.x})
		case '[':
			stack = append(stack, Pos{current.y, current.x + 1})
			stack = append(stack, Pos{current.y + direction, current.x})
			availableMoves = append(availableMoves, current)
		case ']':
			stack = append(stack, Pos{current.y, current.x - 1})
			stack = append(stack, Pos{current.y + direction, current.x})
			availableMoves = append(availableMoves, current)
		}
	}
	return availableMoves, err
}

func moveRobotHorizontal(path string, moveAvailable string, warehouse []string, y, x int, instruction rune) {
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
			pathRunes[i] = rune(moveAvailable[i-1])
		}
		path = string(pathRunes)
	}

	editWarehouse(warehouse, path, y, x, instruction)
}

func sortVertMoves(moves []Pos, direction int) {
	if direction == 1 {
		sort.Slice(moves, func(i, j int) bool {
			return moves[i].y > moves[j].y
		})
	} else if direction == -1 {
		sort.Slice(moves, func(i, j int) bool {
			return moves[i].y < moves[j].y
		})
	}
}

func makeVertMoves(robotPos Pos, moves []Pos, warehouse []string, direction int) {
	sortVertMoves(moves, direction)
	for _, move := range moves {
		moveVertical(move, warehouse, direction)
	}
	moveVertical(robotPos, warehouse, direction)
}

func moveVertical(move Pos, warehouse []string, direction int) {
	posChar := warehouse[move.y][move.x]
	currentLine := []rune(warehouse[move.y])
	nextLine := []rune(warehouse[move.y + direction])
	
	currentLine[move.x] = '.'
	nextLine[move.x] = rune(posChar)

	warehouse[move.y] = string(currentLine)
	warehouse[move.y + direction] = string(nextLine)
}

func getBoxes(warehouse []string) []Pos {
	boxes := []Pos{}
	for y, row := range warehouse {
		for x, cell := range row {
			if cell == '[' {
				boxes = append(boxes, Pos{y, x})
			}
		}
	}
	return boxes
}

func getBoxesSum(boxes []Pos) int {
	sum := 0
	for _, box := range boxes {
		sum += box.y * 100 + box.x
	}

	return sum
}
