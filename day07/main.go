package main

import (
	"advent-of-code-2024/pkg/timeutils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	timeutils.StartTimer("Advent of Code - Day 07")
	defer timeutils.TimeElapsed("Advent of Code - Day 07", true)
	var calibrationValue int
	equations := readInput()
	for i := 0; i < len(equations); i++ {
		calibrationValue += equationCanBeSolved(equations[i])
	}
	fmt.Printf("Calibration Value: %d\n", calibrationValue)
}

func readInput() [][]int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Panic("error reading file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	equations := make([][]int, 0)
	for scanner.Scan() {
		equation := make([]int, 0)
		equationText := scanner.Text()
		parts := strings.Fields(equationText)

		for index, part := range parts {
			if index == 0 {
				part = part[:len(part)-1]
			}

			number, err := strconv.Atoi(part)
			if err != nil {
				log.Fatalf("Error converting %s to int", number)
			}
			equation = append(equation, number)
		}
		equations = append(equations, equation)
	}
	return equations
}

func equationCanBeSolved(equation []int) int {
	if len(equation) < 2 {
		return 0
	}
	if canReachTarget(equation[0], equation[1:], 0) {
		return equation[0]
	} else {
		return 0
	}
}

func canReachTarget(target int, numbers []int, currentValue int) bool {
	if len(numbers) == 0 {
		return currentValue == target
	}
	next, remaining := numbers[0], numbers[1:]
	return canReachTarget(target, remaining, currentValue+next) ||
		canReachTarget(target, remaining, currentValue*next) ||
		canReachTarget(target, remaining, concatenateInts(currentValue, next))
}

func concatenateInts(a, b int) int {
	concatenatedString := strconv.Itoa(a) + strconv.Itoa(b)
	concatenatedInt, err := strconv.Atoi(concatenatedString)
	if err != nil {
		log.Fatalf("Error converting concatenated string to int: %s", concatenatedString)
	}
	return concatenatedInt
}
