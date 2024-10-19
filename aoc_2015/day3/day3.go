package main

import (
	"fmt"
	"os"
)

func part1(content string) int {
	houses := make(map[string]int)
	houses["(0, 0)"] = 0
	var x, y int
	for _, v := range content {
		if v == '^' {
			x++
		}
		if v == 'v' {
			x--
		}
		if v == '>' {
			y++
		}
		if v == '<' {
			y--
		}
		houses[fmt.Sprintf("(%d, %d)", x, y)] = 0
	}
	return len(houses)
}

func checker(x, y *int, v rune) {
	if v == '^' {
		*x++
	}
	if v == 'v' {
		*x--
	}
	if v == '>' {
		*y++
	}
	if v == '<' {
		*y--
	}
}

func part2(content string) int {
	houses := make(map[string]int)
	houses["(0, 0)"] = 0

	var x1, y1 int
	var x2, y2 int
	for i, v := range content {
		if (i+1)%2 == 0 {
			checker(&x1, &y1, v)
		}
		if (i+1)%2 != 0 {
			checker(&x2, &y2, v)
		}
		houses[fmt.Sprintf("(%d, %d)", x1, y1)] = 0
		houses[fmt.Sprintf("(%d, %d)", x2, y2)] = 0
	}

	return len(houses)
}

func main() {
	content, _ := os.ReadFile("input.txt")

	// fmt.Printf("Number of houses: %v\n", part1(string(content)))
	fmt.Println(part2(string(content)))
}
