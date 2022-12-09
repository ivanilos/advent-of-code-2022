package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

	grid := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()

		row := []int{}
		for _, char := range line {
			row = append(row, int(char-'0'))
		}
		grid = append(grid, row)
	}
	return grid
}

func isIn(posX, posY, rows, cols int) bool {
	return 0 <= posX && posX < rows && 0 <= posY && posY < cols
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func calcScenicScoreInDirection(grid [][]int, scenicScoreInDirection [][][]int,
	rows, cols, posX, posY, dx, dy, dir int) {

	const MAX_TREE_HEIGHT = 9
	initialPosX := posX
	initialPosY := posY

	scenicScoreInDirection[posX][posY][dir] = 0
	posX += dx
	posY += dy

	lastHeightSeenPosition := map[int][]int{}

	for isIn(posX, posY, rows, cols) {
		scenicScoreInDirection[posX][posY][dir] = abs(posX-initialPosX) + abs(posY-initialPosY)

		height := grid[posX][posY]
		for i := height; i <= MAX_TREE_HEIGHT; i++ {
			if pos, ok := lastHeightSeenPosition[i]; ok {
				scenicScoreInDirection[posX][posY][dir] =
					min(scenicScoreInDirection[posX][posY][dir], abs(posX-pos[0])+abs(posY-pos[1]))
			}
		}

		lastHeightSeenPosition[height] = []int{posX, posY}

		posX += dx
		posY += dy
	}
}

func calcScenicScore(grid [][]int, scenicScoreInDirection [][][]int, rows int, cols int) {
	for i := 0; i < rows; i++ {
		calcScenicScoreInDirection(grid, scenicScoreInDirection, rows, cols, i, 0, 0, 1, 0)
		calcScenicScoreInDirection(grid, scenicScoreInDirection, rows, cols, i, cols-1, 0, -1, 1)
	}

	for i := 0; i < cols; i++ {
		calcScenicScoreInDirection(grid, scenicScoreInDirection, rows, cols, 0, i, 1, 0, 2)
		calcScenicScoreInDirection(grid, scenicScoreInDirection, rows, cols, rows-1, i, -1, 0, 3)
	}
}

func getMaxScenicScore(scenicScoreInDirection [][][]int, rows, cols int) int {
	result := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			posScore := 1
			for dir := 0; dir < 4; dir++ {
				posScore *= scenicScoreInDirection[i][j][dir]
			}

			if posScore > result {
				result = posScore
			}
		}
	}
	return result
}

func solve(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	scenicScoreInDirection := make([][][]int, rows)
	for i := 0; i < rows; i++ {
		scenicScoreInDirection[i] = make([][]int, cols)

		for j := 0; j < cols; j++ {
			scenicScoreInDirection[i][j] = make([]int, 4)
		}
	}

	calcScenicScore(grid, scenicScoreInDirection, rows, cols)
	return getMaxScenicScore(scenicScoreInDirection, rows, cols)
}

func main() {
	grid := parseInput()
	result := solve(grid)

	fmt.Println(result)
}
