package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "unicode"
)

func createFileScanner(fileName string) (*bufio.Scanner, *os.File) {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatalf("unable to read file: %v", err)
    }

    reader := bufio.NewScanner(file)
    return reader, file
}

func parseInput() []string {
    scanner, file := createFileScanner("input.txt")
    defer file.Close()

    result := []string{}

    for scanner.Scan() {
        line := scanner.Text()
        result = append(result, line)
    }

    return result
}

func getCommonItem(left string, right string) rune {
    leftMap := map[rune]bool{}

    for _, val := range left {
        leftMap[val] = true
    }

    for _, val := range right {
        if leftMap[val] == true {
            return val
        }
    }
    return ' '
}

func getPriority(val rune) int {
    if unicode.IsLower(val) {
        return int(val - 'a' + 1)
    } else {
        return int(val - 'A' + 1 + 26)
    }
}

func solve(input []string) int {
    result := 0

    for _, rucksack := range input {
        left := rucksack[:len(rucksack) / 2]
        right := rucksack[len(rucksack) / 2:]

        commonItem := getCommonItem(left, right)
        result += getPriority(commonItem)
    }
    return result
}

func main() {
    input := parseInput()
    result := solve(input)

    fmt.Println(result)
}
