package main

import (
	"advent-of-code-2024/pkg/arrayutils"
	"advent-of-code-2024/pkg/timeutils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	timeutils.StartTimer("Advent of Code - Day 05")
	defer timeutils.TimeElapsed("Advent of Code - Day 05", true)
	rules, updates := parseInput()
	validUpdates := make([][]int, 0)
	invalidUpdates := make([][]int, 0)
	for _, update := range updates {
		if verifyUpdate(update, rules) {
			validUpdates = append(validUpdates, update)
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}
	sum := sumMiddlePages(validUpdates)
	fmt.Println("Sum of middle pages of initially correct updates: ", sum)

	for _, update := range invalidUpdates {
		update = correctUpdate(update, rules)
	}
	correctedSum := sumMiddlePages(invalidUpdates)
	fmt.Println("Sum of middle pages of corrected updates: ", correctedSum)
	return
}

func parseInput() ([][]int, [][]int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Panic("Error reading file")
	}
	defer file.Close()

	pageOrderingRules := make([][]int, 0)
	updates := make([][]int, 0)
	var isUpdateSection bool

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isUpdateSection = true
			continue
		}

		if !isUpdateSection {
			parts := strings.Split(line, "|")
			rule := make([]int, 2)
			rule[0], _ = strconv.Atoi(parts[0])
			rule[1], _ = strconv.Atoi(parts[1])
			pageOrderingRules = append(pageOrderingRules, rule)
		} else {
			line = strings.Trim(line, "()")
			parts := strings.Split(line, ",")
			update := make([]int, len(parts))
			for i, part := range parts {
				update[i], _ = strconv.Atoi(part)
			}
			updates = append(updates, update)
		}
	}
	return pageOrderingRules, updates
}

func getApplicableRulesForUpdate(update []int, rules [][]int) [][]int {
	applicableRules := make([][]int, 0)
	for i := 1; i < len(rules); i++ {
		if arrayutils.Contains(update, rules[i][0]) && arrayutils.Contains(update, rules[i][1]) {
			applicableRules = append(applicableRules, rules[i])
		}
	}
	return applicableRules
}

func verifyUpdate(update []int, rules [][]int) bool {
	applicableRules := getApplicableRulesForUpdate(update, rules)
	for _, rule := range applicableRules {
		if isRuleBroken(update, rule) {
			return false
		}
	}
	return true
}

func isRuleBroken(update []int, rule []int) bool {
	return arrayutils.IndexOf(update, rule[0]) >= arrayutils.IndexOf(update, rule[1])
}

func sumMiddlePages(validUpdates [][]int) int {
	var sum int
	for _, update := range validUpdates {
		middle := len(update) / 2
		sum += update[middle]
	}
	return sum
}

func correctUpdate(update []int, rules [][]int) []int {
	applicableRules := getApplicableRulesForUpdate(update, rules)
	for _, rule := range applicableRules {
		update = applyRule(update, rule)
	}
	if verifyUpdate(update, rules) {
		return update
	} else {
		return correctUpdate(update, rules)
	}
}

func applyRule(update []int, rule []int) []int {
	if !isRuleBroken(update, rule) {
		return update
	} else {
		illegalVal, newUpdate := arrayutils.RemoveAtIndex(update, arrayutils.IndexOf(update, rule[1]))
		newUpdate = arrayutils.InsertAtIndex(newUpdate, arrayutils.IndexOf(newUpdate, rule[0])+1, illegalVal)
		return newUpdate
	}

}
