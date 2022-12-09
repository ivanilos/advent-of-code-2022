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

func markVisibleInDirection(grid [][]int, visible [][]bool, rows, cols, posX, posY, dx, dy int) {

	curMax := -1
	for isIn(posX, posY, rows, cols) {
		if grid[posX][posY] > curMax {
			visible[posX][posY] = true
			curMax = grid[posX][posY]
		}
		posX += dx
		posY += dy
	}
}

func markVisible(grid [][]int, visible [][]bool, rows int, cols int) {
	for i := 0; i < rows; i++ {
		markVisibleInDirection(grid, visible, rows, cols, i, 0, 0, 1)
		markVisibleInDirection(grid, visible, rows, cols, i, cols-1, 0, -1)
	}

	for i := 0; i < cols; i++ {
		markVisibleInDirection(grid, visible, rows, cols, 0, i, 1, 0)
		markVisibleInDirection(grid, visible, rows, cols, rows-1, i, -1, 0)
	}
}

func countVisible(visible [][]bool, rows, cols int) int {
	result := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if visible[i][j] {
				result++
			}
		}
	}
	return result
}

func solve(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	visible := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visible[i] = make([]bool, cols)
	}

	markVisible(grid, visible, rows, cols)
	return countVisible(visible, rows, cols)
}

func main() {
	grid := parseInput()
	result := solve(grid)

	fmt.Println(result)
}
