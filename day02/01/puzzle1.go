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

		//if len(lineSplit) == 2 && bufTrend >= -3 && bufTrend <= 3 {
		//	safeReport += 1
		//	continue
		//}

		safeLine := true
		for index := 2; index < len(lineSplit); index++ {
			curElem, _ := strconv.Atoi(lineSplit[index])
			prevElem, _ := strconv.Atoi(lineSplit[index-1])
			curTrend := prevElem - curElem


			if curTrend < -3 && curTrend > 3  || curTrend * bufTrend <= 0 || curTrend == 0 {
				safeLine = false
				break
			}
		}

		if (safeLine) {
			safeReport += 1
		}

	}
	fmt.Println("safe report : ", safeReport)

}
