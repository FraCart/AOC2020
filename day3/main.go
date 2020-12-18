package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	localMap := readInput()

	partOneAnswer := partOne(localMap, 3, 1)
	fmt.Printf("Part one answer: %d\n", partOneAnswer)

	partTwo(localMap)
}

func readInput() []string {
	data, readFileErr := ioutil.ReadFile("input.txt")
	if readFileErr != nil {
		panic(readFileErr)
	}
	lines := strings.Split(string(data), "\n")
	return lines
}

func partOne(localMap []string, rightSlope, downSlope int) int {
	coordX, coordY := 0, 0
	numTreesEncountered := 0
	localMapWidth := len(localMap[0])

	for coordY < len(localMap) {
		currentValue := localMap[coordY][coordX]
		if currentValue == '#' {
			numTreesEncountered++
		}

		coordX = (coordX + rightSlope) % localMapWidth
		coordY += downSlope
	}

	return numTreesEncountered
}

func partTwo(localMap []string) {
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	slopeResults := []int{}

	for _, slope := range slopes {
		rightSlope, downSlope := slope[0], slope[1]
		numTreesEncountered := partOne(localMap, rightSlope, downSlope)
		slopeResults = append(slopeResults, numTreesEncountered)
	}

	answer := 1
	for _, slopeResult := range slopeResults {
		answer *= slopeResult
	}

	fmt.Printf("Part two answer: %d", answer)
}
