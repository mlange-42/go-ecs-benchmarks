package query2comp

import (
	"github.com/mlange-42/go-ecs-benchmarks/bench/util"
)

func Benchmarks() util.Benchmarks {
	return util.Benchmarks{
		Benches: []util.Benchmark{
			{Name: "Arche", F: runArche},
			{Name: "Arche (cached)", F: runArcheRegistered},
			{Name: "Ark", F: runArk},
			{Name: "Donburi", F: runDonburi},
			{Name: "ggecs", F: runGGEcs},
			{Name: "uot", F: runUot},
			{Name: "Volt", F: runVolt},
		},
		N: []int{
			1, 4, 16, 64, 256, 1024, 16_000, 256_000, 1_000_000,
		},
	}
}
