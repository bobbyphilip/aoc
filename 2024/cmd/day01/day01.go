package main

import (
	"aoc/internal/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
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
		l := strings.Fields(line)

		lInt, _ := strconv.Atoi(l[0])
		left = append(left, lInt)

		rInt, _ := strconv.Atoi(l[1])
		right = append(right, rInt)

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
		l := strings.Fields(line)

		lInt, _ := strconv.Atoi(l[0])
		left = append(left, lInt)

		rInt, _ := strconv.Atoi(l[1])
		right = append(right, rInt)

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
