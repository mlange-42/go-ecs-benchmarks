package addremove

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/wfranczyk/ento"
)

func runEnto(b *testing.B, n int) {
	b.StopTimer()
	world := ento.NewWorldBuilder().
		WithDenseComponents(comps.Position{}).
		WithSparseComponents(comps.Velocity{}).
		Build(1024)

	add := AddVelSystem{}
	rem := RemVelSystem{}
	world.AddSystems(&add, &rem)

	for i := 0; i < n; i++ {
		world.AddEntity(comps.Position{})
	}

	// Iterate once for more fairness
	world.Update()

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		world.Update()
	}
}

type AddVelSystem struct {
	Pos *comps.Position `ento:"required"`
}

func (s *AddVelSystem) Update(entity *ento.Entity) {
	entity.Set(comps.Position{})
}

type RemVelSystem struct {
	Pos *comps.Position `ento:"required"`
	Vel *comps.Velocity `ento:"required"`
}

func (s *RemVelSystem) Update(entity *ento.Entity) {
	entity.Rem(comps.Position{})
}
