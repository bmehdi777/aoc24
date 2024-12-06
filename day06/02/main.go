package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	//file, err := os.Open("../input.txt")
	file, err := os.Open("../example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	directionPos := []rune{
		'>',
		'v',
		'<',
		'^',
	}
	direction := make(map[rune][]int, 0)
	// x, y
	direction['>'] = []int{1, 0}  // right
	direction['v'] = []int{0, 1}  // down
	direction['<'] = []int{-1, 0} // left
	direction['^'] = []int{0, -1} // top

	guardMap := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		guardMap = append(guardMap, line)
	}

	numTravel := 0
firstLoop:
	for i, row := range guardMap {
		for j, char := range row {
			if char == '>' || char == '<' || char == 'v' || char == '^' {
				// found the guard, now follow him
				dir := slices.Index(directionPos, char)
				lookupChar := char
				lookupX := j + direction[lookupChar][0]
				lookupY := i + direction[lookupChar][1]

				for lookupX < len(row) && lookupX >= 0 &&
					lookupY < len(guardMap) && lookupY >= 0 {
					if guardMap[lookupY][lookupX] == '#' || rune(guardMap[lookupY][lookupX]) == directionPos[dir] {
						// lazy to do better here : just go back once and tada
						lookupX -= direction[lookupChar][0]
						lookupY -= direction[lookupChar][1]
						if dir+1 == len(directionPos) {
							dir = 0
						} else {
							dir += 1
						}
						lookupChar = directionPos[dir]
					} else if guardMap[lookupY][lookupX] == '.' {
						tmpRow := []rune(guardMap[lookupY])
						tmpRow[lookupX] = 'X'
						guardMap[lookupY] = string(tmpRow)
						numTravel += 1
					} else if guardMap[lookupY][lookupX] == 'X' {
						// search in a direction if we have another X

						for _, XDir := range directionPos {
							
						}
					}
					fmt.Println("Dir : ", direction[lookupChar])
					lookupX += direction[lookupChar][0]
					lookupY += direction[lookupChar][1]
				}
				break firstLoop
			}
		}
	}

	for _, row := range guardMap {
		fmt.Println("row", row)
	}

	fmt.Println("Guard traveled : ", numTravel+1)

}
