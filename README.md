# parallel-algos

`go test -bench=Benchmark -benchtime=5x -cpu=4`

# CW 1
```
goos: windows
goarch: amd64
pkg: qsort
cpu: AMD Ryzen 9 5900X 12-Core Processor
BenchmarkSeq-4                 5         114211360 ns/op
BenchmarkPar-4                 5          35603320 ns/op
PASS
```

Result: ~3,2x boost