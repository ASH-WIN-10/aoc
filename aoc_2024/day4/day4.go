package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func (p Point) getPoint(arr []string) string {
	return string(arr[p.y][p.x])
}

func (p Point) verticalCheck(lines []string) bool {
	XMAXappears := false
	var upStr string

	if p.y+3 < len(lines) {
		for i := 0; i < 4; i++ {
			point := Point{p.x, p.y + i}
			upStr += point.getPoint(lines)
		}
	}

	if upStr == "XMAS" || upStr == "SAMX" {
		XMAXappears = true
	}

	return XMAXappears
}

func (p Point) diagonalCheck(lines []string) int {
	XMAXappearances := 0
	var downLeftStr, downRightStr string

	if p.y-3 >= 0 && p.x-3 >= 0 {
		for i := 0; i < 4; i++ {
			point := Point{p.x - i, p.y - i}
			downLeftStr += point.getPoint(lines)
		}
	}

	if p.y-3 >= 0 && p.x+3 < len(lines[0]) {
		for i := 0; i < 4; i++ {
			point := Point{p.x + i, p.y - i}
			downRightStr += point.getPoint(lines)
		}
	}

	if downLeftStr == "XMAS" || downLeftStr == "SAMX" {
		XMAXappearances += 1
	}

	if downRightStr == "XMAS" || downRightStr == "SAMX" {
		XMAXappearances += 1
	}

	return XMAXappearances
}

func part1(lines []string) int {
	wordCount := 0
	for y, line := range lines {
		wordCount += strings.Count(line, "XMAS")
		wordCount += strings.Count(line, "SAMX")

		for x := range line {
			point := Point{x, y}

			if point.x < 0 || point.y < 0 {
				continue
			}

			if point.x >= len(lines) || point.y >= len(line) {
				continue
			}

			if point.verticalCheck(lines) {
				wordCount += 1
			}
			wordCount += point.diagonalCheck(lines)
		}
	}

	return wordCount
}

func part2(lines []string) int {
	for _, line := range lines {
		fmt.Println(line)
	}

	return 0
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

	fmt.Println("\npart1: ", part1(lines))
	// fmt.Println("\npart2: ", part2(lines))
}
