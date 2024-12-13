package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readInput()

	findPrice(input)
}

func readInput() []string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return strings.Split(string(content), "\n")
}

type Position struct {
	y, x int
}

type Region struct {
	letter string
	positions []Position
	perimeter int
	corners int
}

func findPrice(input []string) {
	regions := getRegions(input)
	
	sum := 0
	discountSum := 0
	for _, region := range regions {
		regionArea := len(region.positions)

		sum += regionArea * region.perimeter
		discountSum += regionArea * region.corners
	}

	fmt.Println("Area * Perimiter sum:", sum)
	fmt.Println("Area * Sides sum:", discountSum)
}

func getRegions(grid []string) []Region {
	regions := []Region{}
	visited := map[Position]bool{}
	
	// Right, Down, Left, Up
	for y, line := range grid {
		for x, char := range line {
			if visited[Position{y, x}] {
				continue
			}

			region := floodFill(grid, string(char), y, x, visited)
			regions = append(regions, region)
		}
	}

	return regions
}

func floodFill(grid []string, letter string, y, x int, visited map[Position]bool) Region {
	region := Region{
		letter: letter,
		positions: []Position{{y, x}},
		perimeter: 0,
		corners: 0,
	}
	dirs := []Position{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	stack := []Position{{y, x}}

	for len(stack) > 0 {
		pos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		visited[pos] = true
		neighbours := []Position{}
		
		for _, dir := range dirs {
			neighbours, stack = checkNeighbour(grid, dir, pos, letter, visited, neighbours, stack, &region)
		}

		switch len(neighbours) {
		case 0:
			region.corners += 4
		case 1:
			region.corners += 2
		case 2:
			posOne := neighbours[0]
			posTwo := neighbours[1]
	
			if (posOne.y == posTwo.y) || (posOne.x == posTwo.x) {
				// In the middle of two neighbours
				continue
			}

			newY := (posOne.y + posTwo.y - pos.y)
			newX := (posOne.x + posTwo.x - pos.x)
			diagonalPos := Position{newY, newX}

			region.corners += getCornersByDiagonal(diagonalPos, letter, grid) + 1 // since at least one corner
		case 3:
			for i, current := range neighbours {
				next := neighbours[0]
				if i < len(neighbours) - 1 {
					next = neighbours[i+1]
				}

				if (current.y == next.y) || (current.x == next.x) {
					// Opposing neighbours
					continue
				}

				diagonalPos := Position{(current.y + next.y - pos.y), (current.x + next.x - pos.x)}
				region.corners += getCornersByDiagonal(diagonalPos, letter, grid)
			}
		case 4:
			diagonalDirs := []Position{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
			for _, dir := range diagonalDirs {
				diagonalPos := Position{pos.y + dir.y, pos.x + dir.x}

				region.corners += getCornersByDiagonal(diagonalPos, letter, grid)
			}
		}
	}

	return region
}

func checkNeighbour(
	grid []string,
	dir Position,
	pos Position,
	letter string,
	visited map[Position]bool,
	neighbours []Position,
	stack []Position,
	region *Region,
) ([]Position, []Position) {
	newPos := Position{pos.y + dir.y, pos.x + dir.x}
			
	isOutsideOfGrid := newPos.y < 0 || newPos.y >= len(grid) ||
	newPos.x < 0 || newPos.x >= len(grid[0])
	
	if isOutsideOfGrid ||
		visited[newPos] ||
		string(grid[newPos.y][newPos.x]) != letter {
			if isOutsideOfGrid || string(grid[newPos.y][newPos.x]) != letter {
				region.perimeter += 1
			} else {
				neighbours = append(neighbours, newPos)
			}
			return neighbours, stack
		}
		
	neighbours = append(neighbours, newPos)
	region.positions = append(region.positions, newPos)
	stack = append(stack, newPos)
	visited[newPos] = true

	return neighbours, stack
}

func getCornersByDiagonal(diagonalPos Position, letter string, grid []string) int {
	corners := 0

	isOutsideOfGrid := diagonalPos.y < 0 || diagonalPos.y >= len(grid) ||
		diagonalPos.x < 0 || diagonalPos.x >= len(grid[0])
	
	if !isOutsideOfGrid && string(grid[diagonalPos.y][diagonalPos.x]) != letter {
		corners += 1
	}

	return corners
}