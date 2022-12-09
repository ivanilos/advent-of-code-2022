package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Dir struct {
	name     string
	par      *Dir
	children map[string]*Dir
	data     map[string]int
}

func createFileScanner(fileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	reader := bufio.NewScanner(file)
	return reader, file
}

func processCDOperation(lineFields []string, curNode *Dir) *Dir {
	name := lineFields[2]

	if name == ".." {
		curNode = curNode.par
	} else if _, ok := curNode.children[name]; ok {
		curNode = curNode.children[name]
	} else {
		curNode.children[name] = &Dir{name, curNode, map[string]*Dir{}, map[string]int{}}
		curNode = curNode.children[name]
	}

	return curNode
}

func processData(lineFields []string, curNode *Dir) {
	name := lineFields[1]
	if lineFields[0] == "dir" {
		curNode.children[name] = &Dir{name, curNode, map[string]*Dir{}, map[string]int{}}
	} else {
		sz, _ := strconv.Atoi(lineFields[0])
		curNode.data[name] = sz
	}
}

func parseInput() *Dir {
	scanner, file := createFileScanner("input.txt")
	defer file.Close()

	fileTreeRoot := &Dir{
		name:     "/",
		children: map[string]*Dir{},
		data:     map[string]int{},
	}
	fileTreeRoot.par = fileTreeRoot
	curNode := fileTreeRoot

	// skip first command "$ cmd /"
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		lineFields := strings.Fields(line)

		if lineFields[0] == "$" {
			if lineFields[1] == "cd" {
				curNode = processCDOperation(lineFields, curNode)
			}
		} else {
			processData(lineFields, curNode)
		}
	}
	return fileTreeRoot
}

func DFS(curDir *Dir, directoriesSizes *[]int) int {
	sz := 0

	for _, child := range curDir.children {
		sz += DFS(child, directoriesSizes)
	}

	for _, datumSize := range curDir.data {
		sz += datumSize
	}
	*directoriesSizes = append(*directoriesSizes, sz)
	return sz
}

func getMinSizeToDelete(directoriesSizes []int) int {
	const fileSystemSpace = 70000000
	const neededSpace = 30000000

	sort.Ints(directoriesSizes)

	freeSpace := fileSystemSpace - directoriesSizes[len(directoriesSizes)-1]

	for _, sz := range directoriesSizes {
		if freeSpace+sz >= neededSpace {
			return sz
		}
	}
	return -1
}

func solve(fileTreeRoot *Dir) int {
	const MAX_SIZE = 100000

	directoriesSizes := []int{}
	DFS(fileTreeRoot, &directoriesSizes)

	return getMinSizeToDelete(directoriesSizes)
}

func main() {
	fileTreeRoot := parseInput()
	result := solve(fileTreeRoot)

	fmt.Println(result)
}
