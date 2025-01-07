package query2compfrag

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/unitoftime/ecs"
)

func runUot(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.NewWorld()

	extraComps := []any{comps.C1{}, comps.C2{}, comps.C3{}, comps.C4{}, comps.C5{}}

	for i := 0; i < n; i++ {
		id := world.NewId()
		ecs.Write(world, id,
			ecs.C(comps.Position{}),
			ecs.C(comps.Velocity{}),
			ecs.C(extraComps[i%len(extraComps)]),
		)
	}
	query := ecs.Query2[comps.Position, comps.Velocity](world)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		query.MapId(func(id ecs.Id, pos *comps.Position, vel *comps.Velocity) {
			pos.X += vel.X
			pos.Y += vel.Y
		})
	}
}
