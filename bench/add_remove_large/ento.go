package addremovelarge

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
		WithDenseComponents(comps.C1{}).
		WithDenseComponents(comps.C2{}).
		WithDenseComponents(comps.C3{}).
		WithDenseComponents(comps.C4{}).
		WithDenseComponents(comps.C5{}).
		WithDenseComponents(comps.C6{}).
		WithDenseComponents(comps.C7{}).
		WithDenseComponents(comps.C8{}).
		WithDenseComponents(comps.C9{}).
		WithDenseComponents(comps.C10{}).
		Build(1024)

	add := AddVelSystem{}
	rem := RemVelSystem{}
	world.AddSystems(&add, &rem)

	for i := 0; i < n; i++ {
		world.AddEntity(
			comps.Position{},
			comps.C1{}, comps.C2{}, comps.C3{}, comps.C4{}, comps.C5{},
			comps.C6{}, comps.C7{}, comps.C8{}, comps.C9{}, comps.C10{},
		)
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
