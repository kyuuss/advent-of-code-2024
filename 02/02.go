package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	list := readFromFile("input.txt")

	number := getSafeReports(list)
	fmt.Println(number)
}

// Returns the absolute value of a number
// e.g. -5 -> 5
func absoluteValue(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func getSafeReports(list [][]int) int {
	safeReports := 0
	isReportSafe := true
	for _, row := range list {
		for j, num := range row {
			currentNumber := num

			if j == len(row)-1 {
				break
			}
			nextNumber := row[j+1]

			increasing := currentNumber < nextNumber

			diff := nextNumber - currentNumber
			diffAbsolute := absoluteValue(diff)

			// If the difference is greater than 3 or less than 1, the report is not safe
			if diffAbsolute > 3 || diffAbsolute < 1 {
				isReportSafe = false
			}

			// If the difference is negative and the numbers are increasing, or the difference is positive and the numbers are decreasing, the report is not safe

			if !increasing && diff >= 0 {
				isReportSafe = false
			}

			// If the difference is positive and the numbers are increasing, or the difference is negative and the numbers are decreasing, the report is not safe
			if increasing && diff <= 0 {
				isReportSafe = false
			}

		}

		if isReportSafe {
			safeReports++
		} else {
			isReportSafe = true
		}
	}

	return safeReports
}

func readFromFile(s string) [][]int {
	// Open the file
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	list := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		listLine := make([]int, 0)
		for _, v := range line {
			listLine = append(listLine, int(v))
		}

		list = append(list, listLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return list
}
