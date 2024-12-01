package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(leftList, rightList []int) int {
	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		totalDistance += int(math.Abs(float64(leftList[i] - rightList[i])))
	}

	return totalDistance
}

func part2(leftList, rightList []int) int {
	similarityScore := 0
	for _, leftNum := range leftList {
		appears := 0
		for _, rightNum := range rightList {
			if leftNum == rightNum {
				appears++
			}
		}

		similarityScore += leftNum * appears
	}

	return similarityScore
}

func main() {
	var leftList, rightList []int

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)

		leftNum, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println(err)
		}

		rightNum, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Println(err)
		}

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	// fmt.Println(part1(leftList, rightList))
	fmt.Println(part2(leftList, rightList))
}
