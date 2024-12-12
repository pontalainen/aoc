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

func firstHalf(input []string) {
	fmt.Println(input)
}

// func secondHalf(input []string) {

// }
