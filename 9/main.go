package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := readInput()

	firstHalf(input)
	// secondHalf(input)
}

func readInput() string {
	line, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(line)
}

func firstHalf(input string) {
	dottedSlice := getDottedSlice(input)
	leftedSlice := getLeftedSlice(dottedSlice)
	checkSum := getCheckSum(leftedSlice)

	fmt.Println(checkSum)
}

func getDottedSlice(input string) []string {
	dottedSlice := []string{}
	isFileblock := true
	idx := 0
	for _, char := range input {
		valueType := "."
		if isFileblock {
			valueType = strconv.Itoa(idx)
			idx++
		}

		intChar, _ := strconv.Atoi(string(char))
		for j := 0; j < intChar; j++ {
			dottedSlice = append(dottedSlice, valueType)
		}

		isFileblock = !isFileblock
	}

	return dottedSlice
}

func getLeftedSlice(dottedSlice []string) []int {
	nonDottedSlice := []string{}
	for _, idChar := range dottedSlice {
		if idChar == "." {
			continue
		}
		nonDottedSlice = append(nonDottedSlice, idChar)
	}

	leftedSlice := []int{}
	secondIdx := len(nonDottedSlice) - 1
	for i, idChar := range dottedSlice {
		if i == len(nonDottedSlice) {
			break
		}
		if idChar == "." {
			intValue, _ := strconv.Atoi(nonDottedSlice[secondIdx])
			leftedSlice = append(leftedSlice, intValue)
			secondIdx--
		} else {
			intValue, _ := strconv.Atoi(idChar)
			leftedSlice = append(leftedSlice, intValue)
		}
	}
	return leftedSlice
}

func getCheckSum(leftedSlice []int) int {
	sum := 0
	for i := 0; i < len(leftedSlice); i++ {
		placeInt := leftedSlice[i]
		sum += placeInt * i
	}

	return sum
}
