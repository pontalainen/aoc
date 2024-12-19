package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	GRID = readInput("input.txt")
	y, x := getStartPosition()

	scores := getScores(y, x)
	// fmt.Println(scores)

	lowestScore := getLowestScore(scores)
	fmt.Println(lowestScore)
}

type Coord struct {
	y int
	x int
}

var DIRS = []Coord{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
var GRID = []string{}

func readInput(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(content), "\n")
}

func getStartPosition() (int, int) {
	for y, row := range GRID {
		for x, cell := range row {
			if cell == 'S' {
				return y, x
			}
		}
	}
	return -1, -1
}

func getScores(y, x int) []int {
	visited := make(map[Coord]bool)

	scores := floodFollow(y, x, 0, visited, 0)
	return scores
}

func floodFollow(y, x int, currentDir int, visited map[Coord]bool, currentScore int) []int {
	pos := GRID[y][x]
	if pos == 'E' {
		return []int{currentScore}
	}

	visited[Coord{y, x}] = true

	localVisited := make(map[Coord]bool)
	for k, v := range visited {
		localVisited[k] = v
	}

	finalScores := []int{}
	for i, dir := range DIRS {
		dirChange := math.Abs(float64(currentDir - i))

		nextY, nextX := y + dir.y, x + dir.x
		nextPos := GRID[nextY][nextX]

		if nextPos == '#' || visited[Coord{nextY, nextX}] || dirChange == 2{
			continue
		}

		scoreAdd := 1
		if dirChange != 0 {
			scoreAdd += 1000
		}
		newScore := currentScore + scoreAdd

		finalScores = append(finalScores, floodFollow(nextY, nextX, i, localVisited, newScore)...)
	}

	return finalScores
}

func getLowestScore(scores []int) int {
	minScore := math.MaxInt
	for _, score := range scores {
		if score < minScore {
			minScore = score
		}
	}
	return minScore
}
