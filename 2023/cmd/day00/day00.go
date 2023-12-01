package main

import (
	"aoc/internal/utils"
	"fmt"
)

func main() {
	result := Solve("2023/data/day00/0.txt")
	fmt.Printf("day 0  %d \n", result)
}

func Solve(file string) int {
	return len(utils.ReadFile(file))
}
