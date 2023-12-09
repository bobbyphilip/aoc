package main

import (
	"aoc/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	result := Solve1("2023/cmd/day09/1.data")
	fmt.Printf("day 09 part 1 %d \n", result)
	result = Solve2("2023/cmd/day09/1.data")
	fmt.Printf("day 09 part 2 %d \n", result)
}

func Solve1(file string) int {
	sum := 0
	lines := utils.ReadFile((file))
	for _, line := range lines {
		nums := []int{}
		for _, str := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(str)
			nums = append(nums, num)
		}
		sum += nextNumber(nums, true)
	}
	return sum
}

func nextNumber(nums []int, original bool) int {
	allZeros := true
	for _, num := range nums {
		if num != 0 {
			allZeros = false
			break
		}
	}
	if allZeros {
		return 0
	} else {
		newNums := []int{}
		for i := 0; i < len(nums)-1; i++ {
			newNums = append(newNums, nums[i+1]-nums[i])
		}
		if original {
			return nums[len(nums)-1] + nextNumber(newNums, original)
		} else {
			return nums[0] - nextNumber(newNums, original)
		}
	}
}

func Solve2(file string) int {
	sum := 0
	lines := utils.ReadFile((file))
	for _, line := range lines {
		nums := []int{}
		for _, str := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(str)
			nums = append(nums, num)
		}
		sum += nextNumber(nums, false)
	}
	return sum
}
