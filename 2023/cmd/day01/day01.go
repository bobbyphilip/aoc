package main

import (
	"aoc/internal/utils"
	"fmt"
	"strings"
)

func main() {
	result := Solve1("2023/data/day01/1.txt")
	fmt.Printf("day 01 %d \n", result)
	result = Solve2("2023/data/day01/1.txt")
	fmt.Printf("day 01 part 2 %d \n", result)
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

func Solve2(file string) int {
	lines := utils.ReadFile((file))
	var sum int
	for _, line := range lines {
		sum += findNumber2(line)

	}
	return sum
}

var numberReference = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func findNumber2(text string) int {
	firstIndex := len(text)
	lastIndex := -1
	var first, last int

	for k, v := range numberReference {
		index := strings.Index(text, k)
		if index == -1 {
			continue
		}
		if index < firstIndex {
			firstIndex = index
			first = v
		}
		index = strings.LastIndex(text, k)
		if index > lastIndex {
			lastIndex = index
			last = v
		}
	}
	result := first*10 + last
	fmt.Printf(" line sum %d \n", result)
	return result

}
