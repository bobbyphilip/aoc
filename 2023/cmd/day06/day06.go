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

	speed := 1
	time := totalTime - 1
	for speed*time <= distance {
		speed++
		time--
	}
	return time - speed + 1
}

func Solve2(file string) int {
	lines := utils.ReadFile((file))
	data := strings.Split(lines[0], ":")
	totalTime, _ := strconv.Atoi(strings.ReplaceAll(data[1], " ", ""))
	data = strings.Split(lines[1], ":")
	dist, _ := strconv.Atoi(strings.ReplaceAll(data[1], " ", ""))

	return findWins(totalTime, dist)
}
