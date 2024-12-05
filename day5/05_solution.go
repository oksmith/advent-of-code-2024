package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

type Rules [100][100]bool

func main() {
	fmt.Printf("Part 1 result: %d\n", partone())
	// fmt.Printf("Part 2 result: %d\n", parttwo())
}

func partone() int {
	rules, updates := parseInput()
	ruleTable := ruleTable(rules)

	counter := 0
	for _, update := range updates {
		valid := isCorrectlyOrdered(update, ruleTable)
		if valid {
			middleNumber := update[len(update)/2]
			counter += middleNumber
		}
	}

	return counter
}

func parttwo() int {
	rules, updates := parseInput()
	ruleTable := ruleTable(rules)

	counter := 0
	for _, update := range updates {
		valid := isCorrectlyOrdered(update, ruleTable)
		if valid {
			middleNumber := update[len(update)/2]
			counter += middleNumber
		} else {
			orderedUpdate := correctlyOrder(update, ruleTable)
			middleNumber := orderedUpdate[len(orderedUpdate)/2]
			counter += middleNumber
		}
	}

	return counter
}

func correctlyOrder(update []int, ruleTable Rules) []int {
	// hard to do?
	return update
}

func isCorrectlyOrdered(update []int, ruleTable Rules) bool {
	valid := true

	for i, num := range update {
		if i == 0 {
			continue
		}
		for j := 0; j < i; j++ {
			if !ruleTable[update[j]][num] {
				return false
			}
		}
	}
	return valid
}

func ruleTable(rules [][]int) Rules {
	// I think it would be more efficient to first create a dictionary that's very quick
	// to index, rather than looping through all rules for each update and checking the first
	// element
	var ruleTable Rules
	for _, rule := range rules {
		ruleTable[rule[0]][rule[1]] = true
	}
	return ruleTable
}

func parseInput() ([][]int, [][]int) {
	_, goFilename, _, _ := runtime.Caller(1)
	dataFilename := path.Join(path.Dir(goFilename), "05_input.txt")
	input, _ := os.ReadFile(dataFilename)
	blocks := strings.Split(string(input), "\n\n")
	if len(blocks) != 2 {
		fmt.Println("Invalid input format")
		os.Exit(1)
	}
	rules := strings.Split(blocks[0], "\n")
	updates := strings.Split(blocks[1], "\n")

	var parsedRules [][]int
	for _, rule := range rules {
		parsedRule := strings.Split(rule, "|")
		var parsedRuleIntegers []int
		for _, element := range parsedRule {
			elementInt, _ := strconv.Atoi(element)
			parsedRuleIntegers = append(parsedRuleIntegers, elementInt)
		}
		parsedRules = append(parsedRules, parsedRuleIntegers)
	}

	var parsedUpdates [][]int
	for _, update := range updates {
		parsedUpdate := strings.Split(update, ",")
		var parsedUpdateIntegers []int
		for _, element := range parsedUpdate {
			elementInt, _ := strconv.Atoi(element)
			parsedUpdateIntegers = append(parsedUpdateIntegers, elementInt)
		}
		parsedUpdates = append(parsedUpdates, parsedUpdateIntegers)
	}
	return parsedRules, parsedUpdates
}
