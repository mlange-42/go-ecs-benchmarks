package main

import (
	"flag"
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench"
)

func main() {
	testing.Init()
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		bench.RunAll()
		return
	}

	bench.Run(args)
}
