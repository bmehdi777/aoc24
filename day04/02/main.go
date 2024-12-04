package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//file, err := os.Open("../input.txt")
	file, err := os.Open("../example2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	line := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}

	wordSearched := "XMAS"

	// x, y
	direction := [][]int{
		[]int{-1, -1}, // <-^
		[]int{0, -1},  // ^
		[]int{1, -1},  // ->^
		[]int{-1, 0},  // <-
		[]int{1, 0},   // ->
		[]int{-1, 1},  // v <-
		[]int{0, 1},   // v
		[]int{1, 1},   // v ->
	}

	res := 0

	for posLine, strLine := range line {
		for posChar, _ := range strLine {
			for _, dir := range direction {
				word := ""
				iLetter := 0
				lookupPosY := posLine + (iLetter * dir[1])
				lookupPosX := posChar + (iLetter * dir[0])


				// not out of bound
				for (lookupPosY < len(line) && lookupPosY >= 0) &&
					(lookupPosX < len(strLine) && lookupPosX >= 0) &&
					iLetter <= 3 &&
					line[lookupPosY][lookupPosX] == wordSearched[iLetter] {
					word += string(line[lookupPosY][lookupPosX])
					iLetter++

					lookupPosY = posLine + (iLetter * dir[1])
					lookupPosX = posChar + (iLetter * dir[0])
				}
				if word == wordSearched {
					res += 1
				}
			}

		}
	}

	fmt.Println("Found : ", res)
}
