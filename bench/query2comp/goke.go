package query2comp

import (
	"testing"
	"time"

	"github.com/kjkrol/goke"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	world := goke.New()

	goke.RegisterComponent[comps.Position](world)
	goke.RegisterComponent[comps.Velocity](world)

	posBP := goke.NewBlueprint1[comps.Position](world)
	posVelBP := goke.NewBlueprint2[comps.Position, comps.Velocity](world)

	for range n * 10 {
		_, _ = posBP.Create()
	}
	for range n {
		_, _, v := posVelBP.Create()
		v.X, v.Y = 1, 1
	}

	view := goke.NewView2[comps.Position, comps.Velocity](world)
	movementSystem := goke.RegisterSystemFunc(world, func(schedule *goke.Schedule, d time.Duration) {
		for head := range view.Values() {
			pos, vel := head.V1, head.V2
			pos.X += vel.X
			pos.Y += vel.Y
		}
	})

	goke.Plan(world, func(ctx goke.ExecutionContext, d time.Duration) {
		ctx.Run(movementSystem, d)
		ctx.Sync()
	})

	loop := func() {
		goke.Tick(world, 0)
	}
	for b.Loop() {
		loop()
	}
}
