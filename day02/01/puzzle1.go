package main

import (
	"bufio"
	"os"
	"strings"
)

type Report = []int;

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	report := make([]Report, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")

	}
}
