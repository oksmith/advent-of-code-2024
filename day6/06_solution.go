package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

type Rules [100][100]bool

func main() {
	fmt.Printf("Part 1 result: %d\n", partone())
	// fmt.Printf("Part 2 result: %d\n", parttwo())
}

func partone() int {
	data := parseInput()

	directions := map[int][2]int{
		0: {-1, 0}, // up
		1: {0, 1},  // right
		2: {1, 0},  // down
		3: {0, -1}, // left
	}
	direction := 0 // first element of directions i.e. up

	// find starting position
	x, y := findStartPos(data)
	// fmt.Println(x, y)

	// initialise visitedLocations with the starting position
	visitedLocations := map[string]bool{
		fmt.Sprintf("%d,%d", x, y): true,
	}
	for x > 0 && x < len(data)-1 && y > 0 && y < len(data)-1 {
		x, y, direction = move(data, x, y, direction, directions, visitedLocations)
		// fmt.Println("direction", direction, "which is", directions[direction][0], directions[direction][1])
		// fmt.Println(x, y)
	}
	return len(visitedLocations)
}

func move(data [][]string, x, y, direction int, directions map[int][2]int, visitedLocations map[string]bool) (int, int, int) {
	if data[x+directions[direction][0]][y+directions[direction][1]] == "#" {
		// fmt.Println("identified obstruction at", x+directions[direction][0], y+directions[direction][1])
		// if the next position in current direction is an obstruction
		// move direction 90 degrees and do nothing else
		direction = (direction + 1) % 4
	} else {
		// assume all other elements of the array are "." or "^" (starting position), then
		// we move in the current direction and append the new position to visitedLocations
		x += directions[direction][0]
		y += directions[direction][1]
		visitedLocations[fmt.Sprintf("%d,%d", x, y)] = true
		// fmt.Println("moved to", x, y)
	}
	return x, y, direction
}

func findStartPos(data [][]string) (int, int) {
	for i, row := range data {
		for j, col := range row {
			if col == "^" {
				return i, j
			}
		}
	}
	return -1, -1
}

func parseInput() [][]string {
	_, goFilename, _, _ := runtime.Caller(1)
	dataFilename := path.Join(path.Dir(goFilename), "06_input.txt")
	input, _ := os.ReadFile(dataFilename)
	lines := strings.Split(string(input), "\n")

	var parsedInput [][]string
	for _, line := range lines {
		if line != "" {
			parsedLine := strings.Split(line, "")
			parsedInput = append(parsedInput, parsedLine)
		}
	}

	return parsedInput
}
