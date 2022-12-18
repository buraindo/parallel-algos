package main

import (
	"math/rand"
	"testing"
)

func benchmarkQsortSeq(n int, b *testing.B) {
	arr := rand.Perm(n)

	b.ResetTimer()

	seqQsort(arr)
}

func benchmarkQsortPar(n int, b *testing.B) {
	arr := rand.Perm(n)

	b.ResetTimer()

	parQsort(arr)
}

func BenchmarkQsortSeq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkQsortSeq(10000000, b)
	}
}

func BenchmarkQsortPar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkQsortPar(10000000, b)
	}
}
