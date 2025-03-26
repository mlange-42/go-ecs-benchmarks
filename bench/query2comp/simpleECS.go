package query2comp

import (
	"testing"

	ecs "github.com/BrownNPC/simple-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runSimpleECS(b *testing.B, n int) {
	p := ecs.New(n)
	for range n {
		e := ecs.NewEntity(p)
		ecs.Add2(p, e,
			comps.Position{},
			comps.Velocity{X: 1, Y: 1},
		)
	}
	POSITION, VELOCITY := ecs.GetStorage2[comps.Position, comps.Velocity](p)
	query := POSITION.And(VELOCITY)
	for b.Loop() {
		for _, e := range query {
			pos, vel :=
				POSITION.Get(e), VELOCITY.Get(e)
			pos.X += vel.X
			pos.Y += vel.Y
			POSITION.Update(e, pos)
		}
	}
}
// modify pointers instead of copying components
func runSimpleECS_Ptr(b *testing.B, n int) {
	p := ecs.New(n)
	for range n {
		e := ecs.NewEntity(p)
		ecs.Add2(p, e,
			comps.Position{},
			comps.Velocity{X: 1, Y: 1},
		)
	}
	POSITION, VELOCITY := ecs.GetStorage2[comps.Position, comps.Velocity](p)
	query := POSITION.And(VELOCITY)
	for b.Loop() {
		for _, e := range query {
			pos, vel :=
				POSITION.GetPtrUnsafe(e), VELOCITY.GetPtrUnsafe(e)
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
}
