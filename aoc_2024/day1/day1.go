package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(leftList, rightList []int) int {
	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		totalDistance += abs(leftList[i] - rightList[i])
	}

	return totalDistance
}

func part2(leftList, rightList []int) int {
	mp := make(map[int]int)
	similarityScore := 0

	for _, rightNum := range rightList {
		mp[rightNum]++
	}

	for _, leftNum := range leftList {
		similarityScore += leftNum * mp[leftNum]
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

	fmt.Println("part1: ", part1(leftList, rightList))
	fmt.Println("part2: ", part2(leftList, rightList))
}
