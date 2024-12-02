package main

import (
	"bufio"
	"fmt"
	"os"
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

	scanner := bufio.NewScanner(file)
	safeReport := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")

		// does it work first ?
		safeLine := check(lineSplit)
		if safeLine {
			safeReport += 1
			continue
		}

		// try to remove every one and see if one works
		for i := 0; i < len(lineSplit); i++ {
			tmpLine := removeAt(lineSplit, i)
			safeLine := check(tmpLine)
			if safeLine {
				safeReport += 1
				break
			}
		}


	}
	fmt.Println("safe report : ", safeReport)
}

func check(lineSplit []string) bool {
	firstElem, _ := strconv.Atoi(lineSplit[0])
	secElem, _ := strconv.Atoi(lineSplit[1])
	bufTrend := firstElem - secElem
	if bufTrend < -3 || bufTrend > 3 {
		// no need to continue if the first trend isn't right, right ?
		return false
	}

	for index := 2; index < len(lineSplit); index++ {
		prevElem, _ := strconv.Atoi(lineSplit[index-1])
		curElem, _ := strconv.Atoi(lineSplit[index])
		curTrend := prevElem - curElem

		// if curTrend * bufTrend is negative, it means -A * B or A * -B
		// so they don't have the same trend
		if (curTrend < -3 || curTrend > 3) ||
			curTrend*bufTrend <= 0 {
			return false
		}
	}

	return true
}

func removeAt(s []string, index int) []string { 
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
