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

func parseInput() string {
	scanner, file := createFileScanner("input.txt")
	defer file.Close()

	scanner.Scan()

	return scanner.Text()
}

func unique(s string) bool {
	chars := map[rune]bool{}

	for _, char := range s {
		chars[char] = true
	}
	return len(chars) == len(s)
}

func solve(input string) int {
	const needDiff = 14

	for i, j := 0, needDiff-1; j < len(input); i, j = i+1, j+1 {
		if unique(input[i : j+1]) {
			return j + 1
		}
	}
	return -1
}

func main() {
	input := parseInput()
	result := solve(input)

	fmt.Println(result)
}
