package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	firstHalf()
	secondHalf()
}

func readInput() ([]int, []int) {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		return nil, nil
	}
	lines := strings.Split(string(content), "\n")

	leftSide := []int{}
	rightSide := []int{}

	for _, line := range lines {
		split := strings.Split(line, "   ")
		left, _ := strconv.Atoi(split[0])
		right, _ := strconv.Atoi(split[1])
		leftSide = append(leftSide, left)
		rightSide = append(rightSide, right)
	}

	sort.Ints(leftSide)
	sort.Ints(rightSide)

	return leftSide, rightSide
}

func firstHalf() {
	leftSide, rightSide := readInput()

	sorted := [][]int{}
	for i := 0; i < len(leftSide); i++ {
		sorted = append(sorted, []int{leftSide[i], rightSide[i]})
	}

	distances := []int{}
	for _, pair := range sorted {
		dist := int(math.Abs(float64(pair[1] - pair[0])))
		distances = append(distances, dist)
	}

	sum := 0
	for _, distance := range distances {
		sum += distance
	}

	fmt.Println(sum)
}

func secondHalf() {
	leftSide, rightSide := readInput()
	sum := 0

	for _, left := range leftSide {
		showCount := 0
		for _, right := range rightSide {
			if right > left {
				break
			}
			if left == right {
				showCount++
			}
		}
		sum += showCount * left
	}

	fmt.Println(sum)
}
