package main

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func main() {
	firstHalf()
}

func readInput() [][]int {
	content, err := os.ReadFile("input")
	if err != nil {
		return nil
	}
	lines := strings.Split(string(content), "\n")

	intLines := [][]int{}
	for _, line := range lines {
		split := strings.Split(line, " ")
		strings := []int{}
		for _, s := range split {
			i, _ := strconv.Atoi(s)
			strings = append(strings, i)
		}
		intLines = append(intLines, strings)
	}

	return intLines
}

func firstHalf() {
	intLines := readInput()
	safeLineCount := 0

	for _, line := range intLines {
		safe := true

		ascLine := make([]int, len(line))
		descLine := make([]int, len(line))
		copy(ascLine, line)
		copy(descLine, line)
		sort.Ints(ascLine)
		sort.Sort(sort.Reverse(sort.IntSlice(descLine)))

		if !reflect.DeepEqual(ascLine, line) && !reflect.DeepEqual(descLine, line) {
			safe = false
		}

		for i := 0; i < len(line); i++ {
			if i == len(line)-1 {
				break
			}

			diff := int(math.Abs(float64(line[i] - line[i+1])))
			if diff > 3 || diff == 0 {
				safe = false
			}
		}

		if safe {
			safeLineCount++
		}
	}
	fmt.Println(safeLineCount)
}