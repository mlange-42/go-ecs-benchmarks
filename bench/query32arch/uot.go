package query32arch

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/unitoftime/ecs"
)

func runUot(b *testing.B, n int) {
	world := ecs.NewWorld()

	extraComps := []any{comps.C1{}, comps.C2{}, comps.C3{}, comps.C4{}, comps.C5{}}

	ids := []ecs.Component{}
	for i := range n {
		ids = append(ids, ecs.C(comps.Position{}), ecs.C(comps.Velocity{}))
		for j, id := range extraComps {
			m := 1 << j
			if i&m == m {
				ids = append(ids, ecs.C(id))
			}
		}

		id := world.NewId()
		ecs.Write(world, id, ids...)

		ids = ids[:0]
	}
	query := ecs.Query2[comps.Position, comps.Velocity](world)

	for b.Loop() {
		query.MapId(func(id ecs.Id, pos *comps.Position, vel *comps.Velocity) {
			pos.X += vel.X
			pos.Y += vel.Y
		})
	}
}
