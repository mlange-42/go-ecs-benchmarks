package bench

import (
	"log"
	"os"

	"github.com/mlange-42/go-ecs-benchmarks/bench/query2comp"
	"github.com/mlange-42/go-ecs-benchmarks/bench/query2comp1of10"
	query2compfrag "github.com/mlange-42/go-ecs-benchmarks/bench/query2comp_frag"
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

	util.RunBenchmarks("query2comp.csv", query2comp.Benchmarks(), util.ToCSV)
	util.RunBenchmarks("query2comp1of10.csv", query2comp1of10.Benchmarks(), util.ToCSV)
	util.RunBenchmarks("query2comp_frag.csv", query2compfrag.Benchmarks(), util.ToCSV)
}
