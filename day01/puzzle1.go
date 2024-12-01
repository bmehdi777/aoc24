package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	left := make([]int, 0)
	right := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, "   ")
		leftElem, _ := strconv.Atoi(lineSplit[0])
		rightElem, _ := strconv.Atoi(lineSplit[1])
		left = append(left, leftElem)
		right = append(right, rightElem)
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i,j int) bool {
		return right[i] < right[j]
	})


	sum := 0
	for index := range left {
		// go doesnt seems to have abs function for int... and i'm lazy to cast it
		val := left[index] - right[index]
		if val < 0 {
			sum += -val
		} else {
			sum += val
		}
	} 

	fmt.Println("Global distance : ", sum)

}
