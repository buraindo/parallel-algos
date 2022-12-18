package main

import (
	"testing"
)

func benchmarkBfsSeq(b *testing.B) {
	graph := cubicGraph(500)
	b.ResetTimer()
	seqBfs(graph, 0)
}

func benchmarkBfsPar(b *testing.B) {
	graph := cubicGraph(500)
	b.ResetTimer()
	parBfs(graph, 0)
}

func BenchmarkBfsSeq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkBfsSeq(b)
	}
}

func BenchmarkBfsPar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkBfsPar(b)
	}
}

func TestBfsSeq(t *testing.T) {
	graph := cubicGraph(100)
	seqBfs(graph, 0)
}

func TestBfsPar(t *testing.T) {
	graph := cubicGraph(100)
	parBfs(graph, 0)
}
