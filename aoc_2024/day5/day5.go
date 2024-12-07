package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(rules map[string][]string, pageNums [][]string) (int, [][]string) {
	sumOfMiddlePageNums := 0
	incorrectPageNums := make([][]string, 0)
	for i := 0; i < len(pageNums); i++ {
		good := true
		for j := 0; j < len(pageNums[i])-1; j++ {
			num1 := pageNums[i][j]
			num2 := pageNums[i][j+1]

			if _, ok := rules[num1]; !ok {
				if _, ok := rules[num2]; !ok {
					good = true
					continue
				}
				good = false
				incorrectPageNums = append(incorrectPageNums, pageNums[i])
				break
			} else {
				if slices.Contains(rules[num1], num2) {
					good = true
					continue
				}
				incorrectPageNums = append(incorrectPageNums, pageNums[i])
				good = false
				break
			}
		}

		if good {
			midVal, _ := strconv.Atoi(pageNums[i][len(pageNums[i])/2])
			sumOfMiddlePageNums += midVal
		}
	}

	return sumOfMiddlePageNums, incorrectPageNums
}

func part2(rules map[string][]string, incorrectPageNums [][]string) int {
	return 0
}

func main() {
	// content, err := os.ReadFile("input_test.txt")
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	sections := strings.Split(string(content), "\n\n")

	section1 := strings.Split(string(sections[0]), "\n")
	section2 := strings.Split(string(sections[1]), "\n")
	section2 = section2[:len(section2)-1]

	rules := make(map[string][]string)
	pageNums := make([][]string, 0)

	for _, line := range section1 {
		nums := strings.Split(line, "|")
		if _, ok := rules[nums[0]]; ok {
			rules[nums[0]] = append(rules[nums[0]], nums[1])
		} else {
			rules[nums[0]] = []string{nums[1]}
		}
	}

	for _, line := range section2 {
		pageNums = append(pageNums, strings.Split(line, ","))
	}

	part1Result, incorrectPageNums := part1(rules, pageNums)
	fmt.Println("\npart1: ", part1Result)
	fmt.Println("\npart2: ", part2(rules, incorrectPageNums))
}
