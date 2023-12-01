package main

import (
	"testing"
)

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
