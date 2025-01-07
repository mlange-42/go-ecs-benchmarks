package bench

import (
	"log"
	"os"

	"github.com/mlange-42/go-ecs-benchmarks/bench/create2comp"
	"github.com/mlange-42/go-ecs-benchmarks/bench/query1in10"
	"github.com/mlange-42/go-ecs-benchmarks/bench/query2comp"
	"github.com/mlange-42/go-ecs-benchmarks/bench/query32arch"
	"github.com/mlange-42/go-ecs-benchmarks/bench/util"
)

func RunAll() {
	err := os.Mkdir("results", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	util.RunBenchmarks("query2comp.csv", query2comp.Benchmarks(), util.ToCSV)
	util.RunBenchmarks("query1in10.csv", query1in10.Benchmarks(), util.ToCSV)
	util.RunBenchmarks("query32arch.csv", query32arch.Benchmarks(), util.ToCSV)

	util.RunBenchmarks("create2comp.csv", create2comp.Benchmarks(), util.ToCSV)
}
