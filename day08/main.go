package main

import (
	"advent-of-code-2024/pkg/arrayutils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Objective: Find all unique points where antinodes can occur
// In particular, an antinode occurs at any point that is
// perfectly in line with two antennas of the same frequency -
// but only when one of the antennas is twice as far away as the other.
// This means that for any pair of antennas with the same frequency,
// there are two antinodes, one on either side of them.

type Point struct {
	x int
	y int
}

type AntennaCategory struct {
	label     string
	positions []Point
}

// create an array that will be used to store categories of antennas and their positions
var antennaCategories = make([]AntennaCategory, 0)

func main() {
	input := readInput()
	fmt.Printf("\nInput len: %d.", len(input))
	mapAntennas(input)
	mWidth, mHeight := len(input), len(input[0])
	//fmt.Printf("\nAntenna categories: %v", antennaCategories)
	partTwo := true
	antiNodePositions := findAntinodePositions(antennaCategories, mWidth, mHeight, partTwo)
	fmt.Printf("\nUnique Positions: %d\n", len(antiNodePositions))
}

func readInput() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Panic("error reading file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	input := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		fmt.Println(row)
		input = append(input, row)
	}
	return input
}

func mapAntennas(grid [][]string) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "." {
				continue
			}
			// check if the category already exists
			categoryExists := false
			for k := 0; k < len(antennaCategories); k++ {
				if antennaCategories[k].label == grid[i][j] {
					categoryExists = true
					antennaCategories[k].positions = append(antennaCategories[k].positions, Point{x: i, y: j})
					break
				}
			}
			if !categoryExists {
				antennaCategories = append(antennaCategories, AntennaCategory{label: grid[i][j], positions: []Point{Point{x: i, y: j}}})
			}
		}
	}
}

func findAntinodePositions(antennas []AntennaCategory, mWidth int, mHeight int, partTwo bool) []Point {
	antinodePositions := make([]Point, 0)
	for _, antenna := range antennas {
		for i := 0; i < len(antenna.positions); i++ {
			for j := 0; j < len(antenna.positions); j++ {
				if i == j {
					continue
				}

				// calculate the dx and dy to find the "line" of positions in line with our two antennas
				dx := antenna.positions[j].x - antenna.positions[i].x
				dy := antenna.positions[j].y - antenna.positions[i].y

				// calculate a possible position that fits our parameters of being at least twice the distance
				// from one antenna compared to the other
				antinodeX := antenna.positions[i].x + 2*dx
				antinodeY := antenna.positions[i].y + 2*dy

				if antinodeX >= 0 && antinodeX < mWidth && antinodeY >= 0 && antinodeY < mHeight {
					antinode := Point{x: antinodeX, y: antinodeY}
					if !arrayutils.Contains(antinodePositions, antinode) {
						antinodePositions = append(antinodePositions, Point{x: antinodeX, y: antinodeY})

					}
				}

				// if we are looking for part2 answer, we need to iterate over the entire width/height of the
				// input grid along the "line" between our two antennas
				if partTwo {
					for k := 1; k < mWidth || k < mHeight; k++ {
						//fmt.Printf("Dx: %d, Dy: %d\n", dx, dy)
						antinodeX = antenna.positions[i].x + k*dx
						antinodeY = antenna.positions[i].y + k*dy

						if antinodeX >= 0 && antinodeX < mWidth && antinodeY >= 0 && antinodeY < mHeight {
							antinode := Point{x: antinodeX, y: antinodeY}
							if !arrayutils.Contains(antinodePositions, antinode) {
								antinodePositions = append(antinodePositions, Point{x: antinodeX, y: antinodeY})
							}
						} else {
							break
						}
					}
				}
			}
		}
	}
	return antinodePositions
}
