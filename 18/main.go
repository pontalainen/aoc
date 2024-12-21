package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	gridSize := 70
	bytesFallen := 1024
	coordinates := readInput("input.txt")

	currentCoords := coordinates[:bytesFallen + 1]
	grid := getMemorySpace(currentCoords, gridSize)

	smallest := dijkstras(grid, gridSize)
	fmt.Println("Answer first half:", smallest)

	for i := bytesFallen + 1; i < len(coordinates); i++ {
		newCoord := coordinates[i]
		currentCoords = append(currentCoords, newCoord)
		grid = getMemorySpace(currentCoords, gridSize)
		smallest := dijkstras(grid, gridSize)
		if smallest == 0 {
			finalFallenByte := strings.Join([]string{newCoord.x, newCoord.y}, ",")
			fmt.Println("Answer second half:", finalFallenByte)
			break
		}
	}
}

type GridCoord struct {
	x string
	y string
}

func readInput(filename string) []GridCoord {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	coordStrings := strings.Split(string(content), "\n")
	coordinates := make([]GridCoord, len(coordStrings))
	for i, coordString := range coordStrings {
		xy := strings.Split(coordString, ",")
		coordinates[i] = GridCoord{x: xy[0], y: xy[1]}
	}
	return coordinates
}

func getMemorySpace(coordinates []GridCoord, gridSize int) [][]string {
	grid := make([][]string, gridSize + 1)
	for i := range grid {
		grid[i] = strings.Split(strings.Repeat(".", gridSize + 1), "")
	}

	for i := 0; i < len(coordinates); i++ {
		coord := coordinates[i]
		xInt, _ := strconv.Atoi(coord.x)
		yInt, _ := strconv.Atoi(coord.y)
		grid[yInt][xInt] = "#"
	}

	return grid
}

type Coord struct {
	x int
	y int
}

type State struct {
	x, y int
	visited []Coord
}

func dijkstras(grid [][]string, gridSize int) int {
	// Starts at the top left corner - 0, 0
	startCoord := Coord{0, 0}
	queue := []State{{startCoord.x, startCoord.y, []Coord{}}}
	distance := make(map[Coord]int)
	smallest := 0
	// bestPath := []Coord{}
	dirs := []Coord{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		prev, found := distance[Coord{current.y, current.x}]
		if found && len(current.visited) >= prev {
			continue
		}

		if current.y == gridSize && current.x == gridSize {
			if smallest == 0 || len(current.visited) < smallest {
				smallest = len(current.visited)
				// bestPath = current.visited
			}
			continue
		}
		
		distance[Coord{current.y, current.x}] = len(current.visited)
		current.visited = append(current.visited, Coord{current.y, current.x})

		for _, dir := range dirs {
			nextY, nextX := current.y + dir.y, current.x + dir.x
			outside := nextY < 0 || nextY > gridSize || nextX < 0 || nextX > gridSize
			if outside || grid[nextY][nextX] == "#" {
				continue
			}

			queue = append(queue, State{nextX, nextY, current.visited})
		}
	}

	// fmt.Println(bestPath)

	return smallest
}
