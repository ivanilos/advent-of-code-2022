package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type pos struct {
	x int
	y int
}

type Queue struct {
	items []pos
}

func (q *Queue) Push(x, y int) {
	q.items = append(q.items, pos{x, y})
}

func (q *Queue) Pop() (int, int) {
	val := q.items[0]
	q.items = q.items[1:]

	return val.x, val.y
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
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

func findPositionWith(char rune, grid [][]rune) (int, int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == char {
				return i, j
			}
		}
	}
	return -1, -1
}

func findAllPositionsWith(char rune, grid [][]rune) []pos {
	result := []pos{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == char {
				result = append(result, pos{i, j})
			}
		}
	}
	return result
}

func isIn(x, y, rows, cols int) bool {
	return 0 <= x && x < rows && 0 <= y && y < cols
}

func canMove(grid [][]rune, x, y, nx, ny int) bool {
	from := grid[x][y]
	to := grid[nx][ny]

	return to-from <= 1
}

func BFS(grid [][]rune, start []pos, ex, ey int) int {
	const INF = int(1e9)

	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}

	rows := len(grid)
	cols := len(grid[0])

	dist := make([][]int, rows)
	for row := 0; row < rows; row++ {
		dist[row] = make([]int, cols)
		for col := 0; col < cols; col++ {
			dist[row][col] = INF
		}
	}

	queue := Queue{}

	for _, position := range start {
		sx, sy := position.x, position.y
		queue.Push(sx, sy)
		dist[sx][sy] = 0
	}

	for !queue.IsEmpty() {
		x, y := queue.Pop()

		for dir := 0; dir < 4; dir++ {
			nx, ny := x+dx[dir], y+dy[dir]

			if isIn(nx, ny, rows, cols) && canMove(grid, x, y, nx, ny) && dist[nx][ny] > 1+dist[x][y] {
				dist[nx][ny] = 1 + dist[x][y]
				queue.Push(nx, ny)
			}
		}
	}
	return dist[ex][ey]
}

func solve(grid [][]rune) int {
	sx, sy := findPositionWith('S', grid)
	ex, ey := findPositionWith('E', grid)

	grid[sx][sy] = 'a'
	grid[ex][ey] = 'z'

	start := findAllPositionsWith('a', grid)

	return BFS(grid, start, ex, ey)
}

func main() {
	grid := parseInput()
	result := solve(grid)

	fmt.Println(result)
}
