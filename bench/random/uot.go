package random

import (
	"log"
	"math/rand/v2"
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/mlange-42/go-ecs-benchmarks/bench/util"
	"github.com/unitoftime/ecs"
)

func runUot(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.NewWorld()

	entities := make([]ecs.Id, 0, n)
	for i := 0; i < n; i++ {
		id := world.NewId()
		ecs.Write(world, id,
			ecs.C(comps.Position{}),
			ecs.C(comps.Velocity{}),
		)
		entities = append(entities, id)
	}
	rand.Shuffle(n, util.Swap(entities))

	sum := 0.0
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, e := range entities {
			pos := ecs.ReadPtr[comps.Position](world, e)
			sum += pos.X
		}
	}
	b.StopTimer()
	if sum > 0 {
		log.Fatal("error")
	}
}
