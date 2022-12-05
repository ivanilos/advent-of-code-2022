package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	qt   int
	from int
	to   int
}

func createFileScanner(fileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	reader := bufio.NewScanner(file)
	return reader, file
}

func createStacks(stacksInput []string, stacksQt int) [][]rune {
	result := make([][]rune, stacksQt)

	for i := len(stacksInput) - 2; i >= 0; i-- {
		for j := 1; j < len(stacksInput[i]); j += 4 {
			stackPos := (j - 1) / 4

			if stacksInput[i][j] != ' ' {
				result[stackPos] = append(result[stackPos], rune(stacksInput[i][j]))
			}
		}
	}
	return result
}

func parseStacks(scanner *bufio.Scanner) [][]rune {
	stacksInput := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}
		stacksInput = append(stacksInput, line)
	}

	stacksIdx := strings.Fields(stacksInput[len(stacksInput)-1])
	stacksQt, _ := strconv.Atoi(stacksIdx[len(stacksIdx)-1])

	return createStacks(stacksInput, stacksQt)
}

func parseMoves(scanner *bufio.Scanner) []Move {
	result := []Move{}

	for scanner.Scan() {
		line := scanner.Text()

		move := Move{}
		fmt.Sscanf(line, "move %d from %d to %d", &move.qt, &move.from, &move.to)
		move.from--
		move.to--

		result = append(result, move)
	}

	return result
}

func parseInput() ([][]rune, []Move) {
	scanner, file := createFileScanner("input.txt")
	defer file.Close()

	stacks := parseStacks(scanner)
	moves := parseMoves(scanner)

	return stacks, moves
}

func reverse(s []rune) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func getTopOfStacks(stacks [][]rune) string {
	result := []rune{}

	for _, stack := range stacks {
		result = append(result, stack[len(stack)-1])
	}
	return string(result)
}

func solve(stacks [][]rune, moves []Move) string {
	for _, move := range moves {
		moveStart := len(stacks[move.from]) - move.qt
		crates := stacks[move.from][moveStart:]

		stacks[move.from] = stacks[move.from][:moveStart]

		reverse(crates)
		stacks[move.to] = append(stacks[move.to], crates...)
	}
	return getTopOfStacks(stacks)
}

func main() {
	stacks, moves := parseInput()
	result := solve(stacks, moves)

	fmt.Println(result)
}
