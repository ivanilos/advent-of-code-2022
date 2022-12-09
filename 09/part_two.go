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
	dir   string
	steps int
}

type Pt struct {
	x int
	y int
}

func createFileScanner(fileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	reader := bufio.NewScanner(file)
	return reader, file
}

func parseInput() []Move {
	scanner, file := createFileScanner("input.txt")
	defer file.Close()

	moves := []Move{}

	for scanner.Scan() {
		line := scanner.Text()
		moveFields := strings.Fields(line)

		dir := moveFields[0]
		steps, _ := strconv.Atoi(moveFields[1])

		moves = append(moves, Move{dir, steps})
	}
	return moves
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x > 0 {
		return 1
	} else if x == 0 {
		return 0
	} else {
		return -1
	}
}

func getNewTailPosition(posHead Pt, posTail Pt) Pt {
	deltaX := posHead.x - posTail.x
	deltaY := posHead.y - posTail.y

	if abs(deltaX) > 1 || abs(deltaY) > 1 {
		newPosX := posTail.x + sign(deltaX)
		newPosY := posTail.y + sign(deltaY)
		return Pt{newPosX, newPosY}
	}
	return posTail
}

func solve(moves []Move) int {
	dirTodeltaX := map[string]int{"U": -1, "R": 0, "D": 1, "L": 0}
	dirTodeltaY := map[string]int{"U": 0, "R": 1, "D": 0, "L": -1}

	const ropeSize = 10
	rope := make([]Pt, ropeSize)

	visitedByTail := map[Pt]bool{}
	visitedByTail[Pt{0, 0}] = true

	for _, move := range moves {
		dx := dirTodeltaX[move.dir]
		dy := dirTodeltaY[move.dir]

		for step := 0; step < move.steps; step++ {
			rope[0] = Pt{rope[0].x + dx, rope[0].y + dy}

			for i := 1; i < ropeSize; i++ {
				rope[i] = getNewTailPosition(rope[i-1], rope[i])
			}
			visitedByTail[rope[ropeSize-1]] = true
		}
	}
	return len(visitedByTail)
}

func main() {
	moves := parseInput()
	result := solve(moves)

	fmt.Println(result)
}
