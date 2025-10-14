package random

import (
	"log"
	"math/rand/v2"
	"testing"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/mlange-42/go-ecs-benchmarks/bench/util"
)

func runArk(b *testing.B, n int) {
	world := ecs.NewWorld(1024)

	mapper := ecs.NewMap[comps.Position](&world)
	entities := make([]ecs.Entity, 0, n)
	for range n {
		e := mapper.NewEntity(&comps.Position{})
		entities = append(entities, e)
	}
	rand.Shuffle(n, util.Swap(entities))

	sum := 0.0
	loop := func() {
		for _, e := range entities {
			pos := mapper.Get(e)
			sum += pos.X
		}
	}
	for b.Loop() {
		loop()
	}
	if sum > 0 {
		log.Fatal("error")
	}
}
