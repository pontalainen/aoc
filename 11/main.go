package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput()

	firstVersion(input, 25)

	//* Had to go to Reddit for help on this one :(
	//* First help needed this year though!
	secondVersion(input, 25)
	secondVersion(input, 75)
}

func readInput() []string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return strings.Split(string(content), " ")
}

func firstVersion(input []string, blinks int) {
	stones := []string{}
	stones = append(stones, input...)
	
	sum := 0
	for _, stone := range stones {
		sum += blinkStone(stone, 0, blinks)
	}

	fmt.Println(sum)
}

func blinkStone(stone string, i int, blinks int) int {
	sum := 0
	stones := []string{stone}

	intStone, _ := strconv.Atoi(stone)
	if intStone == 0 {
		stones[0] = "1"
	} else if len(stone)%2 == 0 {
		stones = splitStone(stone)
	} else {
		stones[0] = multiplyStone(stone)
	}
	
	if i == blinks-1 { // 25 blinks, 0-24
		return len(stones)
	}

	for _, stone := range stones {
		sum += blinkStone(stone, i+1, blinks)
	}

	return sum
}

func splitStone(stone string) []string {
	first, _ := strconv.Atoi(stone[:len(stone)/2])
	second, _ := strconv.Atoi(stone[len(stone)/2:])

	firstString := strconv.Itoa(first)
	secondString := strconv.Itoa(second)

	return []string{firstString, secondString}
}

func multiplyStone(stone string) string {
	intStone, _ := strconv.Atoi(stone)
	multStone := intStone * 2024

	return strconv.Itoa(multStone)
}

func secondVersion(input []string, blinks int) {
	stonesMap := make(map[string]int)
	for _, stone := range input {
		stonesMap[stone] = 1
	}

	for i := 0; i < blinks; i++ {
		newStonesMap := make(map[string]int)
		for stone, count := range stonesMap {
			blinkedStones := blinkStone2(stone)
			for _, blinkedStone := range blinkedStones {
				newStonesMap[blinkedStone] += count
			}
		}
		stonesMap = newStonesMap
	}

	sum := 0
	for _, count := range stonesMap {
		sum += count
	}

	fmt.Println(sum)
}

func blinkStone2(stone string) []string {
	stones := []string{}

	intStone, _ := strconv.Atoi(stone)
	if intStone == 0 {
		stones = append(stones, "1")
	} else if len(stone)%2 == 0 {
		stones = append(stones, splitStone(stone)...)
	} else {
		stones = append(stones, multiplyStone(stone))
	}

	return stones
}