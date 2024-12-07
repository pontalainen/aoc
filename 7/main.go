package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput()

	firstHalf(input)
}

func readInput() []string {
	lines, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(lines), "\n")
}

func firstHalf(input []string) {
	sum := 0

	for _, line := range input {
		lineParts := strings.Split(line, ": ")
		testValue, _ := strconv.Atoi(lineParts[0])
		numbers := strings.Split(lineParts[1], " ")

		spacesCount := len(numbers) - 1
		operatorPlacements := generateCombinations("", spacesCount)

		for _, op := range operatorPlacements {
			acc, _ := strconv.Atoi(numbers[0])
			for i := 0; i < len(numbers)-1; i++ {
				num, _ := strconv.Atoi(numbers[i+1])

				if op[i] == '*' {
					acc *= num
				} else {
					acc += num
				}
			}

			if acc == testValue {
				sum += testValue
				break
			}
		}
	}

	fmt.Println(sum)
}

func generateCombinations(current string, length int) []string {
	if len(current) == length {
		return []string{current}
	}
	return append(generateCombinations(current+"*", length), generateCombinations(current+"+", length)...)
}
