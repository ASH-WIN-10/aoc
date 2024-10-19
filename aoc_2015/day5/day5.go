package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func hasConsecutiveLetters(line string) bool {
	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] {
			return true
		}
	}
	return false
}

func part1(content string) (numOfNiceStrings int) {
	splitContent := strings.Split(content, "\n")
	for _, line := range splitContent[:len(splitContent)-1] {
		r, _ := regexp.Compile("[aeiou]")
		vowelsInString := r.FindAllString(line, len(line))
		if len(vowelsInString) < 3 {
			continue
		}

		if !hasConsecutiveLetters(line) {
			continue
		}

		if ok, _ := regexp.MatchString("(ab|cd|pq|xy)", line); !ok {
			numOfNiceStrings++
		}
	}
	return
}

func hasOneBetweenTwo(line string) bool {
	for i := 0; i < len(line)-2; i++ {
		if line[i] == line[i+2] {
			return true
		}
	}
	return false
}

func hasTwoConsecutivePair(line string) bool {
	for i := 0; i < len(line)-2; i++ {
		if strings.Count(line, line[i:i+2]) >= 2 {
			return true
		}
	}
	return false
}

func part2(content string) (numOfNiceStrings int) {
	splitContent := strings.Split(content, "\n")
	for _, line := range splitContent[:len(splitContent)-1] {
		if !hasOneBetweenTwo(line) {
			continue
		}

		if hasTwoConsecutivePair(line) {
			numOfNiceStrings++
		}
	}

	return
}

func main() {
	content, _ := os.ReadFile("input.txt")
	// fmt.Println("Nice strings: ", part1(string(content)))
	fmt.Println("Nice strings: ", part2(string(content)))
}
