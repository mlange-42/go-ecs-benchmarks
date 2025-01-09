package bench

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	addremove "github.com/mlange-42/go-ecs-benchmarks/bench/add_remove"
	addremovelarge "github.com/mlange-42/go-ecs-benchmarks/bench/add_remove_large"
	"github.com/mlange-42/go-ecs-benchmarks/bench/create10comp"
	"github.com/mlange-42/go-ecs-benchmarks/bench/create2comp"
	create2compalloc "github.com/mlange-42/go-ecs-benchmarks/bench/create2comp_alloc"
	"github.com/mlange-42/go-ecs-benchmarks/bench/delete10comp"
	"github.com/mlange-42/go-ecs-benchmarks/bench/delete2comp"
	newworld "github.com/mlange-42/go-ecs-benchmarks/bench/new_world"
	"github.com/mlange-42/go-ecs-benchmarks/bench/query1in10"
	"github.com/mlange-42/go-ecs-benchmarks/bench/query2comp"
	"github.com/mlange-42/go-ecs-benchmarks/bench/query32arch"
	"github.com/mlange-42/go-ecs-benchmarks/bench/util"
	"github.com/shirou/gopsutil/cpu"
)

func RunAll() {
	if err := os.Mkdir("results", os.ModePerm); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	if err := writeInfo(); err != nil {
		log.Fatal(err)
	}

	util.RunBenchmarks("query2comp.csv", query2comp.Benchmarks(), util.ToCSV)
	util.RunBenchmarks("query1in10.csv", query1in10.Benchmarks(), util.ToCSV)
	util.RunBenchmarks("query32arch.csv", query32arch.Benchmarks(), util.ToCSV)

	util.RunBenchmarks("create2comp.csv", create2comp.Benchmarks(), util.ToCSV)
	util.RunBenchmarks("create2comp_alloc.csv", create2compalloc.Benchmarks(), util.ToCSV)
	util.RunBenchmarks("create10comp.csv", create10comp.Benchmarks(), util.ToCSV)
	util.RunBenchmarks("add_remove.csv", addremove.Benchmarks(), util.ToCSV)
	util.RunBenchmarks("add_remove_large.csv", addremovelarge.Benchmarks(), util.ToCSV)

	util.RunBenchmarks("delete2comp.csv", delete2comp.Benchmarks(), util.ToCSV)
	util.RunBenchmarks("delete10comp.csv", delete10comp.Benchmarks(), util.ToCSV)

	util.RunBenchmarks("new_world.csv", newworld.Benchmarks(), util.ToCSV)
}

func writeInfo() error {
	text := strings.Builder{}
	fmt.Fprintf(&text, "Last run: %s  \n", time.Now().Format(time.RFC1123))
	infos, err := cpu.Info()
	if err != nil {
		return err
	}
	for _, info := range infos {
		fmt.Fprintf(&text, "CPU: %s\n", info.ModelName)
		break
	}
	err = os.WriteFile(path.Join("results", "info.md"), []byte(text.String()), 0666)
	if err != nil {
		return err
	}

	return nil
}
