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
		for index, digit := range line {
			if index == len(line)-1 {
				continue
			}

			diffRight := digit - line[index+1]

			if index == 0 {
				if diffRight > 0 {
					isIncreasing = false
				} else {
					isIncreasing = true
				}
			}

			diffRight = absolute(diffRight)

			if isIncreasing {
				if digit > line[index+1] {
					isSafe = false
				}
			}

			if !isIncreasing {
				if digit < line[index+1] {
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
		// log.Println(isSafe)
		if isSafe {
			totalSafeCount += 1
		}
	}
	// log.Fatal(totalSafeCount)
	return totalSafeCount
}

func reprocess() {

}

func absolute(num int) int {
	if num < 0 {
		num = num * -1
	}
	return num
}
