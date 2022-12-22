package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Monkey struct {
	operation func(map[string]Monkey) int
	val       int
}

func (m Monkey) yell(monkeys map[string]Monkey) int {
	if m.operation == nil {
		return m.val
	}
	return m.operation(monkeys)
}

func createFileScanner(fileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	reader := bufio.NewScanner(file)
	return reader, file
}

func getMonkeyOperation(left, op, right string) func(map[string]Monkey) int {
	if op == "*" {
		return func(monkeys map[string]Monkey) int {
			return (monkeys[left]).yell(monkeys) * (monkeys[right]).yell(monkeys)
		}
	} else if op == "/" {
		return func(monkeys map[string]Monkey) int {
			return monkeys[left].yell(monkeys) / monkeys[right].yell(monkeys)
		}
	} else if op == "+" {
		return func(monkeys map[string]Monkey) int {
			return monkeys[left].yell(monkeys) + monkeys[right].yell(monkeys)
		}
	} else {
		return func(monkeys map[string]Monkey) int {
			return monkeys[left].yell(monkeys) - monkeys[right].yell(monkeys)
		}
	}
}

func getMonkey(inputYell string) Monkey {
	val, err := strconv.Atoi(inputYell)

	if err == nil {
		return Monkey{nil, val}
	}

	var left string
	var op string
	var right string

	fmt.Sscanf(inputYell, "%s %s %s", &left, &op, &right)

	return Monkey{getMonkeyOperation(left, op, right), 0}
}

func parseInput() map[string]Monkey {
	scanner, file := createFileScanner("input.txt")
	defer file.Close()

	result := map[string]Monkey{}

	for scanner.Scan() {
		line := scanner.Text()

		r := regexp.MustCompile(`([a-zA-Z]+): ([a-zA-Z0-9+-/* ]+)`)
		matches := r.FindAllStringSubmatch(line, -1)

		name := matches[0][1]
		inputYell := matches[0][2]

		result[name] = getMonkey(inputYell)
	}
	return result
}

func solve(monkeys map[string]Monkey) int {
	return monkeys["root"].yell(monkeys)
}

func main() {
	monkeys := parseInput()
	result := solve(monkeys)

	fmt.Println(result)
}
