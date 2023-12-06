package main

import (
	"testing"
)

func TestS1(t *testing.T) {
	res := Solve1("0.data")
	if res != 288 {
		t.Errorf(" failed was %d expected %d", res, 288)
	}
}
func TestS2(t *testing.T) {
	res := Solve2("0.data")
	if res != 71503 {
		t.Errorf(" failed was %d expected %d", res, 71503)
	}
}
func BenchmarkS1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve1("1.data")
	}
}

func BenchmarkS2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve2("1.data")
	}
}
