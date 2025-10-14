package random

import (
	"log"
	"math/rand/v2"
	"testing"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/mlange-42/go-ecs-benchmarks/bench/util"
)

func runArche(b *testing.B, n int) {
	world := ecs.NewWorld(1024)

	posID := ecs.ComponentID[comps.Position](&world)

	entities := make([]ecs.Entity, 0, n)
	query := world.Batch().NewQ(n, posID)
	for query.Next() {
		entities = append(entities, query.Entity())
	}
	rand.Shuffle(n, util.Swap(entities))

	sum := 0.0
	loop := func() {
		for _, e := range entities {
			pos := (*comps.Position)(world.Get(e, posID))
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
