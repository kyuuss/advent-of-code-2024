package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func read_input_from_file(file_name string) ([]int, []int, error) {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var left, right []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Split the line by whitespace
		numbers := strings.Fields(scanner.Text())
		if len(numbers) < 2 {
			continue // Skip invalid lines
		}

		// Parse the first number (left column)
		leftNum, err := strconv.Atoi(numbers[0])
		if err != nil {
			continue // Skip invalid numbers
		}

		// Parse the second number (right column)
		rightNum, err := strconv.Atoi(numbers[1])
		if err != nil {
			continue // Skip invalid numbers
		}

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	return left, right, nil

}

func get_distance(left, right []int) int {
	sortedLeft := make([]int, len(left))
	sortedRight := make([]int, len(right))
	copy(sortedLeft, left)
	copy(sortedRight, right)

	// Sort the left and right columns
	// This is necessary to calculate the distance
	sort.Ints(sortedLeft)
	sort.Ints(sortedRight)

	totalDistance := 0
	for i := 0; i < len(sortedLeft); i++ {
		distance := int(math.Abs(float64(sortedRight[i] - sortedLeft[i])))
		totalDistance += distance
	}

	return totalDistance
}

func main() {
	right, left, err := read_input_from_file("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	distance := get_distance(right, left)
	similarity := get_similarity(right, left)
	fmt.Println(similarity)
	fmt.Println(distance)
}

func get_similarity(right, left []int) int {
	// Create a frequency map for the right list
	rightFreq := make(map[int]int)
	for _, num := range right {
		rightFreq[num]++
	}

	// Calculate similarity score
	totalScore := 0
	for _, leftNum := range left {
		// Multiply the number by its frequency in the right list
		frequency := rightFreq[leftNum] // Will be 0 if number doesn't exist in right list
		score := leftNum * frequency
		totalScore += score
	}

	return totalScore
}
