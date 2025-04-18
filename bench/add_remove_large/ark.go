package addremovelarge

import (
	"testing"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runArk(b *testing.B, n int) {
	world := ecs.NewWorld(1024)

	velMap := ecs.NewMap1[comps.Velocity](&world)

	mapper := ecs.NewMap11[
		comps.Position,
		comps.C1, comps.C2, comps.C3, comps.C4, comps.C5,
		comps.C6, comps.C7, comps.C8, comps.C9, comps.C10,
	](&world)

	mapper.NewBatchFn(n, nil)

	filterPos := ecs.NewFilter1[comps.Position](&world)
	filterPosVel := ecs.NewFilter2[comps.Position, comps.Velocity](&world)

	entities := make([]ecs.Entity, 0, n)

	// Iterate once for more fairness
	query1 := filterPos.Query()
	for query1.Next() {
		entities = append(entities, query1.Entity())
	}

	for _, e := range entities {
		velMap.AddFn(e, nil)
	}

	entities = entities[:0]
	query2 := filterPosVel.Query()
	for query2.Next() {
		entities = append(entities, query2.Entity())
	}

	for _, e := range entities {
		velMap.Remove(e)
	}

	entities = entities[:0]

	for b.Loop() {
		query1 = filterPos.Query()
		for query1.Next() {
			entities = append(entities, query1.Entity())
		}

		for _, e := range entities {
			velMap.AddFn(e, nil)
		}

		entities = entities[:0]
		query2 = filterPosVel.Query()
		for query2.Next() {
			entities = append(entities, query2.Entity())
		}

		for _, e := range entities {
			velMap.Remove(e)
		}

		entities = entities[:0]
	}
}

func runArkBatched(b *testing.B, n int) {
	world := ecs.NewWorld(1024)

	velMap := ecs.NewMap1[comps.Velocity](&world)

	mapper := ecs.NewMap11[
		comps.Position,
		comps.C1, comps.C2, comps.C3, comps.C4, comps.C5,
		comps.C6, comps.C7, comps.C8, comps.C9, comps.C10,
	](&world)

	mapper.NewBatchFn(n, nil)

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
