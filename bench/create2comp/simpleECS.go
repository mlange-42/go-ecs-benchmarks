package create2comp

import (
	"testing"

	ecs "github.com/BrownNPC/simple-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runSimpleECS(b *testing.B, n int) {
	b.StopTimer()
	p := ecs.New(n)
	b.StartTimer()
	for range n {
		e := ecs.NewEntity(p)
		ecs.Add2(p, e,
			comps.Position{},
			comps.Velocity{X: 1, Y: 1},
		)
	}
}
