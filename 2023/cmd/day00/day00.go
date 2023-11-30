package main

import (
	"aoc/internal/utils"
	"fmt"
	"strconv"
)

func main() {
	result := Solve("2023/data/day00/0.txt")
	fmt.Println("day 0 " + result)
}

func Solve(file string) string {
	length := len(utils.ReadFile(file))
	return strconv.Itoa(length)
}
