package query2comp

import (
	"fmt"
	"runtime"
	"testing"

	ecs "github.com/marioolofo/go-gameengine-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

// Component IDs
const (
	positionComponentID ecs.ComponentID = iota
	velocityComponentID
)

func runGGEcs(b *testing.B, n int) {
	world := ecs.NewWorld(1024)
	world.Register(ecs.NewComponentRegistry[comps.Position](positionComponentID))
	world.Register(ecs.NewComponentRegistry[comps.Velocity](velocityComponentID))

	for i := 0; i < n*10; i++ {
		world.NewEntity(positionComponentID)
	}
	for i := 0; i < n; i++ {
		e := world.NewEntity(positionComponentID, velocityComponentID)
		vel := (*comps.Velocity)(world.Component(e, velocityComponentID))
		vel.X = 1
		vel.Y = 1
	}

	mask := ecs.MakeComponentMask(positionComponentID, velocityComponentID)

	loop := func() {
		query := world.Query(mask)
		for query.Next() {
			pos := (*comps.Position)(query.Component(positionComponentID))
			vel := (*comps.Velocity)(query.Component(velocityComponentID))
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
	for b.Loop() {
		loop()
	}

	sum := 0.0
	query := world.Query(mask)
	for query.Next() {
		pos := (*comps.Position)(query.Component(positionComponentID))
		sum += pos.X + pos.Y
	}
	if sum != float64(n*b.N*2) {
		panic(fmt.Sprintf("Expected sum %d, got %.2f", n*b.N*2, sum))
	}
	runtime.KeepAlive(sum)
}
