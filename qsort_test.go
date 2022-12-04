package main

import (
	"math/rand"
	"testing"
)

func benchmarkSeq(n int, b *testing.B) {
	arr := rand.Perm(n)

	b.ResetTimer()

	seq(arr)
}

func benchmarkPar(n int, b *testing.B) {
	arr := rand.Perm(n)

	b.ResetTimer()

	par(arr)
}

func BenchmarkSeq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkSeq(10000000, b)
	}
}

func BenchmarkPar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkPar(10000000, b)
	}
}
