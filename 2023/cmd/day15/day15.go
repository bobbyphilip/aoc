package main

import (
	"aoc/internal/utils"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	result := Solve1("2023/cmd/day15/1.data")
	fmt.Printf("day 00 part 1 %d \n", result)
	result = Solve2("2023/cmd/day15/1.data")
	fmt.Printf("day 00 part 2 %d \n", result)
}

func Solve1(file string) int {
	sum := 0
	lines := utils.ReadFile((file))
	for _, line := range lines {
		seqs := strings.Split(line, ",")
		for _, seq := range seqs {
			sum += hash(seq)
		}
	}
	return sum
}

func hash(str string) int {
	cur := 0
	for _, l := range str {
		cur += int(l)
		cur *= 17
		cur %= 256
	}
	return cur
}

type KV struct {
	key string
	val int
}

func Solve2(file string) int {
	sum := 0
	lines := utils.ReadFile((file))
	re := regexp.MustCompile(`[=-]`)
	boxes := make([][]KV, 256)
	for _, line := range lines {
		words := strings.Split(line, ",")
		for _, word := range words {
			locs := re.FindStringIndex(word)
			key := word[0:locs[0]]
			num := -1
			if word[locs[0]] == '=' {
				num, _ = strconv.Atoi(word[locs[0]+1:])

			}
			h := hash(key)
			box := boxes[h]
			existingIndex := -1
			for index, kv := range box {
				if kv.key == key {
					existingIndex = index
					break
				}
			}
			newKV := KV{key: key, val: num}
			if num > -1 {
				// this is for =, so we need to append/replace in the list
				if existingIndex == -1 {
					box = append(box, newKV)
					boxes[h] = box
				} else {
					boxes[h][existingIndex].val = num
				}
			} else {
				if existingIndex > -1 {
					boxes[h] = slices.Delete(boxes[h], existingIndex, existingIndex+1)
				}
			}
		}
	}

	for i, box := range boxes {
		for j, kv := range box {
			sum += (i + 1) * (j + 1) * kv.val
		}
	}

	return sum
}
