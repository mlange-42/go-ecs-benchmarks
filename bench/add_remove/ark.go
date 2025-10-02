package addremove

import (
	"testing"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runArk(b *testing.B, n int) {
	world := ecs.NewWorld(1024)

	posMap := ecs.NewMap1[comps.Position](&world)
	velMap := ecs.NewMap1[comps.Velocity](&world)

	entities := make([]ecs.Entity, 0, n)
	posMap.NewBatchFn(n, func(entity ecs.Entity, _ *comps.Position) {
		entities = append(entities, entity)
	})

	for _, e := range entities {
		velMap.AddFn(e, nil)
	}
	for _, e := range entities {
		velMap.Remove(e)
	}

	for b.Loop() {
		for _, e := range entities {
			velMap.AddFn(e, nil)
		}
		for _, e := range entities {
			velMap.Remove(e)
		}
	}
}

func runArkBatched(b *testing.B, n int) {
	world := ecs.NewWorld(1024)

	posMap := ecs.NewMap1[comps.Position](&world)
	velMap := ecs.NewMap1[comps.Velocity](&world)

	posMap.NewBatchFn(n, nil)

	filterPos := ecs.NewFilter1[comps.Position](&world)
	filterPosVel := ecs.NewFilter2[comps.Position, comps.Velocity](&world)

	// Iterate once for more fairness
	velMap.AddBatchFn(filterPos.Batch(), nil)
	velMap.RemoveBatch(filterPosVel.Batch(), nil)

	for b.Loop() {
		velMap.AddBatchFn(filterPos.Batch(), nil)
		velMap.RemoveBatch(filterPosVel.Batch(), nil)
	}
}
