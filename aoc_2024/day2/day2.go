package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isSafe(report []int) bool {
	isInc, isDec, okDiff := true, true, true
	for i := 0; i < len(report)-1; i++ {
		if report[i] >= report[i+1] {
			isInc = false
		}

		if report[i] <= report[i+1] {
			isDec = false
		}

		diff := abs(report[i] - report[i+1])
		if diff < 1 || diff > 3 {
			okDiff = false
			break
		}
	}

	if (isInc || isDec) && okDiff {
		return true
	}

	return false
}

func part1(lines []string) int {
	safeReports := 0

	for _, line := range lines {
		reportS := strings.Split(line, " ")
		report := make([]int, len(reportS))
		for i, v := range reportS {
			report[i], _ = strconv.Atoi(v)
		}

		if isSafe(report) {
			safeReports++
		}
	}

	return safeReports
}

func part2(lines []string) int {
	safeReports := 0

	for _, line := range lines {
		reportS := strings.Split(line, " ")
		report := make([]int, len(reportS))
		for i, v := range reportS {
			report[i], _ = strconv.Atoi(v)
		}

		if isSafe(report) {
			safeReports++
			continue
		}

		for i := 0; i < len(report); i++ {
			tempReport := make([]int, 0, len(report)-1)
			tempReport = append(tempReport, report[:i]...)
			tempReport = append(tempReport, report[i+1:]...)

			if isSafe(tempReport) {
				safeReports++
				break
			}
		}
	}

	return safeReports
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]

	fmt.Println("part1: ", part1(lines))
	fmt.Println("part2: ", part2(lines))
}
