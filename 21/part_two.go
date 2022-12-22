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
	name      string
	left      string
	right     string
	op        string
	operation func(int, int) int
	val       int
}

func (m Monkey) yell(monkeys map[string]Monkey, specialMonkey string) (int, bool) {
	if m.operation == nil {
		return m.val, m.name == specialMonkey
	}

	leftVal, isSpecialOnLeft := monkeys[m.left].yell(monkeys, specialMonkey)
	rightVal, isSpecialOnRight := monkeys[m.right].yell(monkeys, specialMonkey)

	return m.operation(leftVal, rightVal), isSpecialOnLeft || isSpecialOnRight || m.name == specialMonkey
}

func createFileScanner(fileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	reader := bufio.NewScanner(file)
	return reader, file
}

func getMonkeyOperation(op string) func(int, int) int {
	if op == "*" {
		return func(leftVal, rightVal int) int {
			return leftVal * rightVal
		}
	} else if op == "/" {
		return func(leftVal, rightVal int) int {
			return leftVal / rightVal
		}
	} else if op == "+" {
		return func(leftVal, rightVal int) int {
			return leftVal + rightVal
		}
	} else {
		return func(leftVal, rightVal int) int {
			return leftVal - rightVal
		}
	}
}

func getMonkey(name, inputYell string) Monkey {
	val, err := strconv.Atoi(inputYell)

	if err == nil {
		return Monkey{name, "", "", "", nil, val}
	}

	var left string
	var op string
	var right string

	fmt.Sscanf(inputYell, "%s %s %s", &left, &op, &right)

	return Monkey{name, left, right, op, getMonkeyOperation(op), 0}
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

		result[name] = getMonkey(name, inputYell)
	}
	return result
}

func applyInverseOperation(a, b int, op string, specialMonkeySide string) int {
	if op == "*" {
		return a / b
	} else if op == "/" {
		return a * b
	} else if op == "+" {
		return a - b
	} else {
		if specialMonkeySide == "left" {
			return a + b
		} else {
			return b - a
		}
	}
}

// could improve by caching yell call values
// neededVal = unknown op a => unknown = neededVal inv(op) a
func findSpecialMonkeyValue(monkeys map[string]Monkey, rootMonkeyName string,
	specialMonkey string, neededVal int) int {

	curMonkey := monkeys[rootMonkeyName]

	for curMonkey.name != specialMonkey {
		leftVal, isSpecialOnLeft := monkeys[curMonkey.left].yell(monkeys, specialMonkey)
		rightVal, _ := monkeys[curMonkey.right].yell(monkeys, specialMonkey)

		if isSpecialOnLeft {
			neededVal = applyInverseOperation(neededVal, rightVal, curMonkey.op, "left")
			curMonkey = monkeys[curMonkey.left]
		} else {
			neededVal = applyInverseOperation(neededVal, leftVal, curMonkey.op, "right")
			curMonkey = monkeys[curMonkey.right]
		}
	}
	return neededVal
}

// assumes humn (specialMonkey) is only used once (assumes the monkey hierarchy is a tree)
func solve(monkeys map[string]Monkey) int {
	rootMonkey := monkeys["root"]
	specialMonkey := "humn"

	leftVal, isSpecialOnLeft := monkeys[rootMonkey.left].yell(monkeys, specialMonkey)
	rightVal, _ := monkeys[rootMonkey.right].yell(monkeys, specialMonkey)

	if isSpecialOnLeft {
		return findSpecialMonkeyValue(monkeys, monkeys[rootMonkey.name].left, specialMonkey, rightVal)
	} else {
		return findSpecialMonkeyValue(monkeys, monkeys[rootMonkey.name].right, specialMonkey, leftVal)
	}
}

func main() {
	monkeys := parseInput()
	result := solve(monkeys)

	fmt.Println(result)
}
