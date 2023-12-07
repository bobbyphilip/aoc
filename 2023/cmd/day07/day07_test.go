package main

import (
	"testing"
)

func TestS1(t *testing.T) {
	res := Solve1("0.data")
	if res != 6440 {
		t.Errorf(" failed was %d expected %d", res, 6440)
	}
}
func TestS2(t *testing.T) {
	res := Solve2("0.data")
	if res != 5905 {
		t.Errorf(" failed was %d expected %d", res, 5905)
	}
}
func BenchmarkS1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve1("0.data")
	}
}

func BenchmarkS2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve2("0.data")
	}
}
