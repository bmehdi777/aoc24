package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	fileString := string(file)

	res := 0
	enable:= true
	for pos, char := range fileString {
		if fileString[pos] == 'd' {
			if pos+4 < len(fileString) && fileString[pos:pos+4] == "do()" {
				enable = true
				pos = pos+4
			} else if pos+7 < len(fileString) && fileString[pos:pos+7] == "don't()" {
				enable = false
				pos = pos+7
			}
		}


		if enable && unicode.IsDigit(char) {
			lookupIndex := pos
			num1 := ""
			num2 := ""
			num3 := 0
			if fileString[pos-4:pos] == "mul(" {
				for unicode.IsDigit(rune(fileString[lookupIndex])) {
					num1 += string(fileString[lookupIndex])
					lookupIndex++
				}
				if num1 == "" {
					pos = lookupIndex
					continue
				}

				fmt.Println("num1 = ", num1)
				if fileString[lookupIndex] == ',' {
					lookupIndex++
					for unicode.IsDigit(rune(fileString[lookupIndex])) {
						num2 += string(fileString[lookupIndex])
						lookupIndex++
					}

					if num2 == "" {
						// invalid
						pos = lookupIndex
						continue
					}

					fmt.Println("num2 = ", num2)

					if fileString[lookupIndex] == ')' {
						num1, err := strconv.Atoi(num1)
						if err != nil {
							panic(err)
						}

						num2, err := strconv.Atoi(num2)
						if err != nil {
							panic(err)
						}

						fmt.Printf("mul(%v,%v)\n", num1, num2)
						num3 = num1 * num2
						res += num3
						pos = lookupIndex
					}
				} else {
					pos = lookupIndex
				}
			} else {
				pos = lookupIndex
			}
		}
	}

	fmt.Println("Result : ", res)
}
