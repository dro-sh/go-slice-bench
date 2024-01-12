# Go Insert Bench

Benchmarks comparing appending at slice via `append` function and set via index.

Machine **MacBook Air M1 2021**.

`append.txt` contains results of using `append` function.
`index.txt` contains results of using set via index.

## Benchmarking

```bash
go test -bench=. -benchmem -count 10 | tee index.txt
go test -bench=. -benchmem -count 10 | tee append.txt
```

## Result

[benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat) result:

```bash
benchstat append.txt index.txt
```

```
goos: darwin
goarch: arm64
pkg: slices-bench
                                 │ append.txt  │              index.txt              │
                                 │   sec/op    │    sec/op     vs base               │
InsertAtSlice/input_size_5-8       276.9n ± 5%   273.9n ±  2%       ~ (p=0.109 n=10)
InsertAtSlice/input_size_10-8      516.2n ± 1%   504.3n ± 16%       ~ (p=0.210 n=10)
InsertAtSlice/input_size_50-8      2.461µ ± 3%   2.294µ ±  4%  -6.79% (p=0.000 n=10)
InsertAtSlice/input_size_100-8     4.730µ ± 2%   4.403µ ±  2%  -6.91% (p=0.002 n=10)
InsertAtSlice/input_size_500-8     30.34µ ± 3%   27.93µ ±  4%  -7.94% (p=0.000 n=10)
InsertAtSlice/input_size_1000-8    61.74µ ± 3%   58.81µ ±  1%  -4.75% (p=0.001 n=10)
InsertAtSlice/input_size_10000-8   692.4µ ± 2%   665.7µ ±  9%       ~ (p=0.436 n=10)
geomean                            8.033µ        7.644µ        -4.84%

                                 │  append.txt  │               index.txt               │
                                 │     B/op     │     B/op      vs base                 │
InsertAtSlice/input_size_5-8         320.0 ± 0%     320.0 ± 0%       ~ (p=1.000 n=10) ¹
InsertAtSlice/input_size_10-8        640.0 ± 0%     640.0 ± 0%       ~ (p=1.000 n=10) ¹
InsertAtSlice/input_size_50-8      3.406Ki ± 0%   3.406Ki ± 0%       ~ (p=1.000 n=10) ¹
InsertAtSlice/input_size_100-8     6.312Ki ± 0%   6.312Ki ± 0%       ~ (p=1.000 n=10) ¹
InsertAtSlice/input_size_500-8     34.94Ki ± 0%   34.94Ki ± 0%       ~ (p=1.000 n=10) ¹
InsertAtSlice/input_size_1000-8    70.66Ki ± 0%   70.66Ki ± 0%       ~ (p=1.000 n=10) ¹
InsertAtSlice/input_size_10000-8   775.9Ki ± 0%   775.9Ki ± 0%       ~ (p=0.075 n=10)
geomean                            9.694Ki        9.694Ki       +0.00%
¹ all samples are equal

                                 │ append.txt  │              index.txt               │
                                 │  allocs/op  │  allocs/op   vs base                 │
InsertAtSlice/input_size_5-8        11.00 ± 0%    11.00 ± 0%       ~ (p=1.000 n=10) ¹
InsertAtSlice/input_size_10-8       21.00 ± 0%    21.00 ± 0%       ~ (p=1.000 n=10) ¹
InsertAtSlice/input_size_50-8       101.0 ± 0%    101.0 ± 0%       ~ (p=1.000 n=10) ¹
InsertAtSlice/input_size_100-8      201.0 ± 0%    201.0 ± 0%       ~ (p=1.000 n=10) ¹
InsertAtSlice/input_size_500-8     1.401k ± 0%   1.401k ± 0%       ~ (p=1.000 n=10) ¹
InsertAtSlice/input_size_1000-8    2.901k ± 0%   2.901k ± 0%       ~ (p=1.000 n=10) ¹
InsertAtSlice/input_size_10000-8   29.90k ± 0%   29.90k ± 0%       ~ (p=1.000 n=10) ¹
geomean                             344.0         344.0       +0.00%
¹ all samples are equal

```
