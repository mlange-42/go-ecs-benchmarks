package query2comp

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/unitoftime/ecs"
)

func runUot(b *testing.B, n int) {
	world := ecs.NewWorld()

	for i := 0; i < n*10; i++ {
		id := world.NewId()
		ecs.Write(world, id,
			ecs.C(comps.Position{}),
		)
	}
	for i := 0; i < n; i++ {
		id := world.NewId()
		ecs.Write(world, id,
			ecs.C(comps.Position{}),
			ecs.C(comps.Velocity{X: 1, Y: 1}),
		)
	}
	query := ecs.Query2[comps.Position, comps.Velocity](world)

	loop := func() {
		query.MapId(func(id ecs.Id, pos *comps.Position, vel *comps.Velocity) {
			pos.X += vel.X
			pos.Y += vel.Y
		})
	}
	for b.Loop() {
		loop()
	}

	sum := 0.0
	query.MapId(func(id ecs.Id, pos *comps.Position, vel *comps.Velocity) {
		sum += pos.X + pos.Y
	})
	if sum != float64(n*b.N*2) {
		panic(fmt.Sprintf("Expected sum %d, got %.2f", n*b.N*2, sum))
	}
	runtime.KeepAlive(sum)
}
