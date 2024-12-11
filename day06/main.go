package main

import (
	"advent-of-code-2024/pkg/arrayutils"
	"advent-of-code-2024/pkg/timeutils"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var GUARD_UP string = "^"
var GUARD_RIGHT string = ">"
var GUARD_DOWN string = "v"
var GUARD_LEFT string = "<"
var infiniteLoopPositions int
var stepCount int = 0

func main() {
	timeutils.StartTimer("Advent of Code - Day 06")
	defer timeutils.TimeElapsed("Advent of Code - Day 06", true)
	input, startRow, startCol := readInput()
	output := walk(input, startRow, startCol, false)
	count := findUniquePositions(output)
	fmt.Println("Part 1: ", count)

	tryInfiniteLoopPositions(output, startRow, startCol)
	fmt.Println("infinite loop positions: ", infiniteLoopPositions)
}

func readInput() ([][]string, int, int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Panic("error reading file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var startX int
	var startY int
	x := 0
	grid := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		if arrayutils.Contains(row, "^") {
			startX = x
			startY = arrayutils.IndexOf(row, "^")
		}
		x++
		grid = append(grid, row)
	}
	return grid, startX, startY
}

func walk(input [][]string, startRow int, startCol int, testForInfinite bool) [][]string {
	guard := GUARD_UP

	input = step(input, startRow, startCol, guard, testForInfinite)
	return input
}

func step(input [][]string, row int, col int, direction string, testForInfinite bool) [][]string {
	numRows := len(input)
	numCols := len(input[0])

	vector := getVectorBasedOnDirection(direction)
	newRow := row + vector[0]
	newCol := col + vector[1]

	if stepCount > 6000 {
		infiniteLoopPositions++
		return input
	}

	if newRow >= numRows || newRow < 0 || newCol >= numCols || newCol < 0 {
		input[row][col] = "X"
		return input
	} else {
		direction = turnUntilClear(input, direction, row, col)
		vector = getVectorBasedOnDirection(direction)
		newRow = row + vector[0]
		newCol = col + vector[1]
		input[row][col] = "X"
		input[newRow][newCol] = direction
		stepCount++
	}
	return step(input, newRow, newCol, direction, testForInfinite)
}

func getVectorBasedOnDirection(direction string) []int {
	vector := make([]int, 0)
	switch direction {
	case GUARD_UP:
		vector = []int{-1, 0}
	case GUARD_RIGHT:
		vector = []int{0, 1}
	case GUARD_DOWN:
		vector = []int{1, 0}
	case GUARD_LEFT:
		vector = []int{0, -1}
	}
	return vector
}

func turnUntilClear(input [][]string, direction string, row int, col int) string {
	vector := getVectorBasedOnDirection(direction)
	newRow := row + vector[0]
	newCol := col + vector[1]
	if input[newRow][newCol] == "#" {
		switch direction {
		case GUARD_UP:
			direction = GUARD_RIGHT
		case GUARD_RIGHT:
			direction = GUARD_DOWN
		case GUARD_DOWN:
			direction = GUARD_LEFT
		case GUARD_LEFT:
			direction = GUARD_UP
		}
		return turnUntilClear(input, direction, row, col)
	} else {
		return direction
	}
}

func findUniquePositions(input [][]string) int {
	count := 0
	for i := 0; i < len(input); i++ {
		inputString := strings.Join(input[i], "")
		pattern := `X`
		regex := regexp.MustCompile(pattern)
		count += len(regex.FindAllString(inputString, -1))
	}
	return count
}

func tryInfiniteLoopPositions(input [][]string, startRow, startCol int) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] != "^" && input[i][j] != "#" && input[i][j] != "." {
				stepCount = 0
				store := input[i][j]
				input[i][j] = "#"
				walk(input, startRow, startCol, true)
				input[i][j] = store
			}
		}
	}
}
