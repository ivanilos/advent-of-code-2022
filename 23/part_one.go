package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const ELF = '#'
const EMPTY = '.'
const MOVE_PROPOSITIONS = 4
const INF = int(1e9)

type MoveProposition struct {
	check_dx []int
	check_dy []int
	moveX    int
	moveY    int
}

type Pos struct {
	x int
	y int
}

func isAnotherElfAround(x, y int, elvesPositions map[Pos]bool) bool {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			nx := x + dx
			ny := y + dy

			if _, found := elvesPositions[Pos{nx, ny}]; found {
				return true
			}
		}
	}
	return false
}

func (m MoveProposition) check(x, y int, elvesPositions map[Pos]bool) bool {
	if !isAnotherElfAround(x, y, elvesPositions) {
		return false
	}

	for k := 0; k < len(m.check_dx); k++ {
		nx := x + m.check_dx[k]
		ny := y + m.check_dy[k]

		if _, found := elvesPositions[Pos{nx, ny}]; found {
			return false
		}
	}
	return true
}

func createFileScanner(fileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	reader := bufio.NewScanner(file)
	return reader, file
}

func parseInput() [][]rune {
	scanner, file := createFileScanner("input.txt")
	defer file.Close()

	grid := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()

		grid = append(grid, []rune(line))
	}
	return grid
}

func genMovePropositions() []MoveProposition {
	toNorth := MoveProposition{
		check_dx: []int{-1, -1, -1},
		check_dy: []int{0, 1, -1},
		moveX:    -1,
		moveY:    0,
	}
	toSouth := MoveProposition{
		check_dx: []int{1, 1, 1},
		check_dy: []int{0, 1, -1},
		moveX:    1,
		moveY:    0,
	}
	toWest := MoveProposition{
		check_dx: []int{0, -1, 1},
		check_dy: []int{-1, -1, -1},
		moveX:    0,
		moveY:    -1,
	}
	toEast := MoveProposition{
		check_dx: []int{0, -1, 1},
		check_dy: []int{1, 1, 1},
		moveX:    0,
		moveY:    1,
	}
	return []MoveProposition{toNorth, toSouth, toWest, toEast}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getRectangleCornerThatContainsAllElves(elvesPositions map[Pos]bool) (int, int, int, int) {
	miniX := INF
	miniY := INF
	maxiX := -INF
	maxiY := -INF

	for elf, _ := range elvesPositions {
		miniX = min(miniX, elf.x)
		maxiX = max(maxiX, elf.x)

		miniY = min(miniY, elf.y)
		maxiY = max(maxiY, elf.y)
	}
	return miniX, maxiX, miniY, maxiY
}

func countEmptyTiles(elvesPositions map[Pos]bool, miniX, maxiX, miniY, maxiY int) int {
	return (maxiX-miniX+1)*(maxiY-miniY+1) - len(elvesPositions)
}

func countEmptyTilesInSmallestRectangleThatContainsAllElves(elvesPositions map[Pos]bool) int {
	miniX, maxiX, miniY, maxiY := getRectangleCornerThatContainsAllElves(elvesPositions)
	return countEmptyTiles(elvesPositions, miniX, maxiX, miniY, maxiY)
}

func getProposedMoves(round int, movePropositions []MoveProposition, elvesPositions map[Pos]bool) (map[Pos]Pos, map[Pos]int) {
	elfToProposedMoves := map[Pos]Pos{}
	proposedMoveToPosQt := map[Pos]int{}

	for elf, _ := range elvesPositions {
		for k := round; k < round+MOVE_PROPOSITIONS; k++ {
			if movePropositions[k%MOVE_PROPOSITIONS].check(elf.x, elf.y, elvesPositions) {
				dx := movePropositions[k%MOVE_PROPOSITIONS].moveX
				dy := movePropositions[k%MOVE_PROPOSITIONS].moveY

				newPos := Pos{elf.x + dx, elf.y + dy}
				elfToProposedMoves[elf] = newPos
				proposedMoveToPosQt[newPos]++
				break
			}
		}
	}
	return elfToProposedMoves, proposedMoveToPosQt
}

func simulate(round int, movePropositions []MoveProposition, elvesPositions map[Pos]bool) map[Pos]bool {
	elfToProposedMoves, proposedMoveToPosQt := getProposedMoves(round, movePropositions, elvesPositions)

	for elfPos, movePos := range elfToProposedMoves {
		if proposedMoveToPosQt[movePos] == 1 {
			delete(elvesPositions, elfPos)
			elvesPositions[movePos] = true
		}
	}
	return elvesPositions
}

func getElvesPositions(grid [][]rune) map[Pos]bool {
	positions := map[Pos]bool{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == ELF {
				positions[Pos{i, j}] = true
			}
		}
	}
	return positions
}

func solve(grid [][]rune) int {
	const ROUNDS = 10

	movePropositions := genMovePropositions()
	elvesPositions := getElvesPositions(grid)
	for round := 0; round < ROUNDS; round++ {
		elvesPositions = simulate(round, movePropositions, elvesPositions)
	}
	return countEmptyTilesInSmallestRectangleThatContainsAllElves(elvesPositions)
}

func main() {
	grid := parseInput()
	result := solve(grid)

	fmt.Println(result)
}
