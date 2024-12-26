package main

// full context of the code check  this link https://adventofcode.com/2024/day/2 --part 1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var lines []string
var isIncreasing bool = true

func main() {
	file := scanFile()
	rawContents := bufio.NewScanner(file)
	convertFileValuesToString(rawContents)
	processData()
}

func processData() {
	log.Println(checkIfSafeOrUnSafe(convertEachToInteger()))
}

func scanFile() *os.File {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error reading file")
	}

	return file
}

func convertFileValuesToString(contents *bufio.Scanner) {
	for contents.Scan() {
		lines = append(lines, contents.Text())
	}
}

func convertEachToInteger() [][]int {
	linesInInt := [][]int{}

	for index, line := range lines {
		fields := strings.Fields(line)

		linesInInt = append(linesInInt, []int{})

		for _, field := range fields {
			digit, err := strconv.Atoi(field)

			if err != nil {
				log.Fatal("Error Parsing String to Int")
			}
			linesInInt[index] = append(linesInInt[index], digit)
		}
	}

	return linesInInt
}

func checkIfSafeOrUnSafe(data [][]int) int {
	totalSafeCount := 0

	for _, line := range data {
		isSafe := true
		// isReprocessed := false

		for index, digit := range line {
			diffRight := 1

			if index != len(line)-1 {
				diffRight = digit - line[index+1]
			}

			if index == 0 {
				if diffRight > 0 {
					isIncreasing = false
				} else {
					isIncreasing = true
				}
			}

			diffRight = absolute(diffRight)

			if isIncreasing {
				condition := false
				indexToDelete := index + 1

				if index == len(line)-1 {
					indexToDelete = index - 1
					condition = digit < line[indexToDelete]
				} else {
					condition = digit > line[indexToDelete]
				}

				if condition {
					isSafe = false
				}
			}

			if !isIncreasing {
				condition := false
				indexToDelete := index + 1

				if index == len(line)-1 {
					indexToDelete = index - 1
					condition = digit > line[indexToDelete]
				} else {
					condition = digit < line[indexToDelete]
				}

				if condition {
					isSafe = false
				}
			}

			if diffRight == 0 {
				isSafe = false
			}

			if (diffRight) > 3 {
				isSafe = false
			}

			if !isSafe {
				// isReprocessed = true
				// log.Println(line)
				// indexToRemove := index + 1
				tempLine := append([]int{}, line...)
				// tempLine = append(tempLine[:indexToRemove], tempLine[indexToRemove+1:]...)
				isSafe = reprocess(tempLine)

				break
			}
		}
		// log.Println(isSafe)
		if isSafe {
			totalSafeCount += 1
		}
		// isReprocessed = false
	}
	// log.Fatal(totalSafeCount)
	return totalSafeCount
}

func reprocess(line []int) bool {
	isGood := true

	for index := range line {
		isSafe := true

		newLine := append([]int{}, line...)
		newLine = append(newLine[:index], newLine[index+1:]...)
		isIncreasing := true
		// log.Println(newLine)
		for newIndex, newDigit := range newLine {
			diffRight := 1
			if newIndex != len(newLine)-1 {
				diffRight = newDigit - newLine[newIndex+1]
			}

			// log.Println(diffRight)
			if newIndex == 0 {
				if diffRight > 0 {
					isIncreasing = false
				} else {
					isIncreasing = true
				}
			}

			diffRight = absolute(diffRight)

			if isIncreasing {
				condition := false
				indexToDelete := newIndex + 1

				if newIndex == len(newLine)-1 {
					indexToDelete = newIndex - 1
					condition = newDigit < newLine[indexToDelete]
				} else {
					condition = newDigit > newLine[indexToDelete]
				}

				if condition {
					isSafe = false
				}
			}

			if !isIncreasing {
				condition := false
				indexToDelete := newIndex + 1

				if newIndex == len(newLine)-1 {
					indexToDelete = newIndex - 1
					condition = newDigit > newLine[indexToDelete]
				} else {
					condition = newDigit < newLine[indexToDelete]
				}

				if condition {
					isSafe = false
				}
			}

			if diffRight == 0 {
				isSafe = false
			}

			if (diffRight) > 3 {
				isSafe = false
			}
		}
		isGood = isSafe
		if isSafe {
			// log.Println(newLine)
			break
		}
	}
	// log.Println(isGood)
	return isGood

}

func absolute(num int) int {
	if num < 0 {
		num = num * -1
	}
	return num
}
