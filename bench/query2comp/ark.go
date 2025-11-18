package query2comp

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runArk(b *testing.B, n int) {
	world := ecs.NewWorld(1024)

	ecs.NewMap1[comps.Position](world).
		NewBatchFn(n*10, nil)

	ecs.NewMap2[comps.Position, comps.Velocity](world).
		NewBatchFn(n, func(e ecs.Entity, p *comps.Position, v *comps.Velocity) {
			v.X, v.Y = 1, 1
		})

	filter := ecs.NewFilter2[comps.Position, comps.Velocity](world)

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
	if sum != float64(n*b.N*2) {
		panic(fmt.Sprintf("Expected sum %d, got %.2f", n*b.N*2, sum))
	}
	runtime.KeepAlive(sum)
}

func runArkRegistered(b *testing.B, n int) {
	world := ecs.NewWorld(1024)

	ecs.NewMap1[comps.Position](world).
		NewBatchFn(n*10, nil)

	ecs.NewMap2[comps.Position, comps.Velocity](world).
		NewBatchFn(n, func(e ecs.Entity, p *comps.Position, v *comps.Velocity) {
			v.X, v.Y = 1, 1
		})

	filter := ecs.NewFilter2[comps.Position, comps.Velocity](world).Register()

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
	if sum != float64(n*b.N*2) {
		panic(fmt.Sprintf("Expected sum %d, got %.2f", n*b.N*2, sum))
	}
	runtime.KeepAlive(sum)
}
