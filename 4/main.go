package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readInput()
	// firstHalf(input)
	secondHalf(input)
}

func readInput() []string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(content), "\n")
}

func searchVertical(i int, ii int, lines []string) int {
	sum := 0

	// Straight up
	if i - 3 >= 0 {
		topString := ""
		for vIdx := 1; vIdx < 4; vIdx++ {
			byte := lines[i - vIdx][ii]
			topString += string(byte)
		}

		if topString == "MAS" {
			sum++
		}
	}

	// Right up
	if i - 3 >= 0 && ii + 3 < len(lines[i]) {
		topString := ""
		for vIdx := 1; vIdx < 4; vIdx++ {
			byte := lines[i - vIdx][ii + vIdx]
			topString += string(byte)
		}

		if topString == "MAS" {
			sum++
		}
	}

	// Left up
	if i - 3 >= 0 && ii - 3 >= 0 {
		topString := ""
		for vIdx := 1; vIdx < 4; vIdx++ {
			byte := lines[i - vIdx][ii - vIdx]
			topString += string(byte)
		}

		if topString == "MAS" {
			sum++
		}
	}
		
	// Straight down
	if i + 3 < len(lines) {
		bottomString := ""
		for vIdx := 1; vIdx < 4; vIdx++ {
			byte := lines[i + vIdx][ii]
			bottomString += string(byte)
		}

		if bottomString == "MAS" {
			sum++
		}
	}
		
	// Right down
	if i + 3 < len(lines) && ii + 3 < len(lines[i]) {
		bottomString := ""
		for vIdx := 1; vIdx < 4; vIdx++ {
			byte := lines[i + vIdx][ii + vIdx]
			bottomString += string(byte)
		}

		if bottomString == "MAS" {
			sum++
		}
	}
		
	// Left down
	if i + 3 < len(lines) && ii - 3 >= 0 {
		bottomString := ""
		for vIdx := 1; vIdx < 4; vIdx++ {
			byte := lines[i + vIdx][ii - vIdx]
			bottomString += string(byte)
		}

		if bottomString == "MAS" {
			sum++
		}
	}

	return sum
}

func firstHalf(lines []string) {
	sum := 0

	for i, line := range lines {
		for ii := 0; ii < len(line); ii++ {
			if line[ii] != 'X' {
				continue
			}

			hStart := ii - 3
			if hStart >= 0 {
				backward := string(line[hStart:ii])

				if backward == "SAM" {
					sum++
				}
			}

			hEnd := ii + 4
			if hEnd <= len(line) {
				forward := string(line[ii+1:hEnd])

				if forward == "MAS" {
					sum++
				}
			}

			sum += searchVertical(i, ii, lines)
		}
	}

	fmt.Println(sum)
}

func getCharByIdx(lines []string, i int, ii int) string {
	if i < 0 || i >= len(lines) || ii < 0 || ii >= len(lines[i]) {
		return ""
	}

	return string(lines[i][ii])
}

func secondHalf(lines []string) {
	sum := 0

	for i, line := range lines {
		for ii := 0; ii < len(line); ii++ {
			if line[ii] != 'A' {
				continue
			}

			topRight := getCharByIdx(lines, i - 1, ii + 1)
			bottomRight := getCharByIdx(lines, i + 1, ii + 1)
			bottomLeft := getCharByIdx(lines, i + 1, ii - 1)
			topLeft := getCharByIdx(lines, i - 1, ii - 1)

			combinations := []string{topRight, bottomRight, bottomLeft, topLeft}

			target1 := []string{"M", "M", "S", "S"}
			target2 := []string{"S", "M", "M", "S"}
			target3 := []string{"S", "S", "M", "M"}
			target4 := []string{"M", "S", "S", "M"}

			if equal(combinations, target1) || equal(combinations, target2) || equal(combinations, target3) || equal(combinations, target4) {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
