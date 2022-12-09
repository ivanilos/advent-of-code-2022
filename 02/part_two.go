package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func createFileScanner(fileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	reader := bufio.NewScanner(file)
	return reader, file
}

func parseInput() [][]string {
	scanner, file := createFileScanner("input.txt")
	defer file.Close()

	result := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, strings.Fields(line))
	}

	return result
}

func getShapeIdx(shape string, shapes []string) int {
	for idx := range shapes {
		if shape == shapes[idx] {
			return idx
		}
	}
	return -1
}

func getChosenShape(opponent string, result string) string {
	shapes := []string{"A", "B", "C"}

	opponentIdx := getShapeIdx(opponent, shapes)

	if result == "X" {
		return shapes[(opponentIdx-1+len(shapes))%len(shapes)]
	} else if result == "Y" {
		return shapes[opponentIdx]
	} else {
		return shapes[(opponentIdx+1)%len(shapes)]
	}

	return ""
}

func solve(input [][]string) int {
	shapeScores := map[string]int{"A": 1, "B": 2, "C": 3}
	outcomeScores := map[string]int{"X": 0, "Y": 3, "Z": 6}

	result := 0
	for _, match := range input {
		opponent := match[0]
		matchResult := match[1]

		result += shapeScores[getChosenShape(opponent, matchResult)] + outcomeScores[matchResult]
	}
	return result
}

func main() {
	input := parseInput()
	result := solve(input)

	fmt.Println(result)
}
