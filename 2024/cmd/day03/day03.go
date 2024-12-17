package main

import (
	"aoc/internal/utils"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	result := Solve1("2024/cmd/day03/1.data")
	fmt.Printf("day 03 part 1 %d \n", result)
	result = Solve2("2024/cmd/day03/1.data")
	fmt.Printf("day 03 part 2 %d \n", result)
}

func Solve1(file string) int {
	sum := 0
	r, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	n, _ := regexp.Compile(`\d+`)
	lines := utils.ReadFile((file))
	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		for _, match := range matches {
			nums := n.FindAllString(match, -1)
			l, _ := strconv.Atoi(nums[0])
			r, _ := strconv.Atoi(nums[1])
			sum += (l * r)
		}
	}
	return sum
}

func Solve2(file string) int {
	sum := 0
	r := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	n := regexp.MustCompile(`\d+`)
	lines := utils.ReadFile((file))
	do := true
	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		for _, match := range matches {
			if match == "do()" {
				do = true
				continue
			}
			if match == "don't()" {
				do = false
				continue
			}
			if !do {
				continue
			}
			nums := n.FindAllString(match, -1)
			l, _ := strconv.Atoi(nums[0])
			r, _ := strconv.Atoi(nums[1])
			sum += (l * r)
		}
	}
	return sum
}
