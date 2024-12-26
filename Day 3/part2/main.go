package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var lines string
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
	result := ""

	result = removeDonts(lines)
	matches := re.FindAllStringSubmatch(result, -1)

	for _, match := range matches {
		digits := append([]string{}, match[1:]...)
		log.Println(digits)
		product := 0
		for _, digit := range digits {
			convertedDigit := convertStringToInt(digit)
			if product == 0 {
				product = convertedDigit
				continue
			}
			product *= convertedDigit
		}

		total += product
	}

	log.Println(result)
	log.Println(total)
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
		lines += contents.Text()
	}
}

func convertStringToInt(number string) int {
	converted, err := strconv.Atoi(number)

	if err != nil {
		log.Fatal("Invalid input, not a number in string")
	}

	return converted
}

func removeDonts(line string) string {

	dontLines := []string{}
	dontIndices := getIndices(line, "don't()")

	for _, dontIndex := range dontIndices {
		stringToRemove := line[dontIndex:]
		log.Println(dontIndex)
		indexForRemoval := strings.Index(stringToRemove, "do()")

		if indexForRemoval == -1 {
			dontLines = append(dontLines, stringToRemove)
			break
		}

		dontLines = append(dontLines, stringToRemove[:indexForRemoval])
	}

	for _, dontLine := range dontLines {
		temp1 := line
		temp := strings.Replace(temp1, dontLine, "", 1)
		line = temp
	}
	return line
}

func getIndices(line string, keyword string) []int {
	indices := []int{}

	counter := 0
	for {
		index := strings.Index(line[counter:], keyword)

		if index == -1 {
			break
		}
		indices = append(indices, counter+index)
		counter += index + len(keyword)
	}

	return indices
}
