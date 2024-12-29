package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var lines [][]string
var iterationCount int = 3
var xmas string = "XMAS"

func main() {
	file := scanFile()
	rawContents := bufio.NewScanner(file)
	convertFileValuesToString(rawContents)
	processData()
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
		line := strings.Split(contents.Text(), "")
		lines = append(lines, line)
	}
}

func processData() {
	count := 0
	for index := range lines {
		count += checkUp(index)
		count += checkDown(index)
		count += checkRight(index)
		count += checkLeft(index)
		count += checkUpLeft(index)
		count += checkUpRight(index)
		count += checkDownLeft(index)
		count += checkDownRight(index)
	}

	log.Println(count)
}

func checkUp(index int) int {
	if index < 3 {
		return 0
	}

	currentLine := lines[index]
	inputs := []string{}

	for charIndex, character := range currentLine {

		if character != "X" {
			continue
		}

		holder := []string{}

		for i := 0; i < iterationCount; i++ {
			holder = append(holder, lines[index-(i+1)][charIndex])
		}
		inputs = append(inputs, character+strings.Join(holder, ""))
	}

	return countXmasOccurrence(inputs)
}

func checkDown(index int) int {
	if index > (len(lines)-1)-iterationCount {
		return 0
	}

	currentLine := lines[index]
	inputs := []string{}

	for charIndex, character := range currentLine {

		if character != "X" {
			continue
		}

		holder := []string{}

		for i := 0; i < iterationCount; i++ {
			holder = append(holder, lines[index+(i+1)][charIndex])
		}
		inputs = append(inputs, character+strings.Join(holder, ""))
	}

	return countXmasOccurrence(inputs)
}

func checkLeft(index int) int {
	lineSize := len(lines[index])
	line := lines[index]
	inputs := []string{}

	for i := 0; i < lineSize; i++ {
		holder := []string{}
		if i < iterationCount {
			continue
		}

		if line[i] != "X" {
			continue
		}

		for j := i; j > i-iterationCount; j-- {
			// log.Fatal(j)
			holder = append(holder, line[j-1])
		}

		inputs = append(inputs, "X"+strings.Join(holder, ""))
	}

	return countXmasOccurrence(inputs)
}

func checkRight(index int) int {
	lineSize := len(lines[index])
	line := lines[index]
	inputs := []string{}

	for i := 0; i < lineSize-1; i++ {
		holder := []string{}
		if i > lineSize-4 {
			break
		}

		if lines[index][i] != "X" {
			continue
		}

		for j := i; j < i+iterationCount; j++ {
			holder = append(holder, line[j+1])
		}

		inputs = append(inputs, "X"+strings.Join(holder, ""))

	}

	return countXmasOccurrence(inputs)
}

func checkUpLeft(index int) int {
	if index < iterationCount {
		return 0
	}

	line := lines[index]
	lineSize := len(line)
	inputs := []string{}

	for i := 0; i < lineSize; i++ {
		if i < iterationCount {
			continue
		}

		if line[i] != "X" {
			continue
		}

		holder := []string{}

		//j is loop for lines
		indexForChar := i
		for j := index; j > index-iterationCount; j-- {
			currLine := lines[j-1]
			holder = append(holder, currLine[indexForChar-1])
			indexForChar--
		}

		inputs = append(inputs, "X"+strings.Join(holder, ""))
	}
	// log.Println(inputs)
	return countXmasOccurrence(inputs)
}

func checkUpRight(index int) int {
	if index < iterationCount {
		return 0
	}

	line := lines[index]
	lineSize := len(line)
	inputs := []string{}
	// log.Println(line)
	for i := 0; i < lineSize; i++ {
		if i > (lineSize-1)-iterationCount {
			continue
		}

		if line[i] != "X" {
			continue
		}

		holder := []string{}

		//j is loop for lines
		indexForChar := i
		for j := index; j > index-iterationCount; j-- {
			if indexForChar == lineSize-1 {
				continue
			}
			currLine := lines[j-1]
			holder = append(holder, currLine[indexForChar+1])
			indexForChar++
		}

		if len(holder) != 0 {
			inputs = append(inputs, "X"+strings.Join(holder, ""))
		}

	}
	// log.Println(inputs)
	return countXmasOccurrence(inputs)
}

func checkDownLeft(index int) int {
	if index > (len(lines)-1)-iterationCount {
		return 0
	}

	line := lines[index]
	inputs := []string{}
	lineSize := len(line)

	for i := 0; i < lineSize; i++ {
		if line[i] != "X" {
			continue
		}

		indexForChar := i
		holder := []string{}

		for j := index; j < (index + iterationCount); j++ {
			if indexForChar == 0 {
				continue
			}

			currLine := lines[j+1]
			holder = append(holder, currLine[indexForChar-1])

			indexForChar--
		}

		if len(holder) != 0 {
			inputs = append(inputs, "X"+strings.Join(holder, ""))
		}
	}
	// log.Println(inputs)
	return countXmasOccurrence(inputs)
}

func checkDownRight(index int) int {
	if index > (len(lines)-1)-iterationCount {
		return 0
	}

	line := lines[index]
	inputs := []string{}
	lineSize := len(line)

	for i := 0; i < lineSize; i++ {
		if line[i] != "X" {
			continue
		}

		indexForChar := i
		holder := []string{}

		for j := index; j < (index + iterationCount); j++ {
			if indexForChar == lineSize-1 {
				continue
			}

			currLine := lines[j+1]
			holder = append(holder, currLine[indexForChar+1])

			indexForChar++
		}

		if len(holder) != 0 {
			inputs = append(inputs, "X"+strings.Join(holder, ""))
		}
	}

	return countXmasOccurrence(inputs)
}

func countXmasOccurrence(inputs []string) int {
	count := 0
	for _, input := range inputs {
		if input == xmas {
			count++
		}
	}
	return count
}
