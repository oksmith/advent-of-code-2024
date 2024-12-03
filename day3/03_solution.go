package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1 result: %d\n", partone())
	// fmt.Printf("Part 2 result: %d\n", parttwo())
}

func partone() int {
	data := readInput()

	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`) // must be a 1-3 digit number?
	matches := r.FindAllString(data, -1)

	result := 0
	for _, match := range matches {
		// take the firstNumber,secondNumber substring
		subStr := match[4 : len(match)-1]

		digitsStr := strings.Split(subStr, ",")
		firstNumber, _ := strconv.Atoi(digitsStr[0])
		secondNumber, _ := strconv.Atoi(digitsStr[1])
		result += firstNumber * secondNumber
	}
	return result
}

func readInput() string {
	_, goFilename, _, _ := runtime.Caller(1)
	dataFilename := path.Join(path.Dir(goFilename), "03_input.txt")
	input, _ := os.ReadFile(dataFilename)
	return string(input)
}
