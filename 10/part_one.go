package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func getSumOfSignalStrength(registerValueInCycle []int) int {
	neededCycles := []int{20, 60, 100, 140, 180, 220}

	result := 0
	for i := 0; i < len(neededCycles); i++ {
		idx := neededCycles[i]
		result += idx * registerValueInCycle[idx]
	}
	return result
}

func solve(input [][]string) int {
	registerValueInCycle := []int{0, 1}
	curCycle := 1

	for _, operation := range input {
		registerValueInCycle = append(registerValueInCycle, registerValueInCycle[curCycle])
		curCycle++

		if operation[0] == "noop" {
			// do nothing
		} else if operation[0] == "addx" {
			val, _ := strconv.Atoi(operation[1])
			registerValueInCycle = append(registerValueInCycle, registerValueInCycle[curCycle]+val)
			curCycle++
		}
	}
	return getSumOfSignalStrength(registerValueInCycle)
}

func main() {
	input := parseInput()
	result := solve(input)

	fmt.Println(result)
}
