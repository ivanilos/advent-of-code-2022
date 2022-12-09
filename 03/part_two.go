package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func createFileScanner(fileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	reader := bufio.NewScanner(file)
	return reader, file
}

func parseInput() []string {
	scanner, file := createFileScanner("input.txt")
	defer file.Close()

	result := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	return result
}

func getCommonItemForThree(left string, mid string, right string) rune {
	aux := getCommonItems(left, mid)
	result := getCommonItems(string(aux), right)
	return rune(result[0])
}

func getCommonItems(left string, right string) []rune {
	leftMap := map[rune]bool{}

	for _, val := range left {
		leftMap[val] = true
	}

	result := []rune{}
	for _, val := range right {
		if leftMap[val] == true {
			result = append(result, val)
		}
	}
	return result
}

func getPriority(val rune) int {
	if unicode.IsLower(val) {
		return int(val - 'a' + 1)
	} else {
		return int(val - 'A' + 1 + 26)
	}
}

func solve(input []string) int {
	result := 0

	for i := 0; i < len(input); i += 3 {
		commonItem := getCommonItemForThree(input[i], input[i+1], input[i+2])
		result += getPriority(commonItem)
	}
	return result
}

func main() {
	input := parseInput()
	result := solve(input)

	fmt.Println(result)
}
