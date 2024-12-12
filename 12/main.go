package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readInput()

	firstHalf(input)

	// secondHalf(input)
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
}

func firstHalf(input []string) {
	regions := getRegions(input)
	
	sum := 0
	for _, region := range regions {
		regionArea := len(region.positions)
		regionPerimeter := region.perimeter

		fmt.Println(region.letter, regionArea, regionPerimeter)

		sum += regionArea * regionPerimeter
	}

	fmt.Println(sum)
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

func floodFill(grid []string, letter string, y, x int, visited map[Position]bool, ) Region {
	dirs := []Position{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	region := Region{letter: letter, positions: []Position{{y, x}}, perimeter: 0}
	stack := []Position{{y, x}}

	for len(stack) > 0 {
		pos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		visited[pos] = true
		perimeter := 0

		for _, dir := range dirs {
			newPos := Position{pos.y + dir.y, pos.x + dir.x}
			if newPos.y < 0 || newPos.y >= len(grid) ||
				newPos.x < 0 || newPos.x >= len(grid[0]) ||
				visited[newPos] || string(grid[newPos.y][newPos.x]) != letter {
					if newPos.y < 0 || newPos.y >= len(grid) ||
						newPos.x < 0 || newPos.x >= len(grid[0]) ||
						string(grid[newPos.y][newPos.x]) != letter {
							perimeter++
					}
					continue
			}

			region.positions = append(region.positions, newPos)
			stack = append(stack, newPos)
			visited[newPos] = true
		}

		region.perimeter += perimeter
	}

	return region
}

// func getRegionPerimeter(region Region) int {
// 	perimeter := 0

// 	for _, pos := range region {

// 	}
// }