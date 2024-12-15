package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := readInput("input.txt")

	getPrice(input, 0)
	getPrice(input, 10000000000000)
}

func readInput(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return strings.Split(string(content), "\n\n")
}

type PlayPart struct {
	x int
	y int
}

type Play struct {
	a PlayPart
	b PlayPart
	prize PlayPart
}

func getPrice(input []string, add int) {
	sum := 0
	for _, rawPlay := range input {
		play := getPlay(rawPlay, add)
		bValue := getBValue(play)
		aValue := getAValue(play, bValue)

		if isWinnable(play, aValue, bValue) {
			priceSum := getPriceSum(aValue, bValue)
			sum += priceSum
		}
	}

	fmt.Println(sum)
}

func getPlay(rawPlay string, add int) Play {
	lines := strings.Split(rawPlay, "\n")
	re := regexp.MustCompile(`\d+`)
	play := Play{}

	for i, line := range lines {
		part := PlayPart{}
		coords := re.FindAllString(line, -1)
		part.x, _ = strconv.Atoi(coords[0])
		part.y, _ = strconv.Atoi(coords[1])

		switch i {
		case 0:
			play.a = part
		case 1:
			play.b = part
		case 2:
			play.prize = PlayPart{x: part.x + add, y: part.y + add}
		}
	}

	return play
}

func clonePlay(play Play) (PlayPart, PlayPart, PlayPart) {
	return play.a, play.b, play.prize
}

func getBValue(play Play) int {
	a, b, prize := clonePlay(play)

	b.x *= a.y
	b.y *= a.x
	prize.x *= a.y
	prize.y *= a.x

	prizeDiff := math.Abs(float64(prize.x - prize.y))
	bDiff := math.Abs(float64(b.x - b.y))

	return int(prizeDiff / bDiff)
}

func getAValue(play Play, bValue int) int {
	a, b, prize := clonePlay(play)

	b.x *= bValue

	return (prize.x - b.x) / a.x
}

func isWinnable(play Play, aValue, bValue int) bool {
	a, b, prize := clonePlay(play)

	xWin := int(a.x) * aValue + int(b.x) * bValue == int(prize.x)
	yWin := int(a.y) * aValue + int(b.y) * bValue == int(prize.y)

	return xWin && yWin
}

func getPriceSum(a, b int) int {
	return a * 3 + b
}