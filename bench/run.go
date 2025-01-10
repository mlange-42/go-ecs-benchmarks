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
	"github.com/mlange-42/go-ecs-benchmarks/bench/random"
	"github.com/mlange-42/go-ecs-benchmarks/bench/util"
	"github.com/shirou/gopsutil/cpu"
)

var benchmarks = map[string]func() util.Benchmarks{
	"query2comp":  query2comp.Benchmarks,
	"query1in10":  query1in10.Benchmarks,
	"query32arch": query32arch.Benchmarks,

	"random": random.Benchmarks,

	"create2comp":       create2comp.Benchmarks,
	"create2comp_alloc": create2compalloc.Benchmarks,
	"create10comp":      create10comp.Benchmarks,
	"add_remove":        addremove.Benchmarks,
	"add_remove_large":  addremovelarge.Benchmarks,

	"delete2comp":  delete2comp.Benchmarks,
	"delete10comp": delete10comp.Benchmarks,

	"new_world": newworld.Benchmarks,
}

func RunAll() {
	if err := os.Mkdir("results", os.ModePerm); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	if err := writeInfo(); err != nil {
		log.Fatal(err)
	}
	for name, fn := range benchmarks {
		util.RunBenchmarks(name, fn(), util.ToCSV)
	}
}

func Run(benches []string) {
	if err := os.Mkdir("results", os.ModePerm); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	if err := writeInfo(); err != nil {
		log.Fatal(err)
	}

	for _, b := range benches {
		if _, ok := benchmarks[b]; !ok {
			log.Fatalf("benchmark %s not found", b)
		}
	}

	for _, b := range benches {
		util.RunBenchmarks(b, benchmarks[b](), util.ToCSV)
	}
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
