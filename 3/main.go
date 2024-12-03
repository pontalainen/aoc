package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	firstHalf()
	secondHalf()
}

func readInput() string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		return ""
	}
	return string(content)
}

func getNumbers(input string) (int, int) {
	re := regexp.MustCompile(`\d+`)
	numbers := re.FindAllString(input, -1)

	nr1, _ := strconv.Atoi(numbers[0])
	nr2, _ := strconv.Atoi(numbers[1])

	return nr1, nr2
}

func firstHalf() {
	input := readInput()

	re := regexp.MustCompile(`mul\(\d+\,\d+\)`)
	muls := re.FindAllString(input, -1)

	sum := 0
	for _, mul := range muls {
		nr1, nr2 := getNumbers(mul)

		product := nr1 * nr2
		sum += product
	}

	fmt.Println(sum)
}

func secondHalf() {
	input := readInput()

	re := regexp.MustCompile(`mul\(\d+\,\d+\)`)
	
	muls := re.FindAllString(input, -1)
	mulParts := re.Split(input, -1)
	combo := make([][]string, len(mulParts)-1) // Last one will not contain any mul()

	for i, mulPart := range mulParts {
		if i == len(combo) {
			// Last one will not contain any mul()
			break
		}
		combo[i] = []string{mulPart, muls[i]}
	}

	sum := 0
	shouldDo := true
	for _, c := range combo {
		re := regexp.MustCompile(`do\(\)|don't\(\)`)
		instructions := re.FindAllString(c[0], -1)
		if len(instructions) > 0 {
			// Last instruction will decide if do or don't
			shouldDo = instructions[len(instructions)-1] == "do()"
		}

		if !shouldDo {
			continue
		}

		nr1, nr2 := getNumbers(c[1])

		product := nr1 * nr2
		sum += product
	}

	fmt.Println(sum)
}
