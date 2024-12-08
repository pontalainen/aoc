package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	input := readInput()

	firstHalf(input)
	secondHalf(input)
}

func readInput() []string {
	lines, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(lines), "\n")
}

func firstHalf(input []string) {
	antennaTypes := make(map[string][][2]int)
	for i, line := range input {
		re := regexp.MustCompile(`[a-zA-Z]|\d`)
		lineAntennas := re.FindAllStringIndex(line, -1)
		
		for _, match := range lineAntennas {
			aType := string(line[match[0]])
			if _, ok := antennaTypes[aType]; !ok {
				antennaTypes[aType] = [][2]int{}
			}
			antennaTypes[aType] = append(antennaTypes[aType], [2]int{i, match[0]})
		}
	}
	
	antinodes := make(map[[2]int]bool)
	for _, idxs := range antennaTypes {
		for i, coords := range idxs {
			for j, otherCoords := range idxs {
				if i == j {
					continue
				}

				farCoords := getFarCoords(coords, otherCoords)
				farY, farX := farCoords[0], farCoords[1]

				isOutside := farY < 0 || farX < 0 || farY >= len(input) || farX >= len(input[farY])
				if isOutside {
					// outside of map
					continue
				}

				if _, set := antinodes[farCoords]; set {
					continue
				}

				antinodes[farCoords] = true
			}
		}
	}

	antinodesCount := len(antinodes)
	fmt.Println(antinodesCount)
}

func getFarCoords(coords, otherCoords [2]int) ([2]int) {
	diffY := otherCoords[0] - coords[0]
	diffX := otherCoords[1] - coords[1]

	farY := otherCoords[0] + diffY
	farX := otherCoords[1] + diffX

	return [2]int{farY, farX}
}

func secondHalf(input []string) {
	antennaTypes := make(map[string][][2]int)
	for i, line := range input {
		re := regexp.MustCompile(`[a-zA-Z]|\d`)
		lineAntennas := re.FindAllStringIndex(line, -1)
		
		for _, match := range lineAntennas {
			aType := string(line[match[0]])
			if _, ok := antennaTypes[aType]; !ok {
				antennaTypes[aType] = [][2]int{}
			}
			antennaTypes[aType] = append(antennaTypes[aType], [2]int{i, match[0]})
		}
	}
	
	antinodes := make(map[[2]int]bool)
	for _, idxs := range antennaTypes {
		for i, coords := range idxs {
			for j, otherCoords := range idxs {
				if i == j {
					continue
				}

				for k := 1; k < len(input); k++ {
					farY, farX := getLinedFarCoords(coords, otherCoords, k)
					
					isOutside := farY < 0 || farX < 0 || farY >= len(input) || farX >= len(input[farY])
					if isOutside {
						// outside of map
						continue
					}
					
					if _, set := antinodes[[2]int{farY, farX}]; set {
						continue
					}
					
					antinodes[[2]int{farY, farX}] = true
				}

				antinodes[[2]int{coords[0], coords[1]}] = true
			}
		}
	}

	antinodesCount := len(antinodes)
	fmt.Println(antinodesCount)
}

func getLinedFarCoords(coords, otherCoords [2]int, k int) (int, int) {
	diffY := otherCoords[0] - coords[0]
	diffX := otherCoords[1] - coords[1]

	farY := otherCoords[0] + diffY * k
	farX := otherCoords[1] + diffX * k

	return farY, farX
}
