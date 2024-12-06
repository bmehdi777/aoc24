package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	//file, err := os.Open("../example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	directionPos := []rune {
		'>',
		'v',
		'<',
		'^',
	}
	direction := make(map[rune][]int, 0)
	// x, y
	direction['>'] = []int {1, 0}; // right
	direction['v'] = []int {0, -1}; // down
	direction['<'] = []int {-1, 0}; // left
	direction['^'] = []int {0, 1}; // top

	guardMap := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		guardMap = append(guardMap, line)
	}

	for i, row := range guardMap {
		for j, char := range row {
			if char == '>' || char == '<' || char == 'v' || char == '^' {
				// found the guard, now follow him
				lookupChar := char
				lookupX := j + direction[lookupChar][0]
				lookupY := i + direction[lookupChar][1]

				for lookupX < len(row) && lookupX >= 0 &&
				lookupY < len(guardMap) && lookupY >= 0 {
					if guardMap[lookupY][lookupX] == '#' {
					}

				}

			}
		}
	}

}
