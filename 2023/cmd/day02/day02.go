package main

import (
	"aoc/internal/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	result := Solve1("2023/cmd/day02/1.data")
	fmt.Printf("day 02 part 1 %d \n", result)
	result = Solve2("2023/cmd/day02/1.data")
	fmt.Printf("day 02 part 2 %d \n", result)
}

func Solve1(file string) int {
	lines := utils.ReadFile((file))
	sum := 0
	totals := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	r := regexp.MustCompile(`\d+\sblue|\d+\sgreen|\d+\sred`)
	for index, line := range lines {
		invalid := false
		drawnValues := r.FindAllString(line, -1)
		for _, ball := range drawnValues {
			countAndColour := strings.Split(ball, " ")
			count, err := strconv.Atoi(countAndColour[0])
			if err != nil {
				fmt.Println("error")
				continue
			}
			if count > totals[countAndColour[1]] {
				invalid = true
				break
			}
		}
		if !invalid {
			sum += (index + 1)
		}
	}
	return sum
}

func Solve2(file string) int {
	lines := utils.ReadFile((file))
	sum := 0
	r := regexp.MustCompile(`\d+\sblue|\d+\sgreen|\d+\sred`)
	for _, line := range lines {
		maxBalls := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		drawnValues := r.FindAllString(line, -1)
		for _, ball := range drawnValues {
			countAndColour := strings.Split(ball, " ")
			count, err := strconv.Atoi(countAndColour[0])
			if err != nil {
				fmt.Println("error")
				continue
			}
			if count > maxBalls[countAndColour[1]] {
				maxBalls[countAndColour[1]] = count
			}
		}
		power := 1
		for _, count := range maxBalls {
			power *= count
		}
		sum += power
	}
	return sum
}
