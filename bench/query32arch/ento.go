package query32arch

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
		Build(1024)

	system := PosVelSystem{}
	world.AddSystems(&system)

	extraComps := []any{comps.C1{}, comps.C2{}, comps.C3{}, comps.C4{}, comps.C5{}}
	ids := []any{}
	for i := 0; i < n; i++ {
		ids = append(ids, comps.Position{}, comps.Velocity{})
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
