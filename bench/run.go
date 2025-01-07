package bench

import (
	"log"
	"os"

	posvel "github.com/mlange-42/go-ecs-benchmarks/bench/pos_vel"
	"github.com/mlange-42/go-ecs-benchmarks/bench/util"
)

func RunAll() {
	err := os.Mkdir("results", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	err = os.Mkdir("plots", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	util.RunBenchmarks("query.csv", posvel.Benchmarks(), util.ToCSV)
}
