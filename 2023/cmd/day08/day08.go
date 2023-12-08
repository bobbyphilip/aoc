package main

import (
	"aoc/internal/utils"
	"fmt"
	"regexp"
)

func main() {
	result := Solve1("2023/cmd/day08/1.data")
	fmt.Printf("day 08 part 1 %d \n", result)
	result = Solve2("2023/cmd/day08/1.data")
	fmt.Printf("day 08 part 2 %d \n", result)
}

func Solve1(file string) int {
	count := 0
	lines := utils.ReadFile((file))
	directions := []rune(lines[0])
	steps := make([]int, len(directions))
	for i, dir := range directions {
		if dir == 'L' {
			steps[i] = 0
		} else {
			steps[i] = 1
		}
	}
	data := make(map[string][2]string, len(lines)-2)
	re := regexp.MustCompile(`[a-zA-Z]+`)
	for _, line := range lines[2:] {
		strs := re.FindAllString(line, -1)
		data[strs[0]] = [2]string{strs[1], strs[2]}
	}
	curr := "AAA"
	for curr != "ZZZ" {
		index := count % len(steps)
		curr = data[curr][steps[index]]
		count++
	}
	return count
}

func Solve2(file string) int {
	lines := utils.ReadFile((file))
	directions := []rune(lines[0])
	steps := make([]int, len(directions))
	for i, dir := range directions {
		if dir == 'L' {
			steps[i] = 0
		} else {
			steps[i] = 1
		}
	}

	aNodes := []string{}
	data := make(map[string][2]string, len(lines)-2)
	re := regexp.MustCompile(`[a-zA-Z]+`)
	for _, line := range lines[2:] {
		strs := re.FindAllString(line, -1)
		data[strs[0]] = [2]string{strs[1], strs[2]}
		if strs[0][2] == 'A' {
			aNodes = append(aNodes, strs[0])
		}
	}
	res := make([]int, len(aNodes))
	for i, curr := range aNodes {
		count := 0
		for curr[2] != 'Z' {
			index := count % len(steps)
			curr = data[curr][steps[index]]
			count++
		}
		res[i] = count
	}
	a := res[0]
	for _, b := range res[1:] {
		t := lcm(a, b)
		a = t
	}
	return a
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)

}
