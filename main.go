package main

import (
	"flag"
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench"
)

func main() {
	testing.Init()
	flag.Parse()

	bench.RunAll()
}
