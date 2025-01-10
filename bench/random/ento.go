package random

/*
func runEnto(b *testing.B, n int) {
	b.StopTimer()
	world := ento.NewWorldBuilder().
		WithDenseComponents(comps.Position{}).
		Build(1024)

	entities := make([]*ento.Entity, 0, n)
	for i := 0; i < n; i++ {
		e := world.AddEntity(comps.Position{})
		entities = append(entities, e)
	}
	rand.Shuffle(n, util.Swap(entities))

	sum := 0.0
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for _, e := range entities {
			var comp *comps.Position
			e.Get(comp)
			sum += comp.X
		}
	}
	b.StopTimer()
	if sum > 0 {
		log.Fatal("error")
	}
}
*/
