package main

import (
	"aoc/internal/utils"
	"fmt"
	"slices"
)

func main() {
	result := Solve1("2024/cmd/day02/1.data")
	fmt.Printf("day 01 part 1 %d \n", result)
	result = Solve2("2024/cmd/day02/1.data")
	fmt.Printf("day 01 part 2 %d \n", result)
}

func Solve1(file string) int {
	lines := utils.ReadFile((file))
	count := 0
	for _, line := range lines {
		numbers := utils.ReadNumbersFromLine(line)
		if isReportSafe(numbers) {
			count++
		}
	}
	return count
}

func Solve2(file string) int {
	lines := utils.ReadFile((file))
	count := 0
	for _, line := range lines {
		numbers := utils.ReadNumbersFromLine(line)
		if isReportSafe(numbers) {
			count++
		} else if isDampenedReportSafe(numbers) {
			count++
		}
	}
	return count
}

func isReportSafe(input []int) bool {
	dec := false
	inc := false
	for i := 1; i < len(input); i++ {
		diff := input[i] - input[i-1]
		if diff == 0 {
			return false
		}
		if diff < 0 {
			dec = true
			diff *= -1
		} else {
			inc = true
		}
		if dec && inc {
			return false
		}
		if diff > 3 {
			return false
		}

	}
	return true
}
func isDampenedReportSafe(input []int) bool {
	length := len(input)
	for i := 0; i < length; i++ {
		reduced := slices.Concat(input[0:i], input[i+1:length])
		if isReportSafe(reduced) {
			return true
		}
	}
	return false
}
