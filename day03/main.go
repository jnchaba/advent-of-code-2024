package main

import (
	"advent-of-code-2024/pkg/timeutils"
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	timeutils.StartTimer("Advent of Code - Day 03")
	defer timeutils.TimeElapsed("Advent of Code - Day 03", true)
	input := parseInputReports()
	partOne(input)
	partTwo(input)

}

func parseInputReports() string {
	timeutils.StartTimer("Reading input file")
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return ""
	}
	defer file.Close()

	scanner := bufio.NewReader(file)
	var input strings.Builder

	for {
		line, err := scanner.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				input.WriteString(line)
				break
			}
			fmt.Println("Error")
			return ""
		}
		input.WriteString(line)
	}
	return input.String()
}

func getMulCalls(input string, enableFlags bool) []string {
	pattern := `mul\(\d+,\d+\)`
	if enableFlags {
		pattern += `|do\(\)|don\'t\(\)`
	}
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllString(input, -1)
	return matches
}

func processMulCalls(mulCalls []string, enableFlags bool) int {
	sum := 0
	enabled := true
	for _, call := range mulCalls {
		if enableFlags {
			if call == "do()" {
				enabled = true
				continue
			} else if call == "don't()" {
				enabled = false
				continue
			}
		}
		if !enableFlags || enabled {
			pattern := `\d+,\d+`
			regex := regexp.MustCompile(pattern)
			matches := regex.FindAllString(call, -1)
			numOne, _ := strconv.Atoi(strings.Split(matches[0], ",")[0])
			numTwo, _ := strconv.Atoi(strings.Split(matches[0], ",")[1])
			sum += numOne * numTwo
		}
	}
	return sum
}

func partOne(input string) {
	mulCalls := getMulCalls(input, false)
	sum := processMulCalls(mulCalls, false)
	fmt.Println("Sum of mul calls: ", sum)
}

func partTwo(input string) {
	mulCallsWithEnableFlags := getMulCalls(input, true)
	sum := processMulCalls(mulCallsWithEnableFlags, true)
	fmt.Println("Sum of mul calls: ", sum)
}
