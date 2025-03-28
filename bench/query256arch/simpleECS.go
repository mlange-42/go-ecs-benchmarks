package query256arch

import (
	"testing"

	ecs "github.com/BrownNPC/simple-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runSimpleECS(b *testing.B, n int) {
	// we could also initialize to n, decreasing setup time
	world := ecs.New(1024).EnableGrowing()
	//C1 to C8
	extraComponents := 8
	for i := range n * 4 {
		e := ecs.NewEntity(world)
		ecs.Add2(world, e, comps.Position{}, comps.Velocity{})

		for componentNum := range extraComponents {
			mask := 1 << componentNum
			if i&mask == mask {
				addComponent_C1toC8(world, e, componentNum)
			}
		}
	}
	stPosition, stVelocity := ecs.GetStorage2[comps.Position, comps.Velocity](world)
	for b.Loop() {
		for _, e := range stPosition.And(stVelocity) {
			pos, vel := stPosition.Get(e), stVelocity.Get(e)
			pos.X += vel.X
			pos.Y += vel.Y
			stPosition.Update(e, pos)
		}
	}
}
func runSimpleECS_Ptr(b *testing.B, n int) {
	// we could also initialize to n, decreasing setup time
	world := ecs.New(1024).EnableGrowing()
	//C1 to C8
	extraComponents := 8
	for i := range n * 4 {
		e := ecs.NewEntity(world)
		ecs.Add2(world, e, comps.Position{}, comps.Velocity{})

		for componentNum := range extraComponents {
			mask := 1 << componentNum
			if i&mask == mask {
				addComponent_C1toC8(world, e, componentNum)
			}
		}
	}
	stPosition, stVelocity := ecs.GetStorage2[comps.Position, comps.Velocity](world)
	for b.Loop() {
		for _, e := range stPosition.And(stVelocity) {
			pos, vel := stPosition.GetPtrUnsafe(e), stVelocity.GetPtrUnsafe(e)
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
}
func addComponent_C1toC8(p *ecs.Pool, e ecs.Entity, componentNum int) {
	switch componentNum {
	case 1:
		ecs.Add(p, e, comps.C1{})
	case 2:
		ecs.Add(p, e, comps.C2{})
	case 3:
		ecs.Add(p, e, comps.C3{})
	case 4:
		ecs.Add(p, e, comps.C4{})
	case 5:
		ecs.Add(p, e, comps.C5{})
	case 6:
		ecs.Add(p, e, comps.C6{})
	case 7:
		ecs.Add(p, e, comps.C7{})
	case 8:
		ecs.Add(p, e, comps.C8{})
	}
}
