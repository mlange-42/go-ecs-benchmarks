package delete2comp

import (
	"testing"

	ecs "github.com/BrownNPC/simple-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runSimpleECS(b *testing.B, n int) {
	b.StopTimer()
	p := ecs.New(n)
	ecs.Register2[comps.Position, comps.Velocity](p)
	for range n {
		e := ecs.NewEntity(p)
		ecs.Add2(p, e,
			comps.Position{},
			comps.Velocity{X: 1, Y: 1},
		)
	}
	POSITION, VELOCITY := ecs.GetStorage2[comps.Position, comps.Velocity](p)
	for b.Loop() {
		b.StartTimer()
		for _, e := range POSITION.And(VELOCITY) {
			ecs.Kill(p, e)
		}
		b.StopTimer()
		for range n {
			e := ecs.NewEntity(p)
			ecs.Add2(p, e,
				comps.Position{},
				comps.Velocity{X: 1, Y: 1},
			)
		}
	}
}

// modify pointers instead of copying components
func runSimpleECS_Batch(b *testing.B, n int) {
	b.StopTimer()
	p := ecs.New(n)
	ecs.Register2[comps.Position, comps.Velocity](p)
	for range n {
		e := ecs.NewEntity(p)
		ecs.Add2(p, e,
			comps.Position{},
			comps.Velocity{X: 1, Y: 1},
		)
	}
	POSITION, VELOCITY := ecs.GetStorage2[comps.Position, comps.Velocity](p)
	for b.Loop() {
		b.StartTimer()
		e := POSITION.And(VELOCITY)
		ecs.Kill(p, e...)

		b.StopTimer()
		for range n {
			e := ecs.NewEntity(p)
			ecs.Add2(p, e,
				comps.Position{},
				comps.Velocity{X: 1, Y: 1},
			)
		}
	}
}
