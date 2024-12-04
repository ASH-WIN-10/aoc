package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Sphagetti
func part1Alternate(lines []string) int {
	result := 0
	for _, line := range lines {
		x := strings.Split(line, "mul(")
		for _, s := range x {
			str := strings.Split(s, ")")[0]

			// Extract only this pattern from str
			// If its length is less than than its original length then continue
			match := regexp.MustCompile(`\d+\,\d+`).FindString(str)
			if len(match) < len(str) || match == "" {
				continue
			}

			nums := strings.Split(match, ",")
			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])

			result += num1 * num2
		}
	}

	return result
}

func part1(lines []string) int {
	result := 0
	for _, line := range lines {
		pattern := `mul\((\d+),(\d+)\)`
		re := regexp.MustCompile(pattern)

		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) == 3 {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				result += x * y
			}
		}
	}

	return result
}

func part2(lines []string) int {
	result := 0
	do := true // enable mul at the beginning
	for _, line := range lines {
		pattern := `mul\((\d+),(\d+)\)|do\(\)|don't\(\)`
		r := regexp.MustCompile(pattern)

		matches := r.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if match[0] == "do()" {
				do = true
			} else if match[0] == "don't()" {
				do = false
			}

			if do {
				if strings.Contains(match[0], "mul") {
					num1, _ := strconv.Atoi(match[1])
					num2, _ := strconv.Atoi(match[2])
					result += num1 * num2
				}
			}
		}
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

	fmt.Println("\npart1: ", part1(lines))
	fmt.Println("\npart2: ", part2(lines))
}
