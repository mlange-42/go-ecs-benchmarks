package addremove

import (
	"testing"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runArk(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.NewWorld(1024)

	posMap := ecs.NewMap1[comps.Position](&world)
	velMap := ecs.NewMap1[comps.Velocity](&world)

	posMap.NewBatchFn(n, nil)

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

	b.StartTimer()

	for i := 0; i < b.N; i++ {
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
	b.StopTimer()
	world := ecs.NewWorld(1024)

	posMap := ecs.NewMap1[comps.Position](&world)
	velMap := ecs.NewMap1[comps.Velocity](&world)

	posMap.NewBatchFn(n, nil)

	filterPos := ecs.NewFilter1[comps.Position](&world)
	filterPosVel := ecs.NewFilter2[comps.Position, comps.Velocity](&world)

	// Iterate once for more fairness
	velMap.AddBatchFn(filterPos.Batch(), nil)
	velMap.RemoveBatch(filterPosVel.Batch(), nil)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		velMap.AddBatchFn(filterPos.Batch(), nil)
		velMap.RemoveBatch(filterPosVel.Batch(), nil)
	}
}
