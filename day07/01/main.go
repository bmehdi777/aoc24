package main

import (
	"bufio"
	"os"
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

	trackNode := make([]*Node, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, ": ")

		finalRes, _ := strconv.Atoi(lineSplit[0])
		numbers := strings.Split(lineSplit[1], " ")

		num, _ := strconv.Atoi(numbers[0])
		head := &Node{Value: num}

		for i := 1; i < len(numbers); i++ {
			num, _ = strconv.Atoi(numbers[i])
			if head.Value + num <= finalRes {
				head.AddNode = &Node{Value: head.Value + num}
			}
			if head.Value * num <= finalRes {
				head.MulNode = &Node{Value: head.Value * num}
			}
			// should use recursive 

		}
	}

}

type Node struct {
	Value   int
	AddNode *Node
	MulNode *Node
}
