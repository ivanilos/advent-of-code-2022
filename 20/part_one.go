package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func createFileScanner(fileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	reader := bufio.NewScanner(file)
	return reader, file
}

func parseInput() []int {
	scanner, file := createFileScanner("input.txt")
	defer file.Close()

	input := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		val, _ := strconv.Atoi(line)
		input = append(input, val)
	}
	return input
}

func sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	} else {
		return 0
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getZeroIdx(v []int) int {
	for i := 0; i < len(v); i++ {
		if v[i] == 0 {
			return i
		}
	}
	return -1
}

func getResult(v []int) int {
	neededIdxAfterZero := []int{1000, 2000, 3000}

	zeroIdx := getZeroIdx(v)

	result := 0
	for _, delta := range neededIdxAfterZero {
		idx := (zeroIdx + delta) % len(v)
		result += v[idx]
	}
	return result
}

func solve(input []int) int {
	sz := len(input)

	orderToPos := make([]int, sz)
	posToOrder := make([]int, sz)
	for i := 0; i < sz; i++ {
		orderToPos[i] = i
		posToOrder[i] = i
	}

	for _, curPos := range orderToPos {
		val := input[curPos]
		step := sign(val)

		for i := 1; i <= abs(val)%(sz-1); i++ {
			nextPos := (curPos + step + sz) % sz
			orderNextPos := posToOrder[nextPos]
			orderCurPos := posToOrder[curPos]

			orderToPos[orderCurPos] = nextPos
			posToOrder[nextPos] = orderCurPos

			orderToPos[orderNextPos] = curPos
			posToOrder[curPos] = orderNextPos

			input[curPos], input[nextPos] = input[nextPos], input[curPos]
			curPos = nextPos
		}
	}
	return getResult(input)
}

func main() {
	input := parseInput()
	result := solve(input)

	fmt.Println(result)
}
