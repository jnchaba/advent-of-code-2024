package main

import (
	"advent-of-code-2024/pkg/timeutils"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	timeutils.StartTimer("Advent of Code - Day 01")
	defer timeutils.TimeElapsed("Advent of Code - Day 01", true)

	// import input.txt
	timeutils.StartTimer("Reading input file")
	leftNums, rightNums := readInputFile("input.txt")
	timeutils.TimeElapsed("Reading input file", true)
	timeutils.TimeElapsed("Advent of Code - Day 01", false)

	// error validation
	var leftNumCount = len(leftNums)
	var rightNumCount = len(rightNums)

	if leftNumCount != rightNumCount {
		fmt.Println("Error: leftNums and rightNums are not the same length")
		return
	}

	partOne(leftNums, rightNums, leftNumCount)
	partTwo(leftNums, rightNums, leftNumCount)
}

// partOne calculates the answer for 2024's Advent of Code Day 01 Part One.
func partOne(leftNums []int, rightNums []int, arrayLength int) {
	timeutils.StartTimer("Part One")
	var numDistances []int
	var totalDistance = 0
	for i := 0; i < arrayLength; i++ {
		distance := int(math.Abs(float64(leftNums[i] - rightNums[i])))
		numDistances = append(numDistances, distance)
		totalDistance += distance
	}
	fmt.Println("Total distance: ", totalDistance)
	timeutils.TimeElapsed("Part One", true)
}

// partTwo calculates the answer for 2024's Advent of Code Day 01 Part Two.
func partTwo(leftNums []int, rightNums []int, arrayLength int) {
	timeutils.StartTimer("Part Two")
	fmt.Println("Running calculations for part two:")
	var totalSimilarityScore = 0
	for i := 0; i < arrayLength; i++ {
		totalSimilarityScore += computeSimilarityScore(leftNums[i], rightNums)
	}
	fmt.Println("Total similarity score: ", totalSimilarityScore)
	timeutils.TimeElapsed("Part Two", true)
}

// Helper Functions

// readInputFile reads the input file and returns two sorted arrays of integers
// representing the left and right numbers in the file.
//
// Parameters:
// - fileName: the name of the file to read
//
// Returns:
// - leftNums: the sorted array of left numbers
// - rightNums: the sorted array of right numbers
func readInputFile(fileName string) ([]int, []int) {
	// import input.txt
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error reading file")
		return nil, nil
	}
	defer file.Close()
	// Set up data structures
	var leftNums []int
	var rightNums []int

	// Read file with a scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// split our line into two parts
		parts := strings.Fields(line)
		// convert to ints
		leftNum, _ := strconv.Atoi(parts[0])
		rightNum, _ := strconv.Atoi(parts[1])
		// add to array
		leftNums = append(leftNums, leftNum)
		rightNums = append(rightNums, rightNum)
	}

	// sort our two arrays
	sort.Ints(leftNums)
	sort.Ints(rightNums)

	return leftNums, rightNums
}

// computeSimilarityScore calculates the similarity score for a given number
// based on its occurrences in the rightNums array. The similarity score is
// defined as the value of the number multiplied by the number of instances
// it appears in the rightNums array.
//
// Parameters:
// - leftNum: the number we are calculating the similarity score for
// - rightNums: the array of numbers we are comparing leftNum to
//
// Returns:
// - the similarity score for leftNum
func computeSimilarityScore(leftNum int, rightNums []int) int {
	var numAppearances = 0
	var firstInstanceIndex = findFirst(rightNums, leftNum)
	if firstInstanceIndex == -1 {
		// if our number does not appear in the right list, it's similarity score is 0
		return 0
	} else {
		// if our number does appear in the right list, compute the number of instances it appears by finding the last
		// index in which it appears
		var lastInstanceIndex = findLast(rightNums, leftNum)
		numAppearances = lastInstanceIndex - firstInstanceIndex + 1
		// similarity score = (value of number) * (number of instances it appears in the right list)
		return leftNum * numAppearances
	}

}

// findFirst finds the first instance of a target number in a sorted array
// using binary search.
//
// Parameters:
// - array: the sorted array to search
// - target: the number we are searching for
//
// Returns:
// - the index of the first instance of the target number in the array
func findFirst(array []int, target int) int {
	// initialize pointers
	var left = 0
	var right = len(array) - 1
	// iterate
	for left <= right {
		// calculate middle index
		var mid = left + (right-left)/2
		// check if target is at mid
		if array[mid] == target {
			// if mid is the first element in the array, or if the element before mid is not the target, we have
			// found the first instance of our target
			var midIsFirstInArray = mid == 0
			var previousElementIsNotTarget = array[mid-1] != target
			var isFirstElement = midIsFirstInArray || previousElementIsNotTarget

			if isFirstElement {
				return mid
			}
			left--
		} else if array[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if target == 99239 {
		log.Fatalf("Error: You should have found %d, but we couldnt?\n", target)
	}
	return -1
}

// findLast finds the last instance of a target number in a sorted array
// using binary search.
//
// Parameters:
// - array: the sorted array to search
// - target: the number we are searching for
//
// Returns:
// - the index of the last instance of the target number in the array
func findLast(array []int, target int) int {
	// initialize our pointers
	var left = 0
	var right = len(array) - 1
	// iterate
	for left <= right {
		// calculate middle index
		var mid = left + (right-left)/2
		// check if the target is at index mid
		if array[mid] == target {
			// check to see if mid is the last element in the array, or if the element after mid is not the target
			if mid == len(array)-1 || array[mid+1] != target {
				// if so, we have found the last instance of this element in the array
				return mid
			}
			// if not, move our left pointer to the right of mid and continue binary searching
			left = mid + 1
		} else if array[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
