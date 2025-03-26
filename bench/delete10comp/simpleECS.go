package delete10comp

import (
	"testing"

	ecs "github.com/BrownNPC/simple-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runSimpleECS(b *testing.B, n int) {
	b.StopTimer()
	p := ecs.New(n)
	for range n {
		e := ecs.NewEntity(p)
		ecs.Add3(p, e,
			comps.C1{}, comps.C2{}, comps.C3{},
		)
		ecs.Add3(p, e, comps.C4{}, comps.C5{}, comps.C6{})
		ecs.Add3(p, e, comps.C7{}, comps.C8{}, comps.C9{})
		ecs.Add(p, e, comps.C10{})
	}
	c1, c2, c3, c4, c5, c6, c7, c8, c9 := ecs.GetStorage9[comps.C1, comps.C2, comps.C3, comps.C4, comps.C5, comps.C6, comps.C7, comps.C8, comps.C9](p)
	c10 := ecs.GetStorage[comps.C10](p)
	for b.Loop() {
		b.StartTimer()
		for _, e := range c1.And(c2, c3, c4, c5, c6, c7, c8, c9, c10) {
			ecs.Kill(p, e)
		}
		b.StopTimer()
		for range n {
			e := ecs.NewEntity(p)
			ecs.Add3(p, e,
				comps.C1{}, comps.C2{}, comps.C3{},
			)
			ecs.Add3(p, e, comps.C4{}, comps.C5{}, comps.C6{})
			ecs.Add3(p, e, comps.C7{}, comps.C8{}, comps.C9{})
			ecs.Add(p, e, comps.C10{})
		}
	}
}

// modify pointers instead of copying components
func runSimpleECS_Batch(b *testing.B, n int) {
	b.StopTimer()
	p := ecs.New(n)
	for range n {
		e := ecs.NewEntity(p)
		ecs.Add3(p, e,
			comps.C1{}, comps.C2{}, comps.C3{},
		)
		ecs.Add3(p, e, comps.C4{}, comps.C5{}, comps.C6{})
		ecs.Add3(p, e, comps.C7{}, comps.C8{}, comps.C9{})
		ecs.Add(p, e, comps.C10{})
	}
	c1, c2, c3, c4, c5, c6, c7, c8, c9 := ecs.GetStorage9[comps.C1, comps.C2, comps.C3, comps.C4, comps.C5, comps.C6, comps.C7, comps.C8, comps.C9](p)
	c10 := ecs.GetStorage[comps.C10](p)
	for b.Loop() {
		b.StartTimer()
		e := c1.And(c2, c3, c4, c5, c6, c7, c8, c9, c10)
		ecs.Kill(p, e...)

		b.StopTimer()
		for range n {
			e := ecs.NewEntity(p)
			ecs.Add3(p, e,
				comps.C1{}, comps.C2{}, comps.C3{},
			)
			ecs.Add3(p, e, comps.C4{}, comps.C5{}, comps.C6{})
			ecs.Add3(p, e, comps.C7{}, comps.C8{}, comps.C9{})
			ecs.Add(p, e, comps.C10{})
		}
	}
}
