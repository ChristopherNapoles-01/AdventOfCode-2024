package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var lines []string

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
		lines = append(lines, contents.Text())
	}
}

func processData() {
	groupedNumbers := groupNumbers()
	total := 0

	for i := 0; i < len(groupedNumbers[0]); i++ {
		count := 0

		for j := 0; j < len(groupedNumbers[1]); j++ {
			if groupedNumbers[0][i] == groupedNumbers[1][j] {
				count++
			}
		}

		total += groupedNumbers[0][i] * count
	}

	log.Println(total)
}

func groupNumbers() [][]int {
	parsedData := [][]int{}
	sideA := []int{}
	sideB := []int{}

	for _, line := range lines {
		partA, errA := strconv.Atoi(strings.Fields(line)[0])
		partB, errB := strconv.Atoi(strings.Fields(line)[1])

		if errA == nil {
			sideA = append(sideA, partA)
		}

		if errB == nil {
			sideB = append(sideB, partB)
		}
	}

	slices.Sort(sideA)
	slices.Sort(sideB)

	parsedData = append(parsedData, sideA)
	parsedData = append(parsedData, sideB)

	return parsedData
}
