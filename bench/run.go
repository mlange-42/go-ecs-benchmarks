package bench

import (
	posvel "github.com/mlange-42/go-ecs-benchmarks/bench/pos_vel"
	"github.com/mlange-42/go-ecs-benchmarks/bench/util"
)

func RunAll() {
	util.RunBenchmarks("PosVel", posvel.Benchmarks(), util.ToCSV)
}
