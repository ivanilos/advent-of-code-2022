package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Valve struct {
	flow     int
	adjacent []string
}

type State struct {
	nameMe           string
	nameElephant     string
	visitedMask      int
	timeLeftMe       int
	timeLeftElephant int
}

const INF = int(1e9)

func createFileScanner(fileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	reader := bufio.NewScanner(file)
	return reader, file
}

func parseInput() map[string]Valve {
	scanner, file := createFileScanner("input.txt")
	defer file.Close()

	valves := map[string]Valve{}

	for scanner.Scan() {
		line := scanner.Text()

		r := regexp.MustCompile(`Valve ([A-Z]+) has flow rate=([0-9]+); tunnels? leads? to valves? ([A-Z, ]+)`)
		matches := r.FindAllStringSubmatch(line, -1)

		name := matches[0][1]
		flow, _ := strconv.Atoi(matches[0][2])
		adjacentString := strings.ReplaceAll(matches[0][3], " ", "")
		adjacent := strings.Split(adjacentString, ",")

		valves[name] = Valve{flow, adjacent}
	}
	return valves
}

// ValveID := 26 * char[0] + char[1]
func intToValveID(val int) string {
	right := 'A' + val%26
	left := 'A' + (val / 26)
	return string(left) + string(right)
}

// ValveID := 26 * char[0] + char[1]
func valveIDToInt(valveID string) int {
	return 26*int(valveID[0]-'A') + int(valveID[1]-'A')
}

func createInitialDistancesMatrix(valves map[string]Valve, maxValves int) [][]int {
	distances := make([][]int, maxValves)
	for i := 0; i < maxValves; i++ {
		distances[i] = make([]int, maxValves)
		for j := 0; j < maxValves; j++ {
			distances[i][j] = INF
		}
		distances[i][i] = 0
	}

	for name, valve := range valves {
		for _, adjacentValve := range valve.adjacent {
			a := valveIDToInt(name)
			b := valveIDToInt(adjacentValve)

			distances[a][b] = 1
		}
	}
	return distances
}

func getDistances(valves map[string]Valve) [][]int {
	maxValves := 26*26 + 25

	distances := createInitialDistancesMatrix(valves, maxValves)

	for k := 0; k < maxValves; k++ {
		for i := 0; i < maxValves; i++ {
			for j := 0; j < maxValves; j++ {
				distances[i][j] = min(distances[i][j], distances[i][k]+distances[k][j])
			}
		}
	}
	return distances
}

func getValvesWithPositiveFlow(valves map[string]Valve) []string {
	result := []string{}

	for name, valve := range valves {
		if valve.flow > 0 {
			result = append(result, name)
		}
	}
	return result
}

func visited(visitedMask, idx int) bool {
	bit := (visitedMask >> idx) & 1
	return bit == 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calc(valveNameMe string, valveNameElephant string,
	visitedMask int,
	timeLeftMe int, timeLeftElephant int,
	interestingValves []string, distances [][]int, valves map[string]Valve,
	dp map[State]int) int {

	if timeLeftMe < timeLeftElephant {
		timeLeftMe, timeLeftElephant = timeLeftElephant, timeLeftMe
		valveNameMe, valveNameElephant = valveNameElephant, valveNameMe
	}

	state := State{valveNameMe, valveNameElephant, visitedMask, timeLeftMe, timeLeftElephant}
	if _, found := dp[state]; found {
		return dp[state]
	}

	if timeLeftMe <= 0 {
		return 0
	}

	dp[state] = 0
	curValveID := valveIDToInt(valveNameMe)
	for i := 0; i < len(interestingValves); i++ {
		if !visited(visitedMask, i) {
			nextValve := interestingValves[i]
			nextValveID := valveIDToInt(nextValve)
			dist := distances[curValveID][nextValveID]

			if timeLeftMe-dist-1 >= 0 {
				nextTimeLeft := timeLeftMe - dist - 1
				score := nextTimeLeft * valves[nextValve].flow
				nextVisitedMask := visitedMask | (1 << i)

				aux := score + calc(nextValve, valveNameElephant,
					nextVisitedMask,
					nextTimeLeft, timeLeftElephant,
					interestingValves, distances, valves, dp)

				dp[state] = max(dp[state], aux)
			}
		}
	}
	return dp[state]
}

// assumes only a few valves have positive flow
func solve(valves map[string]Valve) int {
	distances := getDistances(valves)

	interestingValves := getValvesWithPositiveFlow(valves)

	const totalTime = 26

	dp := map[State]int{}
	return calc("AA", "AA", 0, totalTime, totalTime,
		interestingValves, distances, valves, dp)
}

func main() {
	valves := parseInput()
	result := solve(valves)

	fmt.Println(result)
}
