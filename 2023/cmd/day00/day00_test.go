package main

import (
	"strconv"
	"testing"
)

func Test00(t *testing.T) {

	res := Solve("0.txt")
	length, _ := strconv.Atoi(res)
	if length != 2 {
		t.Errorf("File should have 2 lines")
	}

}
