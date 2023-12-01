package main

import (
	"testing"
)

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

func TestS2(t *testing.T) {
	res := Solve2("1.data")
	if res != 53539 {
		t.Errorf(" failed was %d expected %d", res, 53539)
	}
}
