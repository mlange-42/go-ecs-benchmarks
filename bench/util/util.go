package util

import (
	"fmt"
	"os"
	"path"
	"runtime"
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

func RunBenchmarks(name string, benchmarks Benchmarks, count int, format func(Benchmarks) string) {
	fmt.Println("Running", name)

	benchmarks.time = make([][]float64, len(benchmarks.Benches))

	for i := range benchmarks.time {
		benchmarks.time[i] = make([]float64, len(benchmarks.N))
	}

	for i, n := range benchmarks.N {
		fmt.Println("   N =", n)
		for j := range benchmarks.Benches {
			bench := &benchmarks.Benches[j]
			fmt.Printf("       %-20s", bench.Name)
			var tSum int64
			var nSum int
			for range count {
				res := testing.Benchmark(func(b *testing.B) {
					bench.F(b, n)
				})
				tSum += res.T.Nanoseconds()
				nSum += res.N
				runtime.GC()
			}
			nanos := float64(tSum) / float64(nSum*n)
			benchmarks.time[j][i] = nanos
			fmt.Printf("%12s\n", toTime(nanos))
		}
	}

	os.WriteFile(path.Join("results", fmt.Sprintf("%s.csv", name)), []byte(format(benchmarks)), 0666)
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

func toTime(v float64) string {
	if v < 1_000 {
		return fmt.Sprintf("%.2fns", v)
	}
	if v < 1_000_000 {
		return fmt.Sprintf("%.2fus", v/1000)
	}
	return fmt.Sprintf("%.2fms", v/1_000_000)
}
