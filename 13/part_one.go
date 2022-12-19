package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Packet struct {
	items []Packet
	val   int
}

const INF = int(1e9)

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

	input := [][]string{}

	pair := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			input = append(input, pair)
			pair = []string{}
		} else {
			pair = append(pair, line)
		}
	}
	input = append(input, pair)

	return input
}

func process(packetInput string) Packet {
	if packetInput == "[]" {
		return Packet{nil, -INF}
	} else if packetInput[0] == '[' {
		leftIdx := 1
		rightIdx := 0
		balance := 0

		result := Packet{[]Packet{}, -1}

		for rightIdx+1 < len(packetInput)-1 {
			rightIdx++

			if packetInput[rightIdx] == ',' && balance == 0 {
				result.items = append(result.items, process(packetInput[leftIdx:rightIdx]))
				leftIdx = rightIdx + 1
			} else if packetInput[rightIdx] == '[' {
				balance++
			} else if packetInput[rightIdx] == ']' {
				balance--
			}
		}
		result.items = append(result.items, process(packetInput[leftIdx:rightIdx+1]))

		return result
	} else {
		val, _ := strconv.Atoi(packetInput)
		return Packet{nil, val}
	}
}

func isOrdered(packet1, packet2 Packet, canTie bool) bool {
	if packet1.items == nil && packet2.items == nil {
		if canTie {
			return packet1.val <= packet2.val
		} else {
			return packet1.val < packet2.val
		}
	} else if packet1.items == nil && packet2.items != nil {
		return isOrdered(packet1, packet2.items[0], true)
	} else if packet1.items != nil && packet2.items == nil {
		return isOrdered(packet1.items[0], packet2, false)
	} else {
		for i := 0; i < len(packet1.items); i++ {
			if i >= len(packet2.items) {
				return false
			}

			if isOrdered(packet1.items[i], packet2.items[i], false) {
				return true
			} else if !isOrdered(packet1.items[i], packet2.items[i], canTie) {
				return false
			}
		}
		return canTie
	}
}

func solve(packetsPairs [][]string) int {
	result := 0

	for i := 0; i < len(packetsPairs); i++ {
		packet1 := process(packetsPairs[i][0])
		packet2 := process(packetsPairs[i][1])

		if isOrdered(packet1, packet2, true) {
			result += i + 1
		}
	}
	return result
}

func main() {
	packetsPairs := parseInput()
	result := solve(packetsPairs)

	fmt.Println(result)
}
