package main

import (
	"fmt"
	"os"
)

func main() {
	input := readFile()

	firstHalf(input)
	// secondHalf(input)
}

func readFile() string {
	line, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	return string(line)
}

func firstHalf(input string) {
	fmt.Println(input)
}
