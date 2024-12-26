package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	re := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)`)
	total := 0
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			digits := append(match[:0], match[1:]...)
			product := 0
			for _, digit := range digits {
				convertedDigit := convertStringToInt(digit)
				// log.Println(convertedDigit)
				if product == 0 {
					product = convertedDigit
					continue
				}

				product *= convertedDigit
			}

			total += product
		}
	}

	log.Println(total)
}

func scanFile() *os.File {
	file, err := os.Open("check.txt")

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

func convertStringToInt(number string) int {
	converted, err := strconv.Atoi(number)

	if err != nil {
		log.Fatal("Invalid input, not a number in string")
	}

	return converted
}
