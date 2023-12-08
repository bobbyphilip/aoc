package main

import (
	"testing"
)

func TestS1(t *testing.T) {
	res := Solve1("0.data")
	if res != 6 {
		t.Errorf(" failed was %d expected %d", res, 6)
	}
}
func BenchmarkS1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve1("1.data")
	}
}
