package utils

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func ReadFile(path string) []string {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func GetAllNeighbouringPoints(x, y, w, h, gridW, gridH int, diagonals bool) [][2]int {
	pts := [][2]int{}
	searchRows := make([]int, 0, 2)
	startRow := y - 1
	if startRow >= 0 {
		searchRows = append(searchRows, startRow)
	}
	endRow := y + h
	if endRow < gridH {
		searchRows = append(searchRows, endRow)
	}

	for _, searchRow := range searchRows {
		startColumn := x - 1
		endColumn := x + w + 1
		for i := startColumn; i < endColumn; i++ {
			if i < 0 || i > gridW {
				continue
			}
			if !diagonals && (i == x-1 || i == x+w) {
				continue
			}
			pts = append(pts, [2]int{i, searchRow})
		}
	}
	for i := 0; i < h; i++ {
		if x-1 >= 0 {
			pts = append(pts, [2]int{x - 1, y + i})
		}
		if x+w < gridW {
			pts = append(pts, [2]int{x + w, y + i})
		}
	}
	return pts
}

func ReadNumbersFromLine(line string) []int {
	re := regexp.MustCompile(`-?\d+`)
	strs := re.FindAllString(line, -1)
	nums := make([]int, len(strs))
	for i, str := range strs {
		nums[i], _ = strconv.Atoi(str)
	}
	return nums
}
