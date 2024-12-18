package main

import (
	"aoc/internal/utils"
	"fmt"
)

func main() {
	result := Solve1("2024/cmd/day04/1.data")
	fmt.Printf("day 04 part 1 %d \n", result)
	result = Solve2("2024/cmd/day04/1.data")
	fmt.Printf("day 04 part 2 %d \n", result)
}

func Solve1(file string) int {
	sum := 0
	lines := utils.ReadFile((file))
	for row, line := range lines {
		for column, c := range line {
			if c == 'X' {
				if lookForXMAS(lines, row, column, 1, 0) {
					sum++
				}
				if lookForXMAS(lines, row, column, 1, 1) {
					sum++
				}
				if lookForXMAS(lines, row, column, 0, 1) {
					sum++
				}
				if lookForXMAS(lines, row, column, -1, 1) {
					sum++
				}
				if lookForXMAS(lines, row, column, -1, 0) {
					sum++
				}
				if lookForXMAS(lines, row, column, -1, -1) {
					sum++
				}
				if lookForXMAS(lines, row, column, 0, -1) {
					sum++
				}
				if lookForXMAS(lines, row, column, 1, -1) {
					sum++
				}
			}
		}

	}
	return sum
}

func lookForXMAS(lines []string, row, column, dirY, dirX int) bool {
	if row+3*dirY < 0 || row+3*dirY >= len(lines) {
		return false
	}
	if column+3*dirX < 0 || column+3*dirX >= len(lines[0]) {
		return false
	}
	word := "XMAS"
	for i := 1; i < 4; i++ {
		if lines[row+dirY*i][column+dirX*i] != word[i] {
			return false
		}
	}
	return true

}

func Solve2(file string) int {
	sum := 0
	lines := utils.ReadFile((file))
	for row, line := range lines {
		for column, c := range line {
			if c == 'M' || c == 'S' {
				if !lookForMAS(lines, c, row, column, 1, 1) {
					continue
				}
				if column+2 >= len(lines[0]) {
					continue
				}
				nextC := lines[row][column+2]
				if nextC == 'M' || nextC == 'S' {
					if lookForMAS(lines, rune(nextC), row, column+2, 1, -1) {
						sum++
					}
				}
			}
		}
	}
	return sum
}
func lookForMAS(lines []string, letter rune, row, column, dirY, dirX int) bool {
	if row+2*dirY < 0 || row+2*dirY >= len(lines) {
		return false
	}
	if column+2*dirX < 0 || column+2*dirX >= len(lines[0]) {
		return false
	}
	word := "MAS"
	if letter == 'S' {
		word = "SAM"
	}
	for i := 1; i < 3; i++ {
		if lines[row+dirY*i][column+dirX*i] != word[i] {
			return false
		}
	}
	return true

}
