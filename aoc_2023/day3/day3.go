package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var dirs = [][]int{
	{-1, 0}, {1, 0},
	{0, -1}, {0, 1},
	{-1, -1}, {-1, 1},
	{1, -1}, {1, 1},
}

type Point struct {
	x, y int
}

func part1(content string) (sumOfAllPartNum int) {
	lines := strings.Split(content, "\n")
	lines = lines[:len(lines)-1] // Remove the last empty line

	for lineIdx, line := range lines {
		r := regexp.MustCompile("[0-9]+")
		numIdxs := r.FindAllStringIndex(line, len(line))

	numLoop:
		for _, numIdx := range numIdxs {
			for i := numIdx[0]; i < numIdx[1]; i++ {
				for _, dir := range dirs {
					point := Point{x: i + dir[0], y: lineIdx + dir[1]}

					// Check out of bounds
					if point.x < 0 || point.y < 0 ||
						point.x >= len(line) || point.y >= len(lines) {
						continue
					}

					// Check if currPoint is a char and sum the number if true
					if currPoint := lines[point.y][point.x]; !strings.ContainsAny(string(currPoint), ".0123456789") {
						num, _ := strconv.Atoi(line[numIdx[0]:numIdx[1]])
						sumOfAllPartNum += num
						continue numLoop // Go to the next num
					}
				}
			}
		}
	}

	return
}

type PointWithChar struct {
	num       int
	isPartNum bool
}

func part2(content string) (sumOfAllPartNum int) {
	lines := strings.Split(content, "\n")
	lines = lines[:len(lines)-1] // Remove the last empty line

	points := make(map[Point]PointWithChar)

	for lineIdx, line := range lines {
		r := regexp.MustCompile("[0-9]+")
		numIdxs := r.FindAllStringIndex(line, len(line))

	numLoop:
		for _, numIdx := range numIdxs {
			for i := numIdx[0]; i < numIdx[1]; i++ {
				for _, dir := range dirs {
					point := Point{x: i + dir[0], y: lineIdx + dir[1]}

					// Check out of bounds
					if point.x < 0 || point.y < 0 ||
						point.x >= len(line) || point.y >= len(lines) {
						continue
					}

					// Check if currPoint is a char and sum the number if true
					if currPoint := lines[point.y][point.x]; strings.ContainsAny(string(currPoint), "*") {
						num, _ := strconv.Atoi(line[numIdx[0]:numIdx[1]])
						if _, ok := points[point]; ok {
							points[point] = PointWithChar{num: points[point].num * num, isPartNum: true}
						} else {
							points[point] = PointWithChar{num: num, isPartNum: false}
						}
						continue numLoop // Go to the next num
					}
				}
			}
		}
	}

	for _, point := range points {
		if point.isPartNum {
			sumOfAllPartNum += point.num
		}
	}

	return
}

func main() {
	content, _ := os.ReadFile("input.txt")
	// fmt.Println(part1(string(content)))
	fmt.Println(part2(string(content)))
}
