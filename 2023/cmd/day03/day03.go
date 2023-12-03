package main

import (
	"aoc/internal/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	result := Solve1("2023/cmd/day03/1.data")
	fmt.Printf("day 03 part 1 %d \n", result)
	result = Solve2("2023/cmd/day03/1.data")
	fmt.Printf("day 03 part 2 %d \n", result)
}

var runes = [][]rune{}

func Solve1(file string) int {
	lines := utils.ReadFile((file))
	runes = make([][]rune, len(lines))
	for i, line := range lines {
		runes[i] = []rune(line)
	}
	sum := 0
	r := regexp.MustCompile(`\d+`)
	for i, line := range lines {
		numbers := r.FindAllStringIndex(line, -1)
		for _, number := range numbers {
			numberStart := number[0]
			numberEnd := number[1]
			if isPartNumber(i, numberStart, numberEnd) {
				partNum, _ := strconv.Atoi(line[numberStart:numberEnd])
				sum += partNum
			}
		}
	}
	return sum
}

func isPartNumber(lineNumber, startIndex, endIndex int) bool {
	lineLength := len(runes[lineNumber])

	pts := utils.GetAllNeighbouringPoints(startIndex, lineNumber, endIndex-startIndex, 1, lineLength, len(runes), true)
	for _, pt := range pts {
		if isMatchPresent(runes[pt[1]], lineLength, pt[0], isSymbol) {
			return true
		}
	}
	return false
}

func isMatchPresent(rs []rune, length, index int, checker func(rune) bool) bool {
	if index < 0 || index >= length {
		return false
	}
	r := rs[index]
	return checker(r)
}
func isSymbol(r rune) bool {
	if (r >= '0' && r <= '9') || r == '.' {
		return false
	}
	return true
}

func isStar(r rune) bool {
	return r == '*'
}

func Solve2(file string) int {
	lines := utils.ReadFile((file))
	runes = make([][]rune, len(lines))
	for i, line := range lines {
		runes[i] = []rune(line)
	}
	sum := 0
	r := regexp.MustCompile(`\d+`)
	starsToNumbers := map[string][]string{}
	for i, line := range lines {
		numbers := r.FindAllStringIndex(line, -1)
		for _, number := range numbers {
			numberStart := number[0]
			numberEnd := number[1]
			starLocs := getStarLocations(i, numberStart, numberEnd)
			for _, starLoc := range starLocs {
				numberList := starsToNumbers[starLoc]
				if numberList == nil {
					numberList = make([]string, 0)
				}
				numberList = append(numberList, fmt.Sprintf("%d:%d:%d", i, numberStart, numberEnd))
				starsToNumbers[starLoc] = numberList
			}
		}
	}
	for _, v := range starsToNumbers {
		if len(v) == 2 {
			gear1 := getGearValue(v[0])
			gear2 := getGearValue(v[1])
			sum += (gear1 * gear2)
		}
	}
	return sum
}

func getGearValue(input string) int {
	locDetails := strings.Split(input, ":")
	lineNumber, _ := strconv.Atoi(locDetails[0])
	start, _ := strconv.Atoi(locDetails[1])
	end, _ := strconv.Atoi(locDetails[2])
	val := runes[lineNumber][start:end]
	gear, _ := strconv.Atoi(string(val))
	return gear

}

func getStarLocations(lineNumber, startIndex, endIndex int) []string {
	lineLength := len(runes[lineNumber])
	pts := utils.GetAllNeighbouringPoints(startIndex, lineNumber, endIndex-startIndex, 1, lineLength, len(runes), true)
	starLocations := map[string]bool{}
	for _, pt := range pts {
		if isMatchPresent(runes[pt[1]], lineLength, pt[0], isStar) {
			starLocations[fmt.Sprintf("%d:%d", pt[0], pt[1])] = true
		}
	}

	locs := make([]string, len(starLocations))
	i := 0
	for k, _ := range starLocations {
		locs[i] = k
		i++
	}
	return locs
}
