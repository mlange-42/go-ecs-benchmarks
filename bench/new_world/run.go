package newworld

import (
	"github.com/mlange-42/go-ecs-benchmarks/bench/util"
)

// Benchmarks runs the benchmarks.
func Benchmarks() util.Benchmarks {
	return util.Benchmarks{
		Benches: []util.Benchmark{
			{Name: "Arche", F: runArche},
			{Name: "Ark", F: runArche},
			{Name: "Donburi", F: runDonburi},
			{Name: "ggecs", F: runGGEcs},
			{Name: "uot", F: runUot},
			{Name: "Volt", F: runVolt},
		},
		N: []int{
			1,
		},
	}
}
