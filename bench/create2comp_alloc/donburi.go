package create2compalloc

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/component"
)

func runDonburi(b *testing.B, n int) {
	allIDs := []component.IComponentType{
		donburi.NewComponentType[comps.Position](),
		donburi.NewComponentType[comps.Velocity](),
	}

	for b.Loop() {
		b.StopTimer()
		world := donburi.NewWorld()

		b.StartTimer()
		for range n {
			world.Create(allIDs...)
		}
	}
}
