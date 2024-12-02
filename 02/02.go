package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isReportSafe(levels []int) bool {
	// Check if there are less than 2 levels
	// If there are less than 2 levels, the report is safe
	if len(levels) < 2 {
		return true
	}

	// Check first difference to determine if we should be increasing or decreasing
	increasing := levels[1] > levels[0]

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		// Check if difference is between 1 and 3 (inclusive)
		if absoluteNumber(diff) < 1 || absoluteNumber(diff) > 3 {
			return false
		}

		// Check if direction matches the initial direction
		if increasing && diff <= 0 {
			return false
		}
		if !increasing && diff >= 0 {
			return false
		}
	}

	return true
}

// Returns the absolute value of x
// e.g. absoluteNumber(-5) = 5
func absoluteNumber(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, err := os.Open("/02/input.txt")

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	scanner := bufio.NewScanner(file)
	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Parse the numbers from the line
		numStrs := strings.Fields(line)
		levels := make([]int, len(numStrs))

		for i, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Printf("Error parsing number: %v\n", err)
				continue
			}
			levels[i] = num
		}

		if isReportSafe(levels) {
			safeCount++
		}
	}

	fmt.Printf("Number of safe reports: %d\n", safeCount)
}
