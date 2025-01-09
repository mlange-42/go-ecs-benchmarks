package util

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"
)

type Benchmarks struct {
	Benches []Benchmark
	N       []int
	time    [][]float64
}

type Benchmark struct {
	Name string
	F    func(b *testing.B, n int)
}

func RunBenchmarks(file string, benchmarks Benchmarks, format func(Benchmarks) string) {
	fmt.Println("Running", file)

	benchmarks.time = make([][]float64, len(benchmarks.Benches))

	for i := range benchmarks.time {
		benchmarks.time[i] = make([]float64, len(benchmarks.N))
	}

	for i, n := range benchmarks.N {
		fmt.Println("   N =", n)
		for j := range benchmarks.Benches {
			bench := &benchmarks.Benches[j]
			fmt.Println("       ", bench.Name)
			res := testing.Benchmark(func(b *testing.B) {
				bench.F(b, n)
			})
			benchmarks.time[j][i] = float64(res.T.Nanoseconds()) / float64(res.N*n)
		}
	}

	os.WriteFile(path.Join("results", file), []byte(format(benchmarks)), 0666)
}

func ToCSV(benchmarks Benchmarks) string {
	b := strings.Builder{}

	b.WriteString("N,")
	for j := range benchmarks.Benches {
		bench := &benchmarks.Benches[j]
		b.WriteString(bench.Name)
		if j < len(benchmarks.Benches)-1 {
			b.WriteString(",")
		}
	}
	b.WriteString("\n")

	for i, n := range benchmarks.N {
		b.WriteString(fmt.Sprintf("%d,", n))
		for j := range benchmarks.Benches {
			t := benchmarks.time[j][i]

			tStr := fmt.Sprintf("%f", t)
			b.WriteString(tStr)
			if j < len(benchmarks.Benches)-1 {
				b.WriteString(",")
			}
		}
		b.WriteString("\n")
	}

	return b.String()
}

func Swap[T any](slice []T) func(i, j int) {
	return func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
