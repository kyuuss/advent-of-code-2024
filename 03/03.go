package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("03/input.txt")

	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}

	var builder strings.Builder
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fileString := builder.String()

	pattern := `mul\((\d{1,3}),(\d{1,3})\)`

	r := regexp.MustCompile(pattern)
	matches := r.FindAllStringSubmatch(fileString, -1) // Returns e.g. [[mul(1,2) 1 2] [mul(3,4) 3 4]]

	fmt.Printf("Matches: %v\n", matches)
	if err != nil {
		log.Fatalf("Error compiling regex: %v\n", err)
	}

	sum := 0

	for _, match := range matches {
		num1, err := strconv.Atoi(match[1])

		if err != nil {
			log.Fatalf("Error converting %s to int: %v\n", match[1], err)
		}

		num2, err := strconv.Atoi(match[2])

		if err != nil {
			log.Fatalf("Error converting %s to int: %v\n", match[2], err)
		}

		sum += num1 * num2
	}

	fmt.Printf("Sum: %d\n", sum)

}
