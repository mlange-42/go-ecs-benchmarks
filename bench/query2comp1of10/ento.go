package query2comp1of10

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
		Build(1024)

	system := PosVelSystem{}
	world.AddSystems(&system)

	for i := 0; i < n*9; i++ {
		world.AddEntity(comps.Position{})
	}
	for i := 0; i < n; i++ {
		world.AddEntity(comps.Position{}, comps.Velocity{})
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
