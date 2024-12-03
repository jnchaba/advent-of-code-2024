package main

import (
	"advent-of-code-2024/pkg/timeutils"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	timeutils.StartTimer("Advent of Code - Day 02")
	defer timeutils.TimeElapsed("Advent of Code - Day 02", true)
	var reports = parseInputReports()
	timeutils.TimeElapsed("Reading input file", true)
	var safeReports = partOne(reports)
	timeutils.TimeElapsed("Computing Part One Answer", true)
	fmt.Println("Safe reports: ", safeReports)
}

func partOne(reports [][]int) int {
	timeutils.StartTimer("Computing Part One Answer")
	partTwo := true
	var safeReports = 0
	for _, report := range reports {
		if isSafeReport(report) {
			safeReports++
		} else if partTwo {
			for i := range len(report) {
				if dampenReportAndRecheck(report, i) {
					safeReports++
					break
				}
			}
		}
	}
	return safeReports
}

func parseInputReports() [][]int {
	timeutils.StartTimer("Reading input file")
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return nil
	}
	defer file.Close()

	var reports [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var report []int
		reportText := scanner.Text()
		stringLevels := strings.Fields(reportText)

		for _, level := range stringLevels {
			numLevel, err := strconv.Atoi(level)
			if err != nil {
				log.Fatalf("Error converting %s to int", level)
			}
			report = append(report, numLevel)
		}
		reports = append(reports, report)
	}
	return reports
}

func isSafeReport(report []int) bool {
	prev := 0
	increasing := false

	for i, level := range report {

		if i == 0 {
			prev = level
			continue
		}

		var delta = math.Abs(float64(level - prev))
		if delta == 0 || delta > 3 {
			return false
		}

		if i == 1 {
			increasing = level > prev
			prev = level
			continue
		}

		if (level > prev) != increasing {
			return false
		}

		prev = level
	}
	return true
}

func dampenReportAndRecheck(report []int, index int) bool {
	var dampenedReport []int
	dampenedReport = append(dampenedReport, report[:index]...)
	dampenedReport = append(dampenedReport, report[index+1:]...)
	var safeAfterDampening = isSafeReport(dampenedReport)
	return safeAfterDampening
}
