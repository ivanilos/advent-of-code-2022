package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pt struct {
	x int
	y int
	z int
}

var dx = []int{1, -1, 0, 0, 0, 0}
var dy = []int{0, 0, 1, -1, 0, 0}
var dz = []int{0, 0, 0, 0, 1, -1}

const sides = 6
const minCoord = -1
const maxCoord = 20

func createFileScanner(fileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	reader := bufio.NewScanner(file)
	return reader, file
}

func parseInput() []Pt {
	scanner, file := createFileScanner("input.txt")
	defer file.Close()

	dropletParts := []Pt{}

	for scanner.Scan() {
		line := scanner.Text()

		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)

		dropletParts = append(dropletParts, Pt{x, y, z})
	}
	return dropletParts
}

func getPartsMap(dropletParts []Pt) map[Pt]bool {
	result := map[Pt]bool{}

	for _, dropletPart := range dropletParts {
		result[dropletPart] = true
	}
	return result
}

func isInsideCoordinates(x, y, z int) bool {
	return minCoord <= x && x <= maxCoord &&
		minCoord <= y && y <= maxCoord &&
		minCoord <= z && z <= maxCoord
}

func DFS(x, y, z int, outsidePoints *map[Pt]bool, partsMap *map[Pt]bool, result *int) {
	pt := Pt{x, y, z}
	(*outsidePoints)[pt] = true

	for side := 0; side < sides; side++ {
		nx := x + dx[side]
		ny := y + dy[side]
		nz := z + dz[side]

		newPt := Pt{nx, ny, nz}

		if _, found := (*partsMap)[newPt]; found {
			*result++
		} else if _, found := (*outsidePoints)[newPt]; !found && isInsideCoordinates(nx, ny, nz) {
			DFS(nx, ny, nz, outsidePoints, partsMap, result)
		}
	}
}

// assumes coordinates are small
func solve(dropletParts []Pt) int {
	result := 0
	partsMap := getPartsMap(dropletParts)
	outsidePoints := map[Pt]bool{}

	DFS(minCoord, minCoord, minCoord, &outsidePoints, &partsMap, &result)

	return result
}

func main() {
	droplets := parseInput()
	result := solve(droplets)

	fmt.Println(result)
}
