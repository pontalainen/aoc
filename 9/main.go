package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := readInput()

	// firstHalf(input)
	secondHalf(input)
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
	fragmentedSlice := getFragmentedSlice(dottedSlice)
	checkSum := getCheckSum(fragmentedSlice)

	fmt.Println(checkSum)
}

func getDottedSlice(input string) []string {
	dottedSlice := []string{}
	isFileblock := true
	idx := 0
	for _, char := range input {
		if char == '0' {
			isFileblock = !isFileblock
			continue
		}

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

func getFragmentedSlice(dottedSlice []string) []int {
	nonDottedSlice := []string{}
	for _, idChar := range dottedSlice {
		if idChar == "." {
			continue
		}
		nonDottedSlice = append(nonDottedSlice, idChar)
	}

	fragmented := []int{}
	secondIdx := len(nonDottedSlice) - 1
	for i, idChar := range dottedSlice {
		if i == len(nonDottedSlice) {
			break
		}
		if idChar == "." {
			intValue, _ := strconv.Atoi(nonDottedSlice[secondIdx])
			fragmented = append(fragmented, intValue)
			secondIdx--
		} else {
			intValue, _ := strconv.Atoi(idChar)
			fragmented = append(fragmented, intValue)
		}
	}
	return fragmented
}

func getCheckSum(finalSlice []int) int {
	sum := 0
	for i := 0; i < len(finalSlice); i++ {
		placeInt := finalSlice[i]
		sum += placeInt * i
	}

	return sum
}

func secondHalf(input string) {
	dottedSlice := getDottedPartedSlice(input)
	compactedSlice := getCompactedSlice(dottedSlice)
	checkSum := getCheckSum(compactedSlice)

	fmt.Println(checkSum)
}

func getDottedPartedSlice(input string) [][]string {
	dottedSlice := [][]string{}
	isFileblock := true
	idx := 0
	for _, char := range input {
		if char == '0' {
			isFileblock = !isFileblock
			continue
		}
		valueType := "."
		if isFileblock {
			valueType = strconv.Itoa(idx)
			idx++
		}

		intChar, _ := strconv.Atoi(string(char))
		part := []string{}
		for j := 0; j < intChar; j++ {
			part = append(part, valueType)
		}

		dottedSlice = append(dottedSlice, part)
		isFileblock = !isFileblock
	}

	return dottedSlice
}

func getCompactedSlice(dottedSlice [][]string) []int {
	compactedSlice := make([][]string, len(dottedSlice))
	copy(compactedSlice, dottedSlice)
	
	for i := len(compactedSlice) - 1; i > 0; i-- {
		slice := compactedSlice[i]
		if slice[0] == "." {
			continue
		}

		sliceLen := len(slice)
		for j, dotSlice := range compactedSlice {
			if i == j {
				break
			}

			if sliceLen > len(dotSlice) || dotSlice[0] != "." {
				continue
			}

			if sliceLen == len(dotSlice) {
				compactedSlice[i] = dotSlice
				compactedSlice[j] = slice
				break
			}

			compactedSlice[j] = slice
			compactedSlice[i] = dotSlice[:sliceLen]
			remainingDotSlice := dotSlice[sliceLen:]
			compactedSlice = append(compactedSlice[:j+1], append([][]string{remainingDotSlice}, compactedSlice[j+1:]...)...)

			i++
			break
		}
	}

	compact := []int{}
	for _, slice := range compactedSlice {
		for _, value := range slice {
			intValue, _ := strconv.Atoi(value)
			compact = append(compact, intValue)
		}
	}

	return compact
}