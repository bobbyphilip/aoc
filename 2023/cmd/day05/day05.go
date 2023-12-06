package main

import (
	"aoc/internal/utils"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	result := Solve1("2023/cmd/day05/1.data")
	fmt.Printf("day 05 part 1 %d \n", result)
	result = Solve2("2023/cmd/day05/0.data")
	fmt.Printf("day 05 part 2 %d \n", result)
}

func Solve1(file string) int {
	min := math.MaxInt64
	lines := utils.ReadFile((file))
	maps := []map[int]int{
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	}

	sortedMapKeys := make([][]int, 7)
	r := regexp.MustCompile(`\d+`)
	mapIndex := -1
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		if r.MatchString(line) {
			vals := strings.Split(line, " ")
			key, _ := strconv.Atoi(vals[1])
			value, _ := strconv.Atoi(vals[0])
			count, _ := strconv.Atoi(vals[2])
			maps[mapIndex][key] = value
			maps[mapIndex][key+count-1] = value + count - 1
			sortedMapKeys[mapIndex] = append(sortedMapKeys[mapIndex], key)
			sortedMapKeys[mapIndex] = append(sortedMapKeys[mapIndex], key+count-1)
		} else {
			mapIndex++
		}
	}
	for i := 0; i < 7; i++ {
		sort.Ints(sortedMapKeys[i])
	}

	seedStrings := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	for _, seedString := range seedStrings {
		pos, _ := strconv.Atoi(seedString)
		for mapIndex := 0; mapIndex < 7; {
			pos = findPosFromMap(maps[mapIndex], sortedMapKeys[mapIndex], pos)
			mapIndex++
		}
		if pos < min {
			min = pos
		}
	}

	return min
}

func findPosFromMap(lookupMap map[int]int, keys []int, lookupValue int) int {
	for i, key := range keys {
		if key == lookupValue {
			return lookupMap[key]
		}
		if i == 0 && lookupValue < key {
			return lookupValue
		}
		if lookupValue < key {
			stored := lookupMap[key]
			return stored - (key - lookupValue)
		}
	}
	return lookupValue
}

func Solve2(file string) int {
	fmt.Printf("\n\n\n\n")
	min := math.MaxInt64
	lines := utils.ReadFile((file))

	seedStrings := strings.Split(strings.Split(lines[0], ": ")[1], " ")

	maps := []map[int]int{
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	}

	sortedMapKeys := make([][]int, 7)
	r := regexp.MustCompile(`\d+`)
	mapIndex := -1
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		if r.MatchString(line) {
			vals := strings.Split(line, " ")
			key, _ := strconv.Atoi(vals[1])
			value, _ := strconv.Atoi(vals[0])
			count, _ := strconv.Atoi(vals[2])
			maps[mapIndex][key] = value
			maps[mapIndex][key+count-1] = value + count - 1
			sortedMapKeys[mapIndex] = append(sortedMapKeys[mapIndex], key)
			sortedMapKeys[mapIndex] = append(sortedMapKeys[mapIndex], key+count-1)
		} else {
			mapIndex++
		}
	}
	for i := 0; i < 7; i++ {
		sort.Ints(sortedMapKeys[i])
	}

	seedGroup := make([]int, len(seedStrings))
	for i, seedString := range seedStrings {
		seedGroup[i], _ = strconv.Atoi(seedString)
	}

	for i := 0; i < len(seedGroup); i += 2 {
		currentSeedGroup := seedGroup[i]
		fmt.Printf(" seed start from %d to %d \n", currentSeedGroup, currentSeedGroup+seedGroup[i+1])
		for j := 0; j < seedGroup[i+1]; j++ {
			pos := currentSeedGroup + j
			for mapIndex := 0; mapIndex < 7; {
				pos = findPosFromMap(maps[mapIndex], sortedMapKeys[mapIndex], pos)
				mapIndex++
			}
			if pos < min {
				min = pos
			}
		}
	}

	return min
}
