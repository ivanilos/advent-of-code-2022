package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func createFileScanner(fileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	reader := bufio.NewScanner(file)
	return reader, file
}

func parseInput() [][]int {
	scanner, file := createFileScanner("input.txt")
	defer file.Close()

	result := [][]int{}
	calories := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			result = append(result, calories)
			calories = []int{}
		} else {
			calorie, _ := strconv.Atoi(line)
			calories = append(calories, calorie)
		}
	}
	result = append(result, calories)

	return result
}

func get_sum(calories []int) int {
	result := 0
	for _, calorie := range calories {
		result += calorie
	}
	return result
}

func solve(input [][]int) int {
	result := []int{}
	for _, calories := range input {
		result = append(result, get_sum(calories))
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})

	return result[0] + result[1] + result[2]
}

func main() {
	input := parseInput()
	result := solve(input)

	fmt.Println(result)
}
