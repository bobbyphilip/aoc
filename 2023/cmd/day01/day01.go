package main

import (
	"aoc/internal/utils"
	"fmt"
)

func main() {
	result := Solve1("2023/data/day01/1.txt")
	fmt.Printf("day 01 %d \n", result)
}

func Solve1(file string) int {
	lines := utils.ReadFile((file))
	var sum int
	for _, line := range lines {
		sum += findNumber(line)

	}
	return sum
}

func findNumber(text string) int {
	result := 0

	runes := []rune(text)
	for _, letter := range runes {
		if letter >= '0' && letter <= '9' {
			result = 10 * int(letter-'0')
			break
		}
	}
	for i := len(runes) - 1; i >= 0; i-- {
		letter := runes[i]
		if letter >= '0' && letter <= '9' {
			result += int(letter - '0')
			break
		}
	}
	fmt.Printf(" line sum %d \n", result)
	return result
}
