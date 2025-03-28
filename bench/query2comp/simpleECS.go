package query2comp

import (
	"testing"

	ecs "github.com/BrownNPC/simple-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runSimpleECS(b *testing.B, n int) {
	world := ecs.New(n)
	for range n {
		e := ecs.NewEntity(world)
		ecs.Add2(world, e,
			comps.Position{},
			comps.Velocity{X: 1, Y: 1},
		)
	}
	stPosition, stVelocity := ecs.GetStorage2[comps.Position, comps.Velocity](world)
	for b.Loop() {
		for _, e := range stPosition.And(stVelocity) {
			pos, vel :=
				stPosition.Get(e), stVelocity.Get(e)
			pos.X += vel.X
			pos.Y += vel.Y
			stPosition.Update(e, pos)
		}
	}
	ecs.Kill(world, stPosition.And(stVelocity)...)

}

// modify pointers instead of copying components
func runSimpleECS_Ptr(b *testing.B, n int) {
	world := ecs.New(n)
	for range n {
		e := ecs.NewEntity(world)
		ecs.Add2(world, e,
			comps.Position{},
			comps.Velocity{X: 1, Y: 1},
		)
	}
	stPosition, stVelocity := ecs.GetStorage2[comps.Position, comps.Velocity](world)
	for b.Loop() {
		for _, e := range stPosition.And(stVelocity) {
			pos, vel :=
				stPosition.GetPtrUnsafe(e), stVelocity.GetPtrUnsafe(e)
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
	ecs.Kill(world, stPosition.And(stVelocity)...)
}
