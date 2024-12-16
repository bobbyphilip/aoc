package main

import (
	"aoc/internal/utils"
	"fmt"
	"slices"
)

func main() {
	result := Solve1("2024/cmd/day01/1.data")
	fmt.Printf("day 01 part 1 %d \n", result)
	result = Solve2("2024/cmd/day01/1.data")
	fmt.Printf("day 01 part 2 %d \n", result)
}

func Solve1(file string) int {
	lines := utils.ReadFile((file))
	left := []int{}
	right := []int{}
	for _, line := range lines {
		numbers := utils.ReadNumbersFromLine(line)
		left = append(left, numbers[0])
		right = append(right, numbers[1])
	}
	slices.Sort(left)
	slices.Sort(right)
	sum := 0
	for i := range left {
		diff := left[i] - right[i]
		if diff < 0 {
			diff *= -1
		}
		sum += diff
	}
	return sum
}

func Solve2(file string) int {
	lines := utils.ReadFile((file))
	left := []int{}
	right := []int{}
	for _, line := range lines {
		numbers := utils.ReadNumbersFromLine(line)
		left = append(left, numbers[0])
		right = append(right, numbers[1])
	}
	rightCounts := make(map[int]int)
	for _, val := range right {
		count := rightCounts[val]
		rightCounts[val] = count + 1
	}
	sum := 0
	for _, val := range left {
		count := rightCounts[val]
		sum += (count * val)
	}
	return sum
}
