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

func getMatchPoints(opponent string, you string, outcomeScores map[string]int) int {
	shapes := []string{"ROCK", "PAPER", "SCISSORS"}

	opponentIdx := getShapeIdx(opponent, shapes)
	youIdx := getShapeIdx(you, shapes)

	if (youIdx+1)%len(shapes) == opponentIdx {
		return outcomeScores["LOST"]
	} else if opponentIdx == youIdx {
		return outcomeScores["DRAW"]
	} else {
		return outcomeScores["WIN"]
	}

	return 0
}

func solve(input [][]string) int {
	shapeScores := map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}
	outcomeScores := map[string]int{"LOST": 0, "DRAW": 3, "WIN": 6}
	shapes := map[string]string{"A": "ROCK", "X": "ROCK",
		"B": "PAPER", "Y": "PAPER",
		"C": "SCISSORS", "Z": "SCISSORS"}

	result := 0
	for _, match := range input {
		opponent := match[0]
		you := match[1]

		result += shapeScores[you] + getMatchPoints(shapes[opponent], shapes[you], outcomeScores)
	}
	return result
}

func main() {
	input := parseInput()
	result := solve(input)

	fmt.Println(result)
}
