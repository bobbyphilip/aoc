package main

import (
	"aoc/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	result := Solve1("2023/cmd/day06/0.data")
	fmt.Printf("day 00 part 1 %d \n", result)
	result = Solve2("2023/cmd/day06/1.data")
	fmt.Printf("day 00 part 2 %d \n", result)
}

func Solve1(file string) int {
	prod := 1
	lines := utils.ReadFile((file))

	lines[0] = strings.Split(lines[0], ":")[1]
	lines[1] = strings.Split(lines[1], ":")[1]
	times := strings.Fields(lines[0])
	distances := strings.Fields(lines[1])
	for i, timeString := range times {
		time, _ := strconv.Atoi(timeString)
		distance, _ := strconv.Atoi(distances[i])
		prod *= findWins(time, distance)
	}
	return prod
}

func findWins(totalTime, distance int) int {
	h := totalTime / 2
	l := 1
	//we want to find h&l where h is the first with a win
	for h-l > 1 {
		mid := (h + l) / 2
		if mid*(totalTime-mid) > distance {
			h = mid
		} else {
			l = mid
		}
	}
	// for 1 -7 if 2 -5 are the winners 7-2 -2-1
	return totalTime - h - (h - 1)
}

func Solve2(file string) int {
	lines := utils.ReadFile((file))
	data := strings.Split(lines[0], ":")
	totalTime, _ := strconv.Atoi(strings.ReplaceAll(data[1], " ", ""))
	data = strings.Split(lines[1], ":")
	dist, _ := strconv.Atoi(strings.ReplaceAll(data[1], " ", ""))

	return findWins(totalTime, dist)
}
