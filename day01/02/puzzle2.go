package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OccurenceDict = map[int]int

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	leftOccurence := make(OccurenceDict)
	left := make([]int, 0)
	right := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineSplit := strings.Split(scanner.Text(), "   ")
		leftElem, _ := strconv.Atoi(lineSplit[0])
		rightElem, _ := strconv.Atoi(lineSplit[1])
		left = append(left, leftElem)
		right = append(right, rightElem)
		leftOccurence[leftElem] = 0
	}

	for _, elem := range right {
		if _, ok := leftOccurence[elem]; ok {
			leftOccurence[elem] += 1
		} else {
			leftOccurence[elem] = 1
		}
	}

	similarityScore := 0
	for _, elem := range left {
		if val, ok := leftOccurence[elem]; ok {
			similarityScore += elem * val
		} 	
	}

	fmt.Println("Similarity score : ", similarityScore)
}
