package main

import (
	"aoc/internal/utils"
	"fmt"
)

func main() {
	result := Solve1("2023/cmd/day00/0.data")
	fmt.Printf("day 00 part 1 %d \n", result)
	result = Solve2("2023/cmd/day00/1.data")
	fmt.Printf("day 00 part 2 %d \n", result)
}

func Solve1(file string) int {
	sum := 0
	lines := utils.ReadFile((file))
	sum = len(lines)
	return sum
}

func Solve2(file string) int {
	sum := 0
	lines := utils.ReadFile((file))
	sum = len(lines)
	return sum
}
