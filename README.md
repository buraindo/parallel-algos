# parallel-algorithms

`go test -bench=Benchmark -benchtime=5x -cpu=4`

# CW 1
```
goos: windows
goarch: amd64
pkg: qsort
cpu: AMD Ryzen 9 5900X 12-Core Processor
BenchmarkQsortSeq-4                 5         114211360 ns/op
BenchmarkQsortPar-4                 5          35603320 ns/op
PASS
```

Result: ~3,2x boost

# CW 2
```
goos: windows
goarch: amd64
pkg: qsort
cpu: AMD Ryzen 9 5900X 12-Core Processor
BenchmarkBfsSeq-4              5        6098503880 ns/op
BenchmarkBfsPar-4              5        1956956780 ns/op
PASS
```

Result: ~3,1x boost