package main

import (
	"fmt"
	"os"
)

func part1() {
	content, _ := os.ReadFile("input.txt")
	floor := 0
	for _, v := range string(content) {
		if string(v) == "(" {
			floor++
		}
		if string(v) == ")" {
			floor--
		}
	}
	fmt.Printf("floor: %v\n", floor)
}

func part2() {
	content, _ := os.ReadFile("input.txt")

	floor := 0
	enterBasementFirstTime := 0
	for i, v := range string(content) {
		if floor == -1 {
			enterBasementFirstTime = i
			break
		}
		if string(v) == "(" {
			floor++
		}
		if string(v) == ")" {
			floor--
		}
	}
	fmt.Printf("enterBasementFirstTime: %v\n", enterBasementFirstTime)
}

func main() {
	// part1()
	part2()
}
