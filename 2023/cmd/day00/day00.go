package main

import (
	"aoc/internal/utils"
	"log"
	"strconv"
)

func main() {
	result := Solve("2023/data/day00/0.txt")
	log.Println("day 0 " + result)
}

func Solve(file string) string {
	length := len(utils.ReadFile(file))
	return strconv.Itoa(length)
}
