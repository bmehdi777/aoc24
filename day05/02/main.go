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
	//file, err := os.Open("../input.txt")
	file, err := os.Open("../example.txt")
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

	for _, rules := range firstSection {
		numbers := strings.Split(rules, "|")

		if _, ok := dict[numbers[0]]; !ok {
			dict[numbers[0]] = []string{numbers[1]}
		} else {
			dict[numbers[0]] = append(dict[numbers[0]], numbers[1])
		}
	}

	resRow := make([]string, 0)
	res := 0

	for _, updatesRow := range secondSection {
		updates := strings.Split(updatesRow, ",")
		validRow := validRow(updates, dict)

		if validRow {
			resRow = append(resRow, updatesRow)
			num, _ := strconv.Atoi(updates[len(updates)/2])
			res += num
		} else {

		}
	}

	fmt.Println("res : ", res)
}

func sortRow(updates []string, dict map[string][]string) []string {
	for i, valI := range updates {

	}

	return nil
}

func validRow(updates []string, dict map[string][]string) bool {
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

func moveStr(array []string, srcIndex int, dstIndex int) []string {
	value := array[srcIndex]
}
