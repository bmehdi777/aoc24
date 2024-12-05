package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	//file, err := os.Open("../example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	isFirstSec := true
	firstSection := make([]string, 0)
	secondSection := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isFirstSec && line == "" {
			isFirstSec = false
		}
		if isFirstSec {
			firstSection = append(firstSection, line)
		} else if line != "" {
			secondSection = append(secondSection, line)
		}
	}

	dict := make(map[string][]string)
	// value : [printAfterValue]

	for _, rules := range firstSection {
		numbers := strings.Split(rules, "|")

		if _, ok := dict[numbers[0]]; !ok {
			dict[numbers[0]] = []string{numbers[1]}
		} else {
			dict[numbers[0]] = append(dict[numbers[0]], numbers[1])
		}
	}
	res := 0

	sortedRow := 0
	for _, updatesRow := range secondSection {
		updates := strings.Split(updatesRow, ",")

		validRow := isValidRow(updates, dict)

		if !validRow {

			for i := 0; i < len(updates); i++ {
				for j := 0; j < len(updates); j++ {
					valI := updates[i]
					valJ := updates[j]
					if i > j && slices.Contains(dict[valI], valJ) {
						nextI := i - 1
						if nextI < 0 {
							nextI = 0
						}
						updates[i] = valJ
						updates[j] = valI

						i = 0
						j = 0
					} else if i < j && !slices.Contains(dict[valI], valJ) {
						nextI := i + 1
						if nextI > len(updates)-1 {
							nextI = len(updates) - 1
						}

						updates[i] = valJ
						updates[j] = valI

						i = 0
						j = 0
					}
				}
			}
			numSorted, _ := strconv.Atoi(updates[len(updates)/2])
			sortedRow += numSorted
		}

		num, _ := strconv.Atoi(updates[len(updates)/2])
		res += num
	}

	fmt.Println("res : ", res)
	fmt.Println("res2: ", sortedRow)
}

func isValidRow(updates []string, dict map[string][]string) bool {
	validRow := true
	for i, valI := range updates {
		for j, valJ := range updates {
			if i > j && slices.Contains(dict[valI], valJ) {
				validRow = false
				break
			} else if i < j && !slices.Contains(dict[valI], valJ) {
				validRow = false
				break
			}
		}

		if !validRow {
			break
		}
	}
	return validRow
}
