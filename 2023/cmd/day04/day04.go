package main

import (
	"aoc/internal/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	result := Solve1("2023/cmd/day04/1.data")
	fmt.Printf("day 04 part 1 %d \n", result)
	result = Solve2("2023/cmd/day04/1.data")
	fmt.Printf("day 04 part 2 %d \n", result)
}

func Solve1(file string) int {
	lines := utils.ReadFile((file))
	sum := 0
	for _, line := range lines {
		line := strings.SplitAfter(line, ":")[1]
		cards := strings.SplitAfter(line, "|")
		winners := map[string]bool{}
		r := regexp.MustCompile(`\d+`)
		winningNumbers := r.FindAllString(cards[0], -1)
		for _, num := range winningNumbers {
			winners[num] = true
		}
		gameNumbers := r.FindAllString(cards[1], -1)
		points := 0
		for _, num := range gameNumbers {
			if winners[num] {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		sum += points
	}
	return sum
}

func Solve2(file string) int {
	lines := utils.ReadFile((file))
	scratchCards := map[int]int{}
	for i := range lines {
		scratchCards[i] = 1
	}

	for cardNo, line := range lines {
		line := strings.SplitAfter(line, ":")[1]
		cards := strings.SplitAfter(line, "|")
		winners := map[string]bool{}
		r := regexp.MustCompile(`\d+`)
		winningNumbers := r.FindAllString(cards[0], -1)
		for _, num := range winningNumbers {
			winners[num] = true
		}
		gameNumbers := r.FindAllString(cards[1], -1)
		numberOfWinners := 0
		for _, num := range gameNumbers {
			if winners[num] {
				numberOfWinners++
			}
		}
		numberOfWinningCards := scratchCards[cardNo]
		for j := 0; j < numberOfWinners; j++ {
			scratchCards[cardNo+j+1] += 1 * numberOfWinningCards
		}
	}
	sum := 0
	for _, v := range scratchCards {
		sum += v
	}
	return sum
}
