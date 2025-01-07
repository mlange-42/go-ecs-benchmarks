package newworld

import (
	"github.com/mlange-42/go-ecs-benchmarks/bench/util"
)

func Benchmarks() util.Benchmarks {
	return util.Benchmarks{
		Benches: []util.Benchmark{
			{Name: "Arche", F: runArche},
			{Name: "Donburi", F: runDonburi},
			{Name: "Ento", F: runEnto},
			{Name: "ggecs", F: runGGEcs},
			{Name: "uot", F: runUot},
		},
		N: []int{
			1,
		},
	}
}
