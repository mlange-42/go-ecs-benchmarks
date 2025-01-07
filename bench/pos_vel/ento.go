package posvel

import (
	"testing"

	"github.com/wfranczyk/ento"
)

func runEnto(b *testing.B, n int) {
	b.StopTimer()
	world := ento.NewWorldBuilder().
		WithDenseComponents(Position{}).
		WithSparseComponents(Velocity{}).
		Build(1024)

	system := PosVelSystem{}
	world.AddSystems(&system)

	for i := 0; i < n*5; i++ {
		world.AddEntity(Position{})
	}
	for i := 0; i < n; i++ {
		world.AddEntity(Position{}, Velocity{})
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		world.Update()
	}
}

type PosVelSystem struct {
	Pos *Position `ento:"required"`
	Vel *Velocity `ento:"required"`
}

func (s *PosVelSystem) Update(entity *ento.Entity) {
	s.Pos.X += s.Vel.X
	s.Pos.Y += s.Vel.Y
}
