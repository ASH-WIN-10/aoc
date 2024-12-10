package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	y, x int
}

type WalkFunc func([]string, Point, rune, map[Point]int)

func getStartingPosition(mappedArea []string) Point {
	p := Point{0, 0}
	for i, row := range mappedArea {
		for j, cell := range row {
			if cell == '^' {
				p.y, p.x = i, j
				return p
			}
		}
	}

	return p
}

func (p Point) getPointValue(mappedArea []string) string {
	if p.x < 0 || p.x >= len(mappedArea[0]) || p.y < 0 || p.y >= len(mappedArea) {
		return " "
	}

	return string(mappedArea[p.y][p.x])
}

func (p Point) move(direction rune) Point {
	switch direction {
	case 'U':
		p.y--
	case 'R':
		p.x++
	case 'D':
		p.y++
	case 'L':
		p.x--
	}

	return p
}

func rotateDirection(currentDirection rune) rune {
	switch currentDirection {
	case 'U':
		return 'R'
	case 'R':
		return 'D'
	case 'D':
		return 'L'
	case 'L':
		return 'U'
	}

	return 'U'
}

func reverseDirection(currentDirection rune) rune {
	switch currentDirection {
	case 'U':
		return 'D'
	case 'R':
		return 'L'
	case 'D':
		return 'U'
	case 'L':
		return 'R'
	}

	return 'U'
}

func placeObstacle(mappedArea []string, position Point) []string {
	updatedMappedArea := make([]string, len(mappedArea))
	copy(updatedMappedArea, mappedArea)

	updatedMappedArea[position.y] = updatedMappedArea[position.y][:position.x] + "#" + updatedMappedArea[position.y][position.x+1:]

	return updatedMappedArea
}

func walk(mappedArea []string,
	currentPosition Point,
	currentDirection rune,
	visitedPositions map[Point]int,
	loopPositions map[Point]int,
	placedObstaclePosition Point,
	visitedPositionWithDirection map[Point]map[rune]int,
) WalkFunc {
	// Check out of bounds
	if currentPosition.x < 0 || currentPosition.x >= len(mappedArea[0]) ||
		currentPosition.y < 0 || currentPosition.y >= len(mappedArea) {
		return nil
	}

	// Check if current position is a wall
	if currentPosition.getPointValue(mappedArea) == "#" {
		currentPosition = currentPosition.move(reverseDirection(currentDirection))
		currentDirection = rotateDirection(currentDirection)

		return walk(mappedArea, currentPosition, currentDirection, visitedPositions, loopPositions, placedObstaclePosition, visitedPositionWithDirection)
	}

	// Check if current position is a space or starting position
	if currentPosition.getPointValue(mappedArea) == "." || currentPosition.getPointValue(mappedArea) == "^" {
		// Check if it's a loop
		if visitedPositions[currentPosition] > 1 &&
			visitedPositionWithDirection[currentPosition][currentDirection] > 1 {
			loopPositions[placedObstaclePosition]++
			return nil
		}

		if _, ok := visitedPositionWithDirection[currentPosition]; !ok {
			visitedPositionWithDirection[currentPosition] = make(map[rune]int)
		}
		visitedPositionWithDirection[currentPosition][currentDirection]++

		visitedPositions[currentPosition]++
		currentPosition = currentPosition.move(currentDirection)

		return walk(mappedArea, currentPosition, currentDirection, visitedPositions, loopPositions, placedObstaclePosition, visitedPositionWithDirection)
	}

	return walk(mappedArea, currentPosition, currentDirection, visitedPositions, loopPositions, placedObstaclePosition, visitedPositionWithDirection)
}

func part1(mappedArea []string) (int, map[Point]int) {
	visitedPositions := make(map[Point]int)

	walk(mappedArea,
		getStartingPosition(mappedArea),
		'U',
		visitedPositions,
		map[Point]int{},
		Point{},
		map[Point]map[rune]int{},
	)

	result := 0
	for range visitedPositions {
		result++
	}

	return result, visitedPositions
}

func part2(mappedArea []string, visitedPositions map[Point]int) int {
	loopPositions := make(map[Point]int)

	for k, _ := range visitedPositions {
		updatedMappedArea := placeObstacle(mappedArea, k)
		walk(updatedMappedArea,
			getStartingPosition(mappedArea),
			'U',
			map[Point]int{},
			loopPositions,
			k,
			map[Point]map[rune]int{},
		)
	}

	result := 0
	for range loopPositions {
		result++
	}

	return result
}

func main() {
	// content, err := os.ReadFile("input_test.txt")
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]

	part1Result, visitedPositions := part1(lines)
	fmt.Println("\npart1: ", part1Result)
	fmt.Println("\npart2: ", part2(lines, visitedPositions))
}
