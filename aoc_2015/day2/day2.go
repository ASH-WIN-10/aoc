package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1and2(file *os.File) (totalSurfaceArea int, totalRibonSize int) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		dimensions := strings.Split(line, "x")
		l, _ := strconv.Atoi(dimensions[0])
		w, _ := strconv.Atoi(dimensions[1])
		h, _ := strconv.Atoi(dimensions[2])

		sortedDimensions := []int{l, w, h}
		slices.Sort(sortedDimensions)

		totalSurfaceArea += (2*l*w + 2*w*h + 2*h*l) + (sortedDimensions[0] * sortedDimensions[1])
		totalRibonSize += 2*(sortedDimensions[0]+sortedDimensions[1]) + l*w*h
	}
	return
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	// Part 1 and 2
	fmt.Println(part1and2(file))
}
