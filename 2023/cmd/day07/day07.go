package main

import (
	"aoc/internal/utils"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	result := Solve1("2023/cmd/day07/1.data")
	fmt.Printf("day 07 part 1 %d \n", result)
	result = Solve2("2023/cmd/day07/1.data")
	fmt.Printf("day 07 part 2 %d \n", result)
}

type Round struct {
	hand string
	bid  int
	rank int
}

func Solve1(file string) int {
	sum := 0
	lines := utils.ReadFile((file))
	rounds := make([]Round, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		hand := fields[0]
		bid, _ := strconv.Atoi(fields[1])
		rank := findHandType([]rune(hand))
		rounds[i] = Round{hand: hand, bid: bid, rank: rank}
	}
	sortRounds(rounds)
	for i := 0; i < len(rounds); i++ {
		round := rounds[i]
		sum += round.bid * (len(rounds) - i)
	}
	return sum
}

func findHandType(hand []rune) int {
	counterMap := map[rune]int{}
	for _, card := range hand {
		counterMap[card] = counterMap[card] + 1
	}

	vals := make([]int, 0, len(counterMap))
	for _, value := range counterMap {
		vals = append(vals, value)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(vals)))
	if vals[0] == 5 {
		return 0
	}
	if vals[0] == 4 {
		return 1
	}
	if vals[0] == 3 {
		if vals[1] == 2 {
			return 2
		} else {
			return 3
		}
	}
	if vals[0] == 2 {
		if vals[1] == 2 {
			return 4
		} else {
			return 5
		}
	}
	return 6
}

func sortRounds(rounds []Round) {
	slices.SortFunc(rounds, func(l, r Round) int {
		if l.rank == r.rank {
			for i := range l.hand {
				lr := rune(l.hand[i])
				rr := rune(r.hand[i])
				if lr == rr {
					continue
				}

				return compareCard(lr, rr)
			}
		} else {
			return l.rank - r.rank
		}
		return 0
	})
}

func compareCard(l, r rune) int {
	if l == 'A' {
		return -1
	}
	if r == 'A' {
		return 1
	}
	if l == 'K' {
		return -1
	}
	if r == 'K' {
		return 1
	}
	if l == 'Q' {
		return -1
	}
	if r == 'Q' {
		return 1
	}
	if l == 'J' {
		return -1
	}
	if r == 'J' {
		return 1
	}
	if l == 'T' {
		return -1
	}
	if r == 'T' {
		return 1
	}
	return int(r - l)

}
func Solve2(file string) int {
	sum := 0
	lines := utils.ReadFile((file))
	rounds := make([]Round, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		hand := fields[0]
		bid, _ := strconv.Atoi(fields[1])
		rank := findHandTypeV2([]rune(hand))
		rounds[i] = Round{hand: hand, bid: bid, rank: rank}
	}
	sortRoundsv2(rounds)
	for i := 0; i < len(rounds); i++ {
		round := rounds[i]
		sum += round.bid * (len(rounds) - i)
	}
	return sum
}

func findHandTypeV2(hand []rune) int {
	counterMap := map[rune]int{}
	jCount := 0
	for _, card := range hand {
		if card == 'J' {
			jCount++
			continue
		}
		counterMap[card] = counterMap[card] + 1
	}
	vals := make([]int, 0, len(counterMap))
	for _, value := range counterMap {
		vals = append(vals, value)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(vals)))
	if len(vals) == 0 {
		vals = append(vals, 0)
	}
	vals[0] += jCount
	if vals[0] == 5 {
		return 0
	}
	if vals[0] == 4 {
		return 1
	}
	if vals[0] == 3 {
		if vals[1] == 2 {
			return 2
		} else {
			return 3
		}
	}
	if vals[0] == 2 {
		if vals[1] == 2 {
			return 4
		} else {
			return 5
		}
	}
	return 6

}

func compareCardv2(l, r rune) int {
	if l == 'A' {
		return -1
	}
	if r == 'A' {
		return 1
	}
	if l == 'K' {
		return -1
	}
	if r == 'K' {
		return 1
	}
	if l == 'Q' {
		return -1
	}
	if r == 'Q' {
		return 1
	}
	if l == 'T' {
		return -1
	}
	if r == 'T' {
		return 1
	}
	if l >= '2' && l <= '9' {
		if r >= '2' && r <= '9' {
			return int(r - l)
		}
		return -1
	}
	if r >= '2' && r <= '9' {
		return 1
	}
	return 0

}

func sortRoundsv2(rounds []Round) {
	slices.SortFunc(rounds, func(l, r Round) int {
		if l.rank == r.rank {
			for i := range l.hand {
				lr := rune(l.hand[i])
				rr := rune(r.hand[i])
				if lr == rr {
					continue
				}

				return compareCardv2(lr, rr)
			}
		} else {
			return l.rank - r.rank
		}
		return 0
	})
}
