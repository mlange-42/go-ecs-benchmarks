package create10comp

import (
	"testing"

	ecs "github.com/BrownNPC/simple-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runSimpleECS(b *testing.B, n int) {
	b.StopTimer()
	p := ecs.New(n)
	ecs.Register2[comps.Position, comps.Velocity](p)
	ecs.Register5[comps.C1, comps.C2, comps.C3, comps.C4, comps.C5](p)
	ecs.Register5[comps.C6, comps.C7, comps.C8, comps.C9, comps.C10](p)
	b.StartTimer()
	for range n {
		e := ecs.NewEntity(p)
		ecs.Add2(p, e,
			comps.Position{},
			comps.Velocity{X: 1, Y: 1},
		)
		ecs.Add3(p, e, comps.C1{}, comps.C2{}, comps.C3{})
		ecs.Add3(p, e, comps.C4{}, comps.C5{}, comps.C6{})
		ecs.Add3(p, e, comps.C7{}, comps.C8{}, comps.C9{})
		ecs.Add(p, e, comps.C10{})
	}
}
