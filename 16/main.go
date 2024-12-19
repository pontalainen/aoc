package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	GRID = readInput("input.txt")
	// fmt.Println(len(GRID))

	scores := floodFollow(getStartPosition())
	// fmt.Println(scores)

	lowestScore := getLowestScore(scores)
	fmt.Println(lowestScore)

	bestTilesCount := getBestTiles(scores, lowestScore)
	fmt.Println(bestTilesCount)
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

type state struct {
	y, x, dir, score int
	visited []Coord
}
type scoreData struct {
	score int
	path []Coord
}

func floodFollow(y, x int) []scoreData {
	queue := []state{{y, x, 0, 0, []Coord{{y, x}}}}
	finalScores := []scoreData{}
	distance := make(map[Coord]int)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		pos := GRID[current.y][current.x]
		if pos == 'E' {
			finalScores = append(finalScores, scoreData{score: current.score, path: current.visited})
			continue
		}

		currentCoord := Coord{current.y, current.x}
		if d, found := distance[currentCoord]; !found || d <= current.score {
			distance[currentCoord] = current.score
		} else {
			continue
		}

		for i, dir := range DIRS {
			nextY, nextX := current.y + dir.y, current.x + dir.x
			nextPos := GRID[nextY][nextX]
			if nextPos == '#' {
				continue
			}

			dirChange := math.Abs(float64(current.dir - i))
			scoreAdd := 1
			if dirChange != 0 {
				scoreAdd += 1000
			}
			newScore := current.score + scoreAdd
			newVisited := append([]Coord{}, current.visited...) // Make a copy of visited
			newVisited = append(newVisited, Coord{nextY, nextX})

			queue = append(queue, state{nextY, nextX, i, newScore, newVisited})
		}
	}

	return finalScores
}

func getLowestScore(scores []scoreData) int {
	minScore := math.MaxInt
	for _, score := range scores {
		if score.score < minScore {
			minScore = score.score
		}
	}
	return minScore
}

func getBestTiles(scores []scoreData, lowestScore int) int {
	bestTiles := map[Coord]bool{}
	count := 0
	for _, score := range scores {
		if score.score == lowestScore {
			count++
			for _, tile := range score.path {
				bestTiles[tile] = true
			}
		}
	}

	vari := count

	fmt.Println(vari)

	return len(bestTiles)
}

