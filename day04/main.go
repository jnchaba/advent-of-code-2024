package main

import (
	"advent-of-code-2024/pkg/timeutils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	timeutils.StartTimer("Advent of Code - Day 04")
	defer timeutils.TimeElapsed("Advent of Code - Day 04", true)
	grid := parseInputGrid()
	rows := len(grid)
	cols := len(grid[0])
	partOne(grid, rows, cols)
	partTwo(grid, rows, cols)
}

func parseInputGrid() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return nil
	}
	defer file.Close()
	grid := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowText := scanner.Text()
		row := strings.Split(rowText, "")
		grid = append(grid, row)
	}
	return grid
}

func partOne(grid [][]string, rowCount, colCount int) {
	var count int
	for _, row := range grid {
		count += scanArray(row)
	}
	cols := getCols(grid, rowCount, colCount)
	for _, col := range cols {
		count += scanArray(col)
	}
	tlbrDiagonals, trblDiagonals := getDiagonals(grid, rowCount, colCount)
	for _, diagonal := range tlbrDiagonals {
		count += scanArray(diagonal)
	}
	for _, diagonal := range trblDiagonals {
		count += scanArray(diagonal)
	}

	fmt.Printf("Total matches: %d\n", count)
}

func partTwo(grid [][]string, rowCount, colCount int) {
	count := countMasShapes(grid, rowCount, colCount)
	fmt.Printf("Total MAS crosses: %d\n", count)
}

func getCols(grid [][]string, rowCount, colCount int) [][]string {
	cols := make([][]string, rowCount)
	for i := range cols {
		col := make([]string, colCount)
		for j := range grid {
			col = append(col, grid[j][i])
		}
		cols = append(cols, col)
	}
	return cols
}

func getDiagonals(grid [][]string, rowCount, colCount int) ([][]string, [][]string) {
	var leftUpRightDiagonals [][]string
	var rightUpLeftDiagonals [][]string

	for i := 0; i < rowCount+colCount-1; i++ {
		var diagonalOne []string
		var diagonalTwo []string
		for j := 0; j < rowCount; j++ {
			k := i - j
			if k >= 0 && k < colCount {
				// leftUpRight
				diagonalOne = append(diagonalOne, grid[k][j])
				// rightUpLeft
				diagonalTwo = append(diagonalTwo, grid[k][colCount-j-1])
			}
		}
		if len(diagonalOne) > 0 {
			leftUpRightDiagonals = append(leftUpRightDiagonals, diagonalOne)
		}
		if len(diagonalTwo) > 0 {
			rightUpLeftDiagonals = append(rightUpLeftDiagonals, diagonalTwo)
		}
	}
	return leftUpRightDiagonals, rightUpLeftDiagonals
}

func scan(text string, target string) int {
	forwardPattern := `XMAS`
	backwardPattern := `SAMX`
	regex := regexp.MustCompile(forwardPattern)
	forwardMatches := len(regex.FindAllString(text, -1))
	regex = regexp.MustCompile(backwardPattern)
	backwardMatches := len(regex.FindAllString(text, -1))
	return forwardMatches + backwardMatches
}

func scanArray(array []string) int {
	stringified := strings.Join(array, "")
	return scan(stringified, "XMAS")
}

func isValidMasShape(grid [][]string, i, j int) bool {
	// check for oob first
	if i-1 < 0 || i+1 > len(grid) || j-1 < 0 || j+1 >= len(grid[0]) {
		return false
	}

	leftMas := (grid[i-1][j-1] == "M" && grid[i+1][j+1] == "S")
	leftSam := (grid[i-1][j-1] == "S" && grid[i+1][j+1] == "M")
	rightMas := (grid[i-1][j+1] == "S" && grid[i+1][j-1] == "M")
	rightSam := (grid[i-1][j+1] == "M" && grid[i+1][j-1] == "S")

	leftDiagonalCorrect := leftMas || leftSam
	rightDiagonalCorrect := rightMas || rightSam

	return leftDiagonalCorrect && rightDiagonalCorrect
}

func countMasShapes(grid [][]string, rowCount, colCount int) int {
	var count int
	for i := 1; i < rowCount-1; i++ {
		for j := 1; j < colCount-1; j++ {
			if grid[i][j] == "A" {
				if isValidMasShape(grid, i, j) {
					count++
				}
			}
		}
	}
	return count
}
