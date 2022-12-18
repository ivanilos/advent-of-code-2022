package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type DropletPart struct {
	x int
	y int
	z int
}

func createFileScanner(fileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	reader := bufio.NewScanner(file)
	return reader, file
}

func parseInput() []DropletPart {
	scanner, file := createFileScanner("input.txt")
	defer file.Close()

	dropletParts := []DropletPart{}

	for scanner.Scan() {
		line := scanner.Text()

		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)

		dropletParts = append(dropletParts, DropletPart{x, y, z})
	}
	return dropletParts
}

// assumes no droplets are coincident
func solve(dropletParts []DropletPart) int {
	result := 0
	seen := map[DropletPart]bool{}

	dx := []int{1, -1, 0, 0, 0, 0}
	dy := []int{0, 0, 1, -1, 0, 0}
	dz := []int{0, 0, 0, 0, 1, -1}
	const sides = 6

	for _, dropletPart := range dropletParts {
		for side := 0; side < sides; side++ {
			nx := dropletPart.x + dx[side]
			ny := dropletPart.y + dy[side]
			nz := dropletPart.z + dz[side]

			if _, found := seen[DropletPart{nx, ny, nz}]; found {
				result--
			} else {
				result++
			}
		}
		seen[dropletPart] = true
	}
	return result
}

func main() {
	droplets := parseInput()
	result := solve(droplets)

	fmt.Println(result)
}
