package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	orderingRules, updatePageNumbers := readFile()

	// firstHalf(orderingRules, updatePageNumbers)
	secondHalf(orderingRules, updatePageNumbers)
}

func readFile() ([]string, []string) {
	content, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	contentParts := strings.Split(string(content), "\n\n")
	orderingRules := strings.Split(contentParts[0], "\n")
	updatePageNumbers := strings.Split(contentParts[1], "\n")

	return orderingRules, updatePageNumbers
}

func firstHalf(orderingRules []string, updatePageNumbers []string) {
	sum := 0

	for _, update := range updatePageNumbers {
		faulty := false
		pageNumbers := strings.Split(update, ",")
		for _, pageNumber := range pageNumbers {
			if faulty {
				break
			}

			reString := fmt.Sprintf(`%s\|\d+|\d+\|%s`, pageNumber, pageNumber)
			re := regexp.MustCompile(reString)

			matchingRules := re.FindAllString(strings.Join(orderingRules, " "), -1)
			if len(matchingRules) == 0 {
				continue
			}

			for _, rule := range matchingRules {
				pages := strings.Split(rule, "|")
				firstPage := pages[0]
				secondPage := pages[1]

				indexFirstPage := strings.Index(update, firstPage)
				indexSecondPage := strings.Index(update, secondPage)
				if indexSecondPage == -1 {
					indexSecondPage = len(update)
				}

				if indexFirstPage > indexSecondPage {
					faulty = true
					break
				}
			}
		}
		
		if faulty {
			continue
		}

		middleNumber := len(pageNumbers) / 2
		middleNumberInt, _ := strconv.Atoi(pageNumbers[middleNumber])

		sum += middleNumberInt
	}

	fmt.Println(sum)
}

func secondHalf(orderingRules []string, updatePageNumbers []string) {
	ruleValues := make(map[string]int)
	for _, rule := range orderingRules {
		pages := strings.Split(rule, "|")
		for _, page := range pages {
			ruleValues[page] = 0
		}
	}

	for _, rule := range orderingRules {
		pages := strings.Split(rule, "|")
		ruleValues[pages[0]]++
		ruleValues[pages[1]]--
	}

	// Create a slice to hold the keys from the map
	var keys []string
	for key := range ruleValues {
		keys = append(keys, key)
	}

	// Sort the slice based on the values in the map
	sort.Slice(keys, func(i, j int) bool {
		return ruleValues[keys[i]] > ruleValues[keys[j]]
	})

	// Print sorted keys
	fmt.Println(keys)

}
