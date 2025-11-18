package addremovelarge

import (
	"testing"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runArk(b *testing.B, n int) {
	world := ecs.NewWorld(1024)

	velMap := ecs.NewMap1[comps.Velocity](world)

	mapper := ecs.NewMap11[
		comps.Position,
		comps.C1, comps.C2, comps.C3, comps.C4, comps.C5,
		comps.C6, comps.C7, comps.C8, comps.C9, comps.C10,
	](world)

	entities := make([]ecs.Entity, 0, n)
	mapper.NewBatchFn(n, func(entity ecs.Entity, a *comps.Position, b *comps.C1,
		c *comps.C2, d *comps.C3, e *comps.C4, f *comps.C5, g *comps.C6,
		h *comps.C7, i *comps.C8, j *comps.C9, k *comps.C10) {
		entities = append(entities, entity)
	})
	// Iterate once for more fairness
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

	velMap := ecs.NewMap1[comps.Velocity](world)

	mapper := ecs.NewMap11[
		comps.Position,
		comps.C1, comps.C2, comps.C3, comps.C4, comps.C5,
		comps.C6, comps.C7, comps.C8, comps.C9, comps.C10,
	](world)

	mapper.NewBatchFn(n, nil)

	filterPos := ecs.NewFilter1[comps.Position](world)
	filterPosVel := ecs.NewFilter2[comps.Position, comps.Velocity](world)

	// Iterate once for more fairness
	velMap.AddBatchFn(filterPos.Batch(), nil)
	velMap.RemoveBatch(filterPosVel.Batch(), nil)

	for b.Loop() {
		velMap.AddBatchFn(filterPos.Batch(), nil)
		velMap.RemoveBatch(filterPosVel.Batch(), nil)
	}
}
