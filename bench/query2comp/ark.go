package query2comp

import (
	"runtime"
	"testing"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runArk(b *testing.B, n int) {
	world := ecs.NewWorld(1024)

	mapper := ecs.NewMap2[comps.Position, comps.Velocity](&world)
	for range n {
		_ = mapper.NewEntity(&comps.Position{}, &comps.Velocity{X: 1, Y: 1})
	}

	filter := ecs.NewFilter2[comps.Position, comps.Velocity](&world)

	loop := func() {
		query := filter.Query()
		for query.Next() {
			pos, vel := query.Get()
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
	for b.Loop() {
		loop()
	}

	sum := 0.0
	query := filter.Query()
	for query.Next() {
		pos, _ := query.Get()
		sum += pos.X + pos.Y
	}
	runtime.KeepAlive(sum)
}

func runArkRegistered(b *testing.B, n int) {
	world := ecs.NewWorld(1024)

	mapper := ecs.NewMap2[comps.Position, comps.Velocity](&world)
	for range n {
		_ = mapper.NewEntity(&comps.Position{}, &comps.Velocity{X: 1, Y: 1})
	}

	filter := ecs.NewFilter2[comps.Position, comps.Velocity](&world).Register()

	loop := func() {
		query := filter.Query()
		for query.Next() {
			pos, vel := query.Get()
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
	for b.Loop() {
		loop()
	}

	sum := 0.0
	query := filter.Query()
	for query.Next() {
		pos, _ := query.Get()
		sum += pos.X + pos.Y
	}
	runtime.KeepAlive(sum)
}
