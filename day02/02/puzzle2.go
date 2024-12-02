package main

import (
	"bufio"
	"fmt"
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

	scanner := bufio.NewScanner(file)
	safeReport := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")

		// determine first trend
		firstElem, _ := strconv.Atoi(lineSplit[0])
		secElem, _ := strconv.Atoi(lineSplit[1])
		bufTrend := firstElem - secElem

		ignoredLevel := 0
		safeLine := true
		if bufTrend < -3 || bufTrend > 3 {
			// no need to continue if the first trend isn't right, right ?
			continue
		}

		for index := 2; index < len(lineSplit); index++ {
			prevElem, _ := strconv.Atoi(lineSplit[index-1])
			curElem, _ := strconv.Atoi(lineSplit[index])
			curTrend := prevElem - curElem

			// if curTrend * bufTrend is negative, it means -A * B or A * -B
			// so they don't have the same trend
			if (curTrend < -3 || curTrend > 3) ||
				curTrend*bufTrend <= 0 {
				// if level isn't good, remove the pref elem and try
				if ignoredLevel == 1 {
					safeLine = false
					break
				}

				// i'm too lazy to refacto this 
				ignoredLevel += 1
				prevElem, _ = strconv.Atoi(lineSplit[index-2])
				curTrend := prevElem - curElem
				if (curTrend < -3 || curTrend > 3) || curTrend*bufTrend <= 0 {
					safeLine = false
					break
				}
			}
		}

		if safeLine {
			safeReport += 1
		}
	}
	fmt.Println("safe report : ", safeReport)

}
