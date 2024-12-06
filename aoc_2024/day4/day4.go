package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func (p Point) getValue(arr []string) string {
	return string(arr[p.y][p.x])
}

func (p Point) verticalCheck(lines []string) bool {
	XMAXappears := false
	var s string

	if p.y+3 < len(lines) {
		for i := 0; i < 4; i++ {
			point := Point{p.x, p.y + i}
			s += point.getValue(lines)
		}
	}

	if s == "XMAS" || s == "SAMX" {
		XMAXappears = true
	}

	return XMAXappears
}

func (p Point) diagonalCheck(lines []string) int {
	XMAXappearances := 0
	var downLeftStr, downRightStr string

	if p.y+3 < len(lines) && p.x-3 >= 0 {
		for i := 0; i < 4; i++ {
			point := Point{p.x - i, p.y + i}
			downLeftStr += point.getValue(lines)
		}
	}

	if p.y+3 < len(lines) && p.x+3 < len(lines[0]) {
		for i := 0; i < 4; i++ {
			point := Point{p.x + i, p.y + i}
			downRightStr += point.getValue(lines)
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

func (p Point) checkX(lines []string) bool {
	isX := false
	var xmasRight, xmasLeft string

	if p.y+2 < len(lines) && p.x+2 < len(lines[0]) {
		for i := 0; i <= 2; i++ {
			point := Point{p.x + i, p.y + i}
			xmasRight += point.getValue(lines)
		}

		p.x += 2
		for i := 0; i <= 2; i++ {
			point := Point{p.x - i, p.y + i}
			xmasLeft += point.getValue(lines)
		}
	}

	if (xmasLeft == "MAS" || xmasLeft == "SAM") && (xmasRight == "MAS" || xmasRight == "SAM") {
		isX = true
	}

	return isX
}

/* Solution */

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

			if point.getValue(lines) == "X" || point.getValue(lines) == "S" {
				if point.verticalCheck(lines) {
					wordCount += 1
				}
				wordCount += point.diagonalCheck(lines)
			}
		}
	}

	return wordCount
}

func part2(lines []string) int {
	wordCount := 0
	for y, line := range lines {
		for x := range line {
			point := Point{x, y}

			if point.x < 0 || point.y < 0 {
				continue
			}

			if point.x >= len(lines) || point.y >= len(line) {
				continue
			}

			if point.getValue(lines) == "M" || point.getValue(lines) == "S" {
				if point.checkX(lines) {
					wordCount += 1
				}
			}
		}
	}

	return wordCount
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
	fmt.Println("\npart2: ", part2(lines))
}
