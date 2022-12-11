package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items          []int
	operation      func(int) int
	test           func(int) bool
	throwToIfTrue  int
	throwToIfFalse int
}

func createFileScanner(fileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	reader := bufio.NewScanner(file)
	return reader, file
}

func getMonkeyItems(itemsInput string) []int {
	aux := strings.Split(itemsInput, ": ")[1]
	items := strings.Split(aux, ", ")

	result := []int{}
	for _, item := range items {
		val, _ := strconv.Atoi(item)
		result = append(result, val)
	}
	return result
}

func buildWithOperand(op string, operand string) func(int) int {
	val, _ := strconv.Atoi(operand)

	if op == "*" {
		return func(old int) int {
			return old * val
		}
	} else {
		return func(old int) int {
			return old + val
		}
	}
}

func buildWithoutOperand(op string) func(int) int {
	if op == "*" {
		return func(old int) int {
			return old * old
		}
	} else {
		return func(old int) int {
			return old + old
		}
	}
}

func getMonkeyOperation(operationInput string) func(int) int {
	aux := strings.Split(operationInput, ": ")[1]

	var op string
	var operand string

	fmt.Sscanf(aux, "new = old %s %s", &op, &operand)

	if operand == "old" {
		return buildWithoutOperand(op)
	} else {
		return buildWithOperand(op, operand)
	}
}

func getMonkeyTest(testInput string) func(int) bool {
	aux := strings.Split(testInput, ": ")[1]

	var mod int
	fmt.Sscanf(aux, "divisible by %d", &mod)

	return func(val int) bool {
		return val%mod == 0
	}
}

func getMonkeyThrowTo(inputThrowTo string) int {
	aux := strings.Split(inputThrowTo, ": ")[1]

	var throwTo int
	fmt.Sscanf(aux, "throw to monkey %d", &throwTo)

	return throwTo
}

func getMonkey(monkeyInput []string) Monkey {
	items := getMonkeyItems(monkeyInput[1])
	operation := getMonkeyOperation(monkeyInput[2])
	test := getMonkeyTest(monkeyInput[3])
	throwIfTrue := getMonkeyThrowTo(monkeyInput[4])
	throwIfFalse := getMonkeyThrowTo(monkeyInput[5])

	return Monkey{items, operation, test, throwIfTrue, throwIfFalse}
}

func scanAllMonkeys(scanner *bufio.Scanner) []Monkey {
	result := []Monkey{}

	monkeyInput := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		monkeyInput = append(monkeyInput, line)

		if len(line) == 0 {
			result = append(result, getMonkey(monkeyInput))
			monkeyInput = []string{}
		}
	}
	result = append(result, getMonkey(monkeyInput))

	return result
}

func parseInput() []Monkey {
	scanner, file := createFileScanner("input.txt")
	defer file.Close()

	result := scanAllMonkeys(scanner)
	return result
}

func multiplyTwoMaximum(v []int) int {
	sort.Slice(v, func(i, j int) bool {
		return v[i] > v[j]
	})

	return v[0] * v[1]
}

func solve(monkeys []Monkey) int {
	const rounds = 20
	const worryLevelDecrease = 3

	timesInspected := make([]int, len(monkeys))

	for round := 0; round < rounds; round++ {
		for idx, monkey := range monkeys {
			for _, item := range monkey.items {
				timesInspected[idx]++

				itemWorryLevel := monkey.operation(item)
				itemWorryLevel /= worryLevelDecrease

				if monkey.test(itemWorryLevel) {
					monkeys[monkey.throwToIfTrue].items = append(
						monkeys[monkey.throwToIfTrue].items, itemWorryLevel)
				} else {
					monkeys[monkey.throwToIfFalse].items = append(
						monkeys[monkey.throwToIfFalse].items, itemWorryLevel)
				}
			}
			monkeys[idx].items = []int{}
		}
	}
	return multiplyTwoMaximum(timesInspected)
}

func main() {
	monkeys := parseInput()
	result := solve(monkeys)

	fmt.Println(result)
}
