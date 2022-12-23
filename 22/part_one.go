package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

const EMPTY = '.'
const WALL = '#'
const N_DIRS = 4

var dx []int = []int{0, 1, 0, -1}
var dy []int = []int{1, 0, -1, 0}

type Pos struct {
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

func parseGrid(scanner *bufio.Scanner) [][]rune {
	grid := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}
		grid = append(grid, []rune(line))
	}
	return grid
}

func parseMoves(scanner *bufio.Scanner) []string {
	moves := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		curMove := []rune{}

		for _, char := range line {
			if unicode.IsLetter(char) {
				moves = append(moves, string(curMove))
				moves = append(moves, string(char))

				curMove = []rune{}
			} else {
				curMove = append(curMove, char)
			}
		}
		if len(curMove) != 0 {
			moves = append(moves, string(curMove))
		}
	}
	return moves
}

func parseInput() ([][]rune, []string) {
	scanner, file := createFileScanner("input.txt")
	defer file.Close()

	grid := parseGrid(scanner)
	moves := parseMoves(scanner)

	return grid, moves
}

func getStartingPos(grid [][]rune) (int, int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == EMPTY {
				return i, j
			}
		}
	}
	return -1, -1
}

func isIn(x, y, rows, cols int) bool {
	return 0 <= x && x < rows && 0 <= y && y < cols
}

func getNextPos(grid [][]rune, x, y, dir, rows, cols int) Pos {
	for step := 1; ; step++ {
		nx := (x + step*dx[dir] + rows) % rows
		ny := (y + step*dy[dir] + cols) % cols

		if isIn(nx, ny, rows, len(grid[nx])) && (grid[nx][ny] == EMPTY || grid[nx][ny] == WALL) {
			return Pos{nx, ny}
		}
	}
}

func getNextPosSlice(grid [][]rune) [][][]Pos {
	nextPos := make([][][]Pos, len(grid))
	for i := 0; i < len(grid); i++ {
		nextPos[i] = make([][]Pos, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			nextPos[i][j] = make([]Pos, N_DIRS)
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == EMPTY {
				for dir := 0; dir < N_DIRS; dir++ {
					nextPos[i][j][dir] = getNextPos(grid, i, j, dir, len(grid), len(grid[i]))
				}
			}
		}
	}
	return nextPos
}

func posValue(x, y, dir int) int {
	return 1000*(x+1) + 4*(y+1) + dir
}

func solve(grid [][]rune, moves []string) int {
	sx, sy := getStartingPos(grid)
	nextPos := getNextPosSlice(grid)

	dir := 0
	for _, move := range moves {
		if move[0] == 'L' {
			dir = (dir - 1 + N_DIRS) % N_DIRS
		} else if move[0] == 'R' {
			dir = (dir + 1 + N_DIRS) % N_DIRS
		} else {
			steps, _ := strconv.Atoi(move)

			for step := 0; step < steps; step++ {
				pos := nextPos[sx][sy][dir]
				nx, ny := pos.x, pos.y

				if grid[nx][ny] == WALL {
					break
				}

				sx, sy = nx, ny
			}
		}
	}
	return posValue(sx, sy, dir)
}

func main() {
	grid, moves := parseInput()
	result := solve(grid, moves)

	fmt.Println(result)
}
