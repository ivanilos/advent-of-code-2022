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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func drawImage(registerValueInCycle []int) {
	const cycles = 240
	const crtWide = 40
	const crtHigh = 6

	curCycle := 1
	for row := 0; row < crtHigh; row++ {
		for col := 0; col < crtWide; col++ {
			if abs(col-registerValueInCycle[curCycle]) <= 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
			curCycle++
		}
		fmt.Println()
	}
}

func solve(input [][]string) {
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
	drawImage(registerValueInCycle)
}

func main() {
	input := parseInput()
	solve(input)
}
