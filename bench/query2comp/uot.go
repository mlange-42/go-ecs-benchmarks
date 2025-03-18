package query2comp

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/unitoftime/ecs"
)

func runUot(b *testing.B, n int) {
	world := ecs.NewWorld()

	for i := 0; i < n; i++ {
		id := world.NewId()
		ecs.Write(world, id,
			ecs.C(comps.Position{}),
			ecs.C(comps.Velocity{X: 1, Y: 1}),
		)
	}
	query := ecs.Query2[comps.Position, comps.Velocity](world)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		query.MapId(func(id ecs.Id, pos *comps.Position, vel *comps.Velocity) {
			pos.X += vel.X
			pos.Y += vel.Y
		})
	}
	b.StopTimer()

	query.MapId(func(id ecs.Id, pos *comps.Position, vel *comps.Velocity) {
		if pos.X == 0 || pos.Y == 0 {
			panic("assertion failed")
		}
	})
}
