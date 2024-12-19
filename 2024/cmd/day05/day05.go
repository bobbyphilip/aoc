package main

import (
	"aoc/internal/utils"
	"fmt"
	"slices"
)

func main() {
	result := Solve1("2024/cmd/day05/1.data")
	fmt.Printf("day 05 part 1 %d \n", result)
	result = Solve2("2024/cmd/day05/1.data")
	fmt.Printf("day 05 part 2 %d \n", result)
}

func Solve1(file string) int {
	sum := 0
	lines := utils.ReadFile((file))
	order := map[int][]int{}
	inputs := [][]int{}
	orderering := true
	for _, line := range lines {
		if line == "" {
			orderering = false
			continue
		}
		nums := utils.ReadNumbersFromLine(line)
		if orderering {
			order[nums[0]] = append(order[nums[0]], nums[1])
		} else {
			inputs = append(inputs, nums)
		}
	}
	for _, input := range inputs {
		valid := isValid(input, order)
		if valid {
			sum += input[len(input)/2]
		}
	}
	return sum
}

func isValid(input []int, order map[int][]int) bool {
	valid := true
	for i := len(input) - 1; i >= 0; i-- {
		page := input[i]
		rules := order[page]
		for j := 0; j < i; j++ {
			if slices.Contains(rules, input[j]) {
				valid = false
				break
			}
		}
		if !valid {
			break
		}
	}
	return valid
}

func Solve2(file string) int {
	sum := 0
	lines := utils.ReadFile((file))
	order := map[int][]int{}
	inputs := [][]int{}
	orderering := true
	for _, line := range lines {
		if line == "" {
			orderering = false
			continue
		}
		nums := utils.ReadNumbersFromLine(line)
		if orderering {
			order[nums[0]] = append(order[nums[0]], nums[1])
		} else {
			inputs = append(inputs, nums)
		}
	}
	for _, input := range inputs {
		valid := isValid(input, order)
		if valid {
			continue
		}
		slices.SortFunc(input, func(a, b int) int {
			rules := order[a]
			if slices.Contains(rules, b) {
				return -1
			}
			return 0
		})
		sum += input[len(input)/2]

	}
	return sum
}
