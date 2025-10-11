package main

import (
	"flag"
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench"
)

func main() {
	testing.Init()

	count := flag.Int("count", 1, "number of times to run the benchmark")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		bench.RunAll(*count)
		return
	}

	bench.Run(args, *count)
}
