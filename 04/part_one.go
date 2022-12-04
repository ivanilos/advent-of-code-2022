package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func createFileScanner(fileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	reader := bufio.NewScanner(file)
	return reader, file
}

func parseInput() [][][]int {
	scanner, file := createFileScanner("input.txt")
	defer file.Close()

	result := [][][]int{}

	for scanner.Scan() {
		line := scanner.Text()

		sectionsFirstElf := make([]int, 2)
		sectionsSecondElf := make([]int, 2)
		fmt.Sscanf(line, "%d-%d,%d-%d", &sectionsFirstElf[0], &sectionsFirstElf[1],
			&sectionsSecondElf[0], &sectionsSecondElf[1])

		result = append(result, [][]int{sectionsFirstElf, sectionsSecondElf})
	}

	return result
}

func isContained(rangeA, rangeB []int) bool {
	return rangeB[0] <= rangeA[0] && rangeA[1] <= rangeB[1]
}

func solve(input [][][]int) int {
	result := 0
	for _, sections := range input {
		sectionsFirstElf := sections[0]
		sectionsSecondElf := sections[1]

		if isContained(sectionsFirstElf, sectionsSecondElf) ||
			isContained(sectionsSecondElf, sectionsFirstElf) {
			result++
		}
	}
	return result
}

func main() {
	input := parseInput()
	result := solve(input)

	fmt.Println(result)
}
