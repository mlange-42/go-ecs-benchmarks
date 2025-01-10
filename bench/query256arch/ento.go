package query256arch

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/wfranczyk/ento"
)

func runEnto(b *testing.B, n int) {
	b.StopTimer()
	world := ento.NewWorldBuilder().
		WithDenseComponents(comps.Position{}).
		WithDenseComponents(comps.Velocity{}).
		WithDenseComponents(comps.C1{}).
		WithDenseComponents(comps.C2{}).
		WithDenseComponents(comps.C3{}).
		WithDenseComponents(comps.C4{}).
		WithDenseComponents(comps.C5{}).
		WithDenseComponents(comps.C6{}).
		WithDenseComponents(comps.C7{}).
		WithDenseComponents(comps.C8{}).
		Build(1024)

	system := PosVelSystem{}
	world.AddSystems(&system)

	extraComps := []any{comps.C1{}, comps.C2{}, comps.C3{}, comps.C4{}, comps.C5{}, comps.C6{}, comps.C7{}, comps.C8{}}

	for range n {
		world.AddEntity(comps.Position{}, comps.Velocity{})
	}

	ids := []any{}
	for i := range n * 4 {
		ids = append(ids, comps.Position{})
		for j, id := range extraComps {
			m := 1 << j
			if i&m == m {
				ids = append(ids, id)
			}
		}
		world.AddEntity(ids...)

		ids = ids[:0]
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		world.Update()
	}
}

type PosVelSystem struct {
	Pos *comps.Position `ento:"required"`
	Vel *comps.Velocity `ento:"required"`
}

func (s *PosVelSystem) Update(entity *ento.Entity) {
	s.Pos.X += s.Vel.X
	s.Pos.Y += s.Vel.Y
}
