package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(f *os.File) (totalPoints int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Split(strings.Split(line, ":")[1], "|")
		winningNums := strings.Fields(numbers[0])
		numbersOnCard := strings.Fields(numbers[1])

		points := 0
		for _, num := range winningNums {
			if slices.Contains(numbersOnCard, num) {
				if points == 0 {
					points++
					continue
				}
				points = points * 2
			}
		}

		totalPoints += points
	}

	return
}

func part2(f *os.File) (totalCards int) {
	cards := make(map[int]int)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		partition := strings.Split(line, ":")
		gameNumber, _ := strconv.Atoi(strings.Fields(partition[0])[1])

		numbers := strings.Split(partition[1], "|")
		winningNums := strings.Fields(numbers[0])
		numbersOnCard := strings.Fields(numbers[1])

		totalMatchingNumbers := 0
		for _, num := range winningNums {
			if slices.Contains(numbersOnCard, num) {
				totalMatchingNumbers++
			}
		}

		// Atleast one copy of the card is guaranteed
		cards[gameNumber]++

		// Play the game for the number of cards
		for i := 0; i < cards[gameNumber]; i++ {
			updateCards(cards, gameNumber, totalMatchingNumbers)
		}
	}

	for _, v := range cards {
		totalCards += v
	}

	return
}

func updateCards(cards map[int]int, gameNumber, totalMatchingNumbers int) {
	for i := gameNumber + 1; i <= totalMatchingNumbers+gameNumber; i++ {
		if i > 214 {
			break
		}
		if _, ok := cards[i]; !ok {
			cards[i] = 1
			continue
		}
		cards[i]++
	}
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	// fmt.Println(part1(file))
	fmt.Println(part2(file))
}
