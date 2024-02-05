### Benchmarks

```
goos: windows
goarch: amd64
pkg: learn-go-with-tests/13-sync
cpu: AMD Ryzen 7 5700X 8-Core Processor
BenchmarkCounter-16                     249205449                4.800 ns/op
BenchmarkConcurrentCounter-16            8987026               135.5 ns/op
BenchmarkAtomicCounter-16               705815508                1.675 ns/op
BenchmarkConcurrentAtomicCounter-16      8331214               133.0 ns/op
```